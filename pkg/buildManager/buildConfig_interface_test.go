// DONT EDIT: Auto generated

package build_config_test

import (
	config_loc_strategies "github.com/Uh-little-less-dum/build/pkg/buildConstants/configLocationStrategies"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	plugin_component_docs_data "github.com/Uh-little-less-dum/build/pkg/plugin/componentDocs"
	navigation_link "github.com/Uh-little-less-dum/build/pkg/plugin/navigationLink"
	plugin_setting_page_data "github.com/Uh-little-less-dum/build/pkg/plugin/settingPageData"
	"github.com/Uh-little-less-dum/build/pkg/types"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	copy_file "github.com/Uh-little-less-dum/go-utils/pkg/fs/copyFile"
	tea "github.com/charmbracelet/bubbletea"
)

// HumanIface makes human interaction easy
type BuildManagerInterface interface {
	SetTargetDir(d string)
	AppConfig() *app_config.AppConfig
	SetAppConfigPath(p string)
	SetPackageManager(p package_managers.PackageManagerId)
	GetCopyStyleSheets() []*copy_file.CopyFile
	PackageManager() package_managers.PackageManager
	Installables() []types.Installable
	SetAppConfig(ac *app_config.AppConfig)
	SetPlugins(p []*ulld_plugin.Plugin)
	SetRootPackageJson(p file_handlers_package_json.PackageJsonHandler)
	SetConfigLocationStrategy(strategyType config_loc_strategies.ConfigLocationStrategyId, value ...string)
	SetInitialStage(s build_stages.BuildStage)
	// Adds a build stage that should be skipped to the list of skipped stages. This list can then be accessed from within each build stage at which point the stage can be bypassed.
	AddSkippedStage(stage build_stages.BuildStage)
	RemoveStageFromSkipped(stage build_stages.BuildStage)
	Init(args []string)
	// Utility to check if stage should be skipped. Reads data formerly set by AddSkippedStage.
	ShouldSkipStage(stageId build_stages.BuildStage) bool
	// Checks if build stage is active build stage.
	IsActiveStage(stageId build_stages.BuildStage) bool
	Stage() build_stages.BuildStage
	Stack() []build_stages.BuildStage
	SendToPreviousStageMsg() tea.Cmd
	GatherPluginConflicts()
	Embeddables() []ulld_plugin.PluginEmbeddableTemplateStruct
	// Returns an array of all pluginNames that should be transpiled.
	PluginsToTranspile() []string
	ComponentDocsData() []plugin_component_docs_data.ComponenDocData
	// Returns an array of items matching the buildStaticData.settingPages.# field.
	SettingPageData() []plugin_setting_page_data.SettingsPageData
	// FIX: Implement this.
	GetNavigationLinks() []navigation_link.NavigationLink
}
