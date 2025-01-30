package build_config

import (
	"embed"
	"fmt"
	"os"
	"sync"

	additional_sources_manager "github.com/Uh-little-less-dum/build/pkg/additionalSourcesManager"
	config_loc_strategies "github.com/Uh-little-less-dum/build/pkg/buildConstants/configLocationStrategies"
	event_handler_types "github.com/Uh-little-less-dum/build/pkg/buildConstants/eventTypes"
	form_data "github.com/Uh-little-less-dum/build/pkg/buildManager/formData"
	conflicts_handler "github.com/Uh-little-less-dum/build/pkg/conflicts/conflictsManager"
	database_manager "github.com/Uh-little-less-dum/build/pkg/databaseManager"
	env_vars "github.com/Uh-little-less-dum/build/pkg/envVars"
	event_handlers "github.com/Uh-little-less-dum/build/pkg/eventHandlers"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	plugin_component_docs_data "github.com/Uh-little-less-dum/build/pkg/plugin/componentDocs"
	navigation_link "github.com/Uh-little-less-dum/build/pkg/plugin/navigationLink"
	plugin_setting_page_data "github.com/Uh-little-less-dum/build/pkg/plugin/settingPageData"
	styles_manager "github.com/Uh-little-less-dum/build/pkg/stylesManager"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/Uh-little-less-dum/build/pkg/utils"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	build_static_data_output "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/buildStaticData/outputStruct"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	copy_file "github.com/Uh-little-less-dum/go-utils/pkg/fs/copyFile"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var (
	//go:embed "staticData/*"
	staticFiles embed.FS
)

//go:generate ifacemaker -f buildConfig.go -s BuildManager -i BuildManagerInterface -p build_config_test -y "HumanIface makes human interaction easy" -c "DONT EDIT: Auto generated" -o buildConfig_interface_test.go
//go:generate mockgen -package build_config_test -source=buildConfig_interface_test.go -destination=mocks_test.go *
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
	Conflicts            conflicts_handler.BuildConflictsManager
	Program              *tea.Program
	Paths                *target_paths.TargetPaths
	Db                   *database_manager.DatabaseManager
	AdditionalSources    *additional_sources_manager.AdditionalSourcesManager
	Styles               *styles_manager.StylesManager
}

var b *BuildManager

var once sync.Once

// Returns the BuildManager singleton.
func GetBuildManager() *BuildManager {
	once.Do(func() {
		b = GetInitialBuildManager()
	})
	return b
}

func (b *BuildManager) SlotMapData() gjson.Result {
	f, err := staticFiles.ReadFile("staticData/slotData.json")
	if err != nil {
		log.Fatal(err)
	}
	return gjson.ParseBytes(f)
}

func (b *BuildManager) SetTargetDir(d string) {
	bp := afero.NewBasePathFs(afero.NewOsFs(), d)
	b.Fs = &bp
	b.Paths = target_paths.NewTargetPaths(d)
	b.Db.SetRootPath(b.Paths)
	b.BuildFormData.SetTargetDirOnlyInBuildConfig(d)
}

func (b *BuildManager) AppConfig() *app_config.AppConfig {
	return b.appConfig
}

func (b *BuildManager) SetAppConfigPath(p string) {
	b.AppConfigPath = p
	b.appConfig.SetPath(p)
}

func (b *BuildManager) SetPackageManager(p package_managers.PackageManagerId) {
	b.packageManager = p
}

func (b *BuildManager) GetCopyStyleSheets() []*copy_file.CopyFile {
	var res []*copy_file.CopyFile = b.AdditionalSources.StylePaths(b.Paths)
	for _, p := range b.Plugins {
		c, ok := p.CopyFileSheet(b.Paths)
		if ok {
			res = append(res, c)
		}
	}
	return res
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
		Paths:         target_paths.NewTargetPaths(ac),
		appConfig:     app_config.NewAppConfig(ac),
	}
	val.Db = database_manager.NewDatabaseManager(val.Paths)

	if val.BuildFormData.TargetDir() == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		val.SetTargetDir(cwd)
	}
	val.Styles = styles_manager.NewStylesManager()
	val.AdditionalSources = additional_sources_manager.NewAdditionalSourcesManager()
	val.packageManager = package_managers.PnpmId
	// Handle initial values with environment variables here.
	val.ConfigDirPath = os.Getenv(string(env_vars.AdditionalSources))
	val.LogLevel = os.Getenv(string(env_vars.LogLevel))
	val.LogFile = os.Getenv(string(env_vars.LogFile))
	val.stack = []build_stages.BuildStage{build_stages.ConfirmCurrentDirStage}
	return &val
}

// func init() {
// 	b = GetInitialBuildManager()
// }

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

func (c *BuildManager) SetInitialTargetDir(useCwd bool, targetDir string) {
	if targetDir != "" {
		c.SetTargetDir(targetDir)
	} else {

	}
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
	// Set initial target directory.
	if hasTargetDir {
		c.SetTargetDir(args[0])
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		c.SetTargetDir(cwd)
	}
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

func (b *BuildManager) GatherPluginConflicts() {
	b.Conflicts.GatherPluginConflicts(b.Plugins)
}

func (b *BuildManager) Embeddables() []ulld_plugin.PluginEmbeddableTemplateStruct {
	var embeddables []ulld_plugin.PluginEmbeddableTemplateStruct
	for _, p := range b.Plugins {
		for i, c := range p.Components() {
			embeddables = append(embeddables, c.Embeddables(p.Name, c.ComponentName(), c.ExportPath(), i)...)
		}
	}
	return embeddables
}

func (b *BuildManager) TailwindSources(outputStruct *build_static_data_output.BuildStaticDataOutput) {
	ch := make(chan string)
	var wg sync.WaitGroup
	for _, p := range b.Plugins {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := p.Config().Get("tailwind.sources")
			if data.Exists() {
				items := data.Array()
				for _, item := range items {
					ch <- item.Str
				}
			}
		}()
	}
	wg.Wait()
	close(ch)
	outputStruct.TailwindSources = make([]string, len(ch))
	for item := range ch {
		outputStruct.TailwindSources = append(outputStruct.TailwindSources, item)
	}
}

// Returns an array of all pluginNames that should be transpiled.
func (b *BuildManager) PluginsToTranspile(outputStruct *build_static_data_output.BuildStaticDataOutput) {
	ch := make(chan string)
	var wg sync.WaitGroup
	for _, p := range b.Plugins {
		wg.Add(1)
		go func() {
			if p.ShouldTranspile() {
				ch <- p.Name
			}
		}()
	}
	wg.Wait()
	close(ch)
	outputStruct.TranspilePackages = make([]string, len(ch))
	for l := range ch {
		outputStruct.TranspilePackages = append(outputStruct.TranspilePackages, l)
	}
}

func (b *BuildManager) EventHandlerListOfType(eventType event_handler_types.PluginEventType) event_handlers.EventHandlerOfTypeList {
	l := event_handlers.NewEventHandlerOfTypeList(eventType)
	for _, p := range b.Plugins {
		data := p.Config()
		val := data.Get(fmt.Sprintf("events.%s", string(eventType)))
		if val.Exists() {
			l.Append(p.Name, val)
		}
	}
	return *l
}

func (b *BuildManager) ComponentDocsData(outputStruct *build_static_data_output.BuildStaticDataOutput) {
	var wg sync.WaitGroup
	ch := make(chan plugin_component_docs_data.ComponenDocData)
	for _, p := range b.Plugins {
		for _, c := range p.Components() {
			wg.Add(1)
			go func() {
				data, ok := c.GetComponentDocData(p.Name)
				if ok {
					ch <- data
				}
			}()
		}
	}
	wg.Wait()
	n := len(ch)
	outputStruct.ComponentDocs = make([]plugin_component_docs_data.ComponenDocData, n)
	for item := range ch {
		outputStruct.ComponentDocs = append(outputStruct.ComponentDocs, item)
	}
}

// Returns an array of items matching the buildStaticData.settingPages.# field.
func (b *BuildManager) SettingPageData(outputStruct *build_static_data_output.BuildStaticDataOutput) {
	var wg sync.WaitGroup
	ch := make(chan plugin_setting_page_data.SettingsPageData)
	for _, p := range b.Plugins {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, ok := p.SettingsPageData()
			if ok {
				ch <- data
			}
		}()
	}
	wg.Wait()
}

func (b *BuildManager) GetNavigationLinks(outputStruct *build_static_data_output.BuildStaticDataOutput) {
	var wg sync.WaitGroup
	ch := make(chan navigation_link.NavigationLink)
	ch <- navigation_link.NavigationLink{Label: "Settings", Href: "/settings", Icon: "cog"}
	for _, p := range b.Plugins {
		wg.Add(1)
		go func() {
			p.NavigationLinks(ch)
			defer wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	outputStruct.NavigationLinks = make([]navigation_link.NavigationLink, len(ch))
	for navLink := range ch {
		outputStruct.NavigationLinks = append(outputStruct.NavigationLinks, navLink)
	}
}
