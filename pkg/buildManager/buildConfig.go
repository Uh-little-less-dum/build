package build_config

import (
	"os"
	"time"

	form_data "github.com/Uh-little-less-dum/build/pkg/buildManager/formData"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/classesKinda/plugin"
	env_vars "github.com/Uh-little-less-dum/build/pkg/envVars"
	"github.com/Uh-little-less-dum/build/pkg/utils"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type BuildManager struct {
	form_data.BuildFormData
	AppConfigPath   string
	ConfigDirPath   string
	LogLevel        string
	RootPackageJson file_handlers_package_json.PackageJsonHandler
	Plugins         []ulld_plugin.Plugin
	LogFile         string
	stack           []build_stages.BuildStage
	skipStages      []build_stages.BuildStage
	allowGoBack     bool
	appConfig       *app_config.AppConfig
}

var b *BuildManager

// Returns the BuildManager singleton.
func GetBuildManager() *BuildManager {
	return b
}

func (b *BuildManager) AppConfig() *app_config.AppConfig {
	return b.appConfig
}

func (b *BuildManager) SetAppConfig(ac *app_config.AppConfig) {
	b.appConfig = ac
}

func (b *BuildManager) SetPlugins(p []ulld_plugin.Plugin) {
	b.Plugins = p
}

func (b *BuildManager) SetRootPackageJson(p file_handlers_package_json.PackageJsonHandler) {
	b.RootPackageJson = p
}

func init() {
	val := BuildManager{
		AppConfigPath: utils.GetDefaultAppConfigPath(),
	}

	if val.TargetDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		val.TargetDir = cwd
	}
	// Handle initial values with environment variables here.
	val.ConfigDirPath = os.Getenv(string(env_vars.AdditionalSources))
	val.LogLevel = os.Getenv(string(env_vars.LogLevel))
	val.LogFile = os.Getenv(string(env_vars.LogFile))
	// FIX: This was changed for development only! Remove this once this model's working properly.
	val.stack = []build_stages.BuildStage{build_stages.ConfirmCurrentDirStage}
	b = &val
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

// TODO: Set up the rest of the config flags here. Handle what you can with viper first though. Might not even need this struct if viper can handle everything.
func (c *BuildManager) Init(args []string) {
	v := viper.GetViper()
	acPath := v.GetString(string(viper_keys.AppConfigPath))
	hasTargetDir := (len(args) > 0) && (args[0] != "")
	var initialStage build_stages.BuildStage
	if (v.GetBool(string(viper_keys.UseCwd))) || (hasTargetDir) {
		initialStage = utils.Ternary(acPath != "", build_stages.CloneTemplateAppStage, build_stages.ConfirmConfigLocFromEnv)
	} else if acPath != "" {
		initialStage = build_stages.CloneTemplateAppStage
	} else {
		initialStage = build_stages.ConfirmCurrentDirStage
	}
	// FIX: This is temporary, for development only!!! Remove this once the stream model is working well.
	if os.Getenv("WAIT_FOR_DEBUG") != "" {
		time.Sleep(15 * time.Second)
	}
	initialStage = build_stages.PreConflictResolveBuild
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

func (b *BuildManager) SetTargetDir(targetDir string) {
	b.TargetDir = targetDir
}

func (b *BuildManager) SendToPreviousStageMsg() tea.Cmd {
	ToPreviousStage()
	return signals.SetStage(b.stack[len(b.stack)-1])
}
