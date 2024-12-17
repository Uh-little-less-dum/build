package build_config

import (
	"os"

	config_loc_strategies "github.com/Uh-little-less-dum/build/pkg/buildConstants/configLocationStrategies"
	form_data "github.com/Uh-little-less-dum/build/pkg/buildManager/formData"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/classesKinda/plugin"
	env_vars "github.com/Uh-little-less-dum/build/pkg/envVars"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/Uh-little-less-dum/build/pkg/utils"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// FIX:
// RESUME: Move most of these fields, particularly those with structs to pointers.
type BuildManager struct {
	form_data.BuildFormData
	AppConfigPath        string
	ConfigDirPath        string
	LogLevel             string
	Fs                   *afero.Fs
	RootPackageJson      file_handlers_package_json.PackageJsonHandler
	Plugins              []*ulld_plugin.Plugin
	LogFile              string
	stack                []build_stages.BuildStage
	skipStages           []build_stages.BuildStage
	allowGoBack          bool
	appConfig            *app_config.AppConfig
	appConfigLocStrategy config_loc_strategies.ConfigLocationStrategy
	packageManager       package_managers.PackageManagerId
	Program              *tea.Program
}

var b *BuildManager

// Returns the BuildManager singleton.
func GetBuildManager() *BuildManager {
	return b
}

func (b *BuildManager) SetTargetDir(d string) {
	bp := afero.NewBasePathFs(afero.NewOsFs(), d)
	b.Fs = &bp
	b.BuildFormData.SetTargetDirOnlyInBuildConfig(d)
}

func (b *BuildManager) AppConfig() *app_config.AppConfig {
	return b.appConfig
}

func (b *BuildManager) SetPackageManager(p package_managers.PackageManagerId) {
	b.packageManager = p
}

func (b *BuildManager) PackageManager() package_managers.PackageManager {
	return package_managers.GetPackageManagerStruct(b.packageManager)
}

func (b *BuildManager) Installables() []types.Installable {
	res := make([]types.Installable, len(b.Plugins))
	for i, l := range b.Plugins {
		var r types.Installable = l
		res[i] = r
	}
	return res
}

func (b *BuildManager) SetAppConfig(ac *app_config.AppConfig) {
	b.appConfig = ac
}

func (b *BuildManager) SetPlugins(p []*ulld_plugin.Plugin) {
	b.Plugins = p
}

func (b *BuildManager) SetRootPackageJson(p file_handlers_package_json.PackageJsonHandler) {
	b.RootPackageJson = p
}

func (b *BuildManager) SetConfigLocationStrategy(strategyType config_loc_strategies.ConfigLocationStrategyId, value ...string) {
	switch strategyType {
	case config_loc_strategies.WaitForUserBeforeBuild:
		b.appConfigLocStrategy = config_loc_strategies.ConfigLocationStrategy{Id: strategyType, Value: ""}
	case config_loc_strategies.CopyFromLocation:
		if (len(value) == 1) && (value[0] != "") {
			b.appConfigLocStrategy = config_loc_strategies.ConfigLocationStrategy{Id: strategyType, Value: value[0]}
			b.AppConfigPath = value[0]
		}
	}
}

func GetInitialBuildManager() *BuildManager {
	ac := utils.GetDefaultAppConfigPath()
	val := BuildManager{
		AppConfigPath: ac,
		appConfig:     app_config.NewAppConfig(ac),
	}

	if val.BuildFormData.TargetDir() == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		val.SetTargetDir(cwd)
	}
	val.packageManager = package_managers.PnpmId
	// Handle initial values with environment variables here.
	val.ConfigDirPath = os.Getenv(string(env_vars.AdditionalSources))
	val.LogLevel = os.Getenv(string(env_vars.LogLevel))
	val.LogFile = os.Getenv(string(env_vars.LogFile))
	val.stack = []build_stages.BuildStage{build_stages.ConfirmCurrentDirStage}
	return &val
}

func init() {
	b = GetInitialBuildManager()
}

func (b *BuildManager) SetInitialStage(s build_stages.BuildStage) {
	b.stack = []build_stages.BuildStage{s}
}

// Adds a build stage that should be skipped to the list of skipped stages. This list can then be accessed from within each build stage at which point the stage can be bypassed.
func (b *BuildManager) AddSkippedStage(stage build_stages.BuildStage) {
	b.skipStages = append(b.skipStages, stage)
}

func (b *BuildManager) RemoveStageFromSkipped(stage build_stages.BuildStage) {
	var skipStages []build_stages.BuildStage
	for _, s := range b.skipStages {
		if s != stage {
			skipStages = append(skipStages, s)
		}
	}
	b.skipStages = skipStages
}

func (c *BuildManager) Init(args []string) {
	v := viper.GetViper()
	acPath := v.GetString(string(viper_keys.AppConfigPath))
	hasTargetDir := (len(args) > 0) && (args[0] != "")
	var initialStage build_stages.BuildStage
	if acPath != "" {
		c.SetConfigLocationStrategy(config_loc_strategies.CopyFromLocation, acPath)
	}
	if (v.GetBool(string(viper_keys.UseCwd))) || (hasTargetDir) {
		initialStage = utils.Ternary(acPath != "", build_stages.CloneTemplateAppStage, build_stages.ConfirmConfigLocFromEnv)
	} else if acPath != "" {
		initialStage = build_stages.CloneTemplateAppStage
	} else {
		initialStage = build_stages.ConfirmCurrentDirStage
	}
	c.SetInitialStage(initialStage)
}

// Utility to check if stage should be skipped. Reads data formerly set by AddSkippedStage.
func (b *BuildManager) ShouldSkipStage(stageId build_stages.BuildStage) bool {
	for _, s := range b.skipStages {
		if s == stageId {
			return true
		}
	}
	return false
}

// Sets the active stage without the ability to go backwards.
func ToPreviousStage() {
	if (len(b.stack) >= 1) && (b.allowGoBack) {
		b.stack = b.stack[0 : len(b.stack)-1]
	}
}

// Sets active build stage for global access.
func SetActiveStage(stageId build_stages.BuildStage) {
	b.stack = append(b.stack, stageId)
}

// Sets active build stage for global access.
func SetAppConfigPath(p string) {
	b.AppConfigPath = p
}

// Sets active build stage for global access.
func SetConfigDirPath(p string) {
	b.ConfigDirPath = p
}

// Checks if build stage is active build stage.
func (b *BuildManager) IsActiveStage(stageId build_stages.BuildStage) bool {
	return b.Stage() == stageId
}

func (b *BuildManager) Stage() build_stages.BuildStage {
	return b.stack[len(b.stack)-1]
}

func (b *BuildManager) Stack() []build_stages.BuildStage {
	return b.stack
}

func (b *BuildManager) SendToPreviousStageMsg() tea.Cmd {
	ToPreviousStage()
	return signals.SetStage(b.stack[len(b.stack)-1])
}

func (b *BuildManager) PluginConflicts() []ulld_plugin.PluginConflict {
	log.Fatal("Need to implement this!")
	var res []ulld_plugin.PluginConflict
	return res
}
