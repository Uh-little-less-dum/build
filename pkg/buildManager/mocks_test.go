// Code generated by MockGen. DO NOT EDIT.
// Source: buildConfig_interface_test.go

// Package build_config_test is a generated GoMock package.
package build_config_test

import (
	reflect "reflect"

	config_loc_strategies "github.com/Uh-little-less-dum/build/pkg/buildConstants/configLocationStrategies"
	navigation_link "github.com/Uh-little-less-dum/build/pkg/buildManager/auxillaryStructs/navigationLink"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	types "github.com/Uh-little-less-dum/build/pkg/types"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	copy_file "github.com/Uh-little-less-dum/go-utils/pkg/fs/copyFile"
	tea "github.com/charmbracelet/bubbletea"
	gomock "github.com/golang/mock/gomock"
)

// MockBuildManagerInterface is a mock of BuildManagerInterface interface.
type MockBuildManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBuildManagerInterfaceMockRecorder
}

// MockBuildManagerInterfaceMockRecorder is the mock recorder for MockBuildManagerInterface.
type MockBuildManagerInterfaceMockRecorder struct {
	mock *MockBuildManagerInterface
}

// NewMockBuildManagerInterface creates a new mock instance.
func NewMockBuildManagerInterface(ctrl *gomock.Controller) *MockBuildManagerInterface {
	mock := &MockBuildManagerInterface{ctrl: ctrl}
	mock.recorder = &MockBuildManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildManagerInterface) EXPECT() *MockBuildManagerInterfaceMockRecorder {
	return m.recorder
}

// AddSkippedStage mocks base method.
func (m *MockBuildManagerInterface) AddSkippedStage(stage build_stages.BuildStage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSkippedStage", stage)
}

// AddSkippedStage indicates an expected call of AddSkippedStage.
func (mr *MockBuildManagerInterfaceMockRecorder) AddSkippedStage(stage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSkippedStage", reflect.TypeOf((*MockBuildManagerInterface)(nil).AddSkippedStage), stage)
}

// AppConfig mocks base method.
func (m *MockBuildManagerInterface) AppConfig() *app_config.AppConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppConfig")
	ret0, _ := ret[0].(*app_config.AppConfig)
	return ret0
}

// AppConfig indicates an expected call of AppConfig.
func (mr *MockBuildManagerInterfaceMockRecorder) AppConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppConfig", reflect.TypeOf((*MockBuildManagerInterface)(nil).AppConfig))
}

// ComponentDocsData mocks base method.
func (m *MockBuildManagerInterface) ComponentDocsData() []ulld_plugin.ComponenDocData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentDocsData")
	ret0, _ := ret[0].([]ulld_plugin.ComponenDocData)
	return ret0
}

// ComponentDocsData indicates an expected call of ComponentDocsData.
func (mr *MockBuildManagerInterfaceMockRecorder) ComponentDocsData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentDocsData", reflect.TypeOf((*MockBuildManagerInterface)(nil).ComponentDocsData))
}

// Embeddables mocks base method.
func (m *MockBuildManagerInterface) Embeddables() []ulld_plugin.PluginEmbeddableTemplateStruct {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Embeddables")
	ret0, _ := ret[0].([]ulld_plugin.PluginEmbeddableTemplateStruct)
	return ret0
}

// Embeddables indicates an expected call of Embeddables.
func (mr *MockBuildManagerInterfaceMockRecorder) Embeddables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Embeddables", reflect.TypeOf((*MockBuildManagerInterface)(nil).Embeddables))
}

// GatherPluginConflicts mocks base method.
func (m *MockBuildManagerInterface) GatherPluginConflicts() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GatherPluginConflicts")
}

// GatherPluginConflicts indicates an expected call of GatherPluginConflicts.
func (mr *MockBuildManagerInterfaceMockRecorder) GatherPluginConflicts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatherPluginConflicts", reflect.TypeOf((*MockBuildManagerInterface)(nil).GatherPluginConflicts))
}

// GetCopyStyleSheets mocks base method.
func (m *MockBuildManagerInterface) GetCopyStyleSheets() []*copy_file.CopyFile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCopyStyleSheets")
	ret0, _ := ret[0].([]*copy_file.CopyFile)
	return ret0
}

// GetCopyStyleSheets indicates an expected call of GetCopyStyleSheets.
func (mr *MockBuildManagerInterfaceMockRecorder) GetCopyStyleSheets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCopyStyleSheets", reflect.TypeOf((*MockBuildManagerInterface)(nil).GetCopyStyleSheets))
}

// GetNavigationLinks mocks base method.
func (m *MockBuildManagerInterface) GetNavigationLinks() []navigation_link.NavigationLink {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNavigationLinks")
	ret0, _ := ret[0].([]navigation_link.NavigationLink)
	return ret0
}

// GetNavigationLinks indicates an expected call of GetNavigationLinks.
func (mr *MockBuildManagerInterfaceMockRecorder) GetNavigationLinks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNavigationLinks", reflect.TypeOf((*MockBuildManagerInterface)(nil).GetNavigationLinks))
}

// Init mocks base method.
func (m *MockBuildManagerInterface) Init(args []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", args)
}

// Init indicates an expected call of Init.
func (mr *MockBuildManagerInterfaceMockRecorder) Init(args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockBuildManagerInterface)(nil).Init), args)
}

// Installables mocks base method.
func (m *MockBuildManagerInterface) Installables() []types.Installable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Installables")
	ret0, _ := ret[0].([]types.Installable)
	return ret0
}

// Installables indicates an expected call of Installables.
func (mr *MockBuildManagerInterfaceMockRecorder) Installables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Installables", reflect.TypeOf((*MockBuildManagerInterface)(nil).Installables))
}

// IsActiveStage mocks base method.
func (m *MockBuildManagerInterface) IsActiveStage(stageId build_stages.BuildStage) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActiveStage", stageId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActiveStage indicates an expected call of IsActiveStage.
func (mr *MockBuildManagerInterfaceMockRecorder) IsActiveStage(stageId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActiveStage", reflect.TypeOf((*MockBuildManagerInterface)(nil).IsActiveStage), stageId)
}

// PackageManager mocks base method.
func (m *MockBuildManagerInterface) PackageManager() package_managers.PackageManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackageManager")
	ret0, _ := ret[0].(package_managers.PackageManager)
	return ret0
}

// PackageManager indicates an expected call of PackageManager.
func (mr *MockBuildManagerInterfaceMockRecorder) PackageManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackageManager", reflect.TypeOf((*MockBuildManagerInterface)(nil).PackageManager))
}

// PluginsToTranspile mocks base method.
func (m *MockBuildManagerInterface) PluginsToTranspile() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluginsToTranspile")
	ret0, _ := ret[0].([]string)
	return ret0
}

// PluginsToTranspile indicates an expected call of PluginsToTranspile.
func (mr *MockBuildManagerInterfaceMockRecorder) PluginsToTranspile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluginsToTranspile", reflect.TypeOf((*MockBuildManagerInterface)(nil).PluginsToTranspile))
}

// RemoveStageFromSkipped mocks base method.
func (m *MockBuildManagerInterface) RemoveStageFromSkipped(stage build_stages.BuildStage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveStageFromSkipped", stage)
}

// RemoveStageFromSkipped indicates an expected call of RemoveStageFromSkipped.
func (mr *MockBuildManagerInterfaceMockRecorder) RemoveStageFromSkipped(stage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStageFromSkipped", reflect.TypeOf((*MockBuildManagerInterface)(nil).RemoveStageFromSkipped), stage)
}

// SendToPreviousStageMsg mocks base method.
func (m *MockBuildManagerInterface) SendToPreviousStageMsg() tea.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendToPreviousStageMsg")
	ret0, _ := ret[0].(tea.Cmd)
	return ret0
}

// SendToPreviousStageMsg indicates an expected call of SendToPreviousStageMsg.
func (mr *MockBuildManagerInterfaceMockRecorder) SendToPreviousStageMsg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendToPreviousStageMsg", reflect.TypeOf((*MockBuildManagerInterface)(nil).SendToPreviousStageMsg))
}

// SetAppConfig mocks base method.
func (m *MockBuildManagerInterface) SetAppConfig(ac *app_config.AppConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAppConfig", ac)
}

// SetAppConfig indicates an expected call of SetAppConfig.
func (mr *MockBuildManagerInterfaceMockRecorder) SetAppConfig(ac interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAppConfig", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetAppConfig), ac)
}

// SetAppConfigPath mocks base method.
func (m *MockBuildManagerInterface) SetAppConfigPath(p string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAppConfigPath", p)
}

// SetAppConfigPath indicates an expected call of SetAppConfigPath.
func (mr *MockBuildManagerInterfaceMockRecorder) SetAppConfigPath(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAppConfigPath", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetAppConfigPath), p)
}

// SetConfigLocationStrategy mocks base method.
func (m *MockBuildManagerInterface) SetConfigLocationStrategy(strategyType config_loc_strategies.ConfigLocationStrategyId, value ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{strategyType}
	for _, a := range value {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetConfigLocationStrategy", varargs...)
}

// SetConfigLocationStrategy indicates an expected call of SetConfigLocationStrategy.
func (mr *MockBuildManagerInterfaceMockRecorder) SetConfigLocationStrategy(strategyType interface{}, value ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{strategyType}, value...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigLocationStrategy", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetConfigLocationStrategy), varargs...)
}

// SetInitialStage mocks base method.
func (m *MockBuildManagerInterface) SetInitialStage(s build_stages.BuildStage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInitialStage", s)
}

// SetInitialStage indicates an expected call of SetInitialStage.
func (mr *MockBuildManagerInterfaceMockRecorder) SetInitialStage(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitialStage", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetInitialStage), s)
}

// SetPackageManager mocks base method.
func (m *MockBuildManagerInterface) SetPackageManager(p package_managers.PackageManagerId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPackageManager", p)
}

// SetPackageManager indicates an expected call of SetPackageManager.
func (mr *MockBuildManagerInterfaceMockRecorder) SetPackageManager(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPackageManager", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetPackageManager), p)
}

// SetPlugins mocks base method.
func (m *MockBuildManagerInterface) SetPlugins(p []*ulld_plugin.Plugin) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPlugins", p)
}

// SetPlugins indicates an expected call of SetPlugins.
func (mr *MockBuildManagerInterfaceMockRecorder) SetPlugins(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlugins", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetPlugins), p)
}

// SetRootPackageJson mocks base method.
func (m *MockBuildManagerInterface) SetRootPackageJson(p file_handlers_package_json.PackageJsonHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRootPackageJson", p)
}

// SetRootPackageJson indicates an expected call of SetRootPackageJson.
func (mr *MockBuildManagerInterfaceMockRecorder) SetRootPackageJson(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRootPackageJson", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetRootPackageJson), p)
}

// SetTargetDir mocks base method.
func (m *MockBuildManagerInterface) SetTargetDir(d string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTargetDir", d)
}

// SetTargetDir indicates an expected call of SetTargetDir.
func (mr *MockBuildManagerInterfaceMockRecorder) SetTargetDir(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTargetDir", reflect.TypeOf((*MockBuildManagerInterface)(nil).SetTargetDir), d)
}

// SettingPageData mocks base method.
func (m *MockBuildManagerInterface) SettingPageData() []ulld_plugin.SettingsPageData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettingPageData")
	ret0, _ := ret[0].([]ulld_plugin.SettingsPageData)
	return ret0
}

// SettingPageData indicates an expected call of SettingPageData.
func (mr *MockBuildManagerInterfaceMockRecorder) SettingPageData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettingPageData", reflect.TypeOf((*MockBuildManagerInterface)(nil).SettingPageData))
}

// ShouldSkipStage mocks base method.
func (m *MockBuildManagerInterface) ShouldSkipStage(stageId build_stages.BuildStage) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldSkipStage", stageId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldSkipStage indicates an expected call of ShouldSkipStage.
func (mr *MockBuildManagerInterfaceMockRecorder) ShouldSkipStage(stageId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldSkipStage", reflect.TypeOf((*MockBuildManagerInterface)(nil).ShouldSkipStage), stageId)
}

// Stack mocks base method.
func (m *MockBuildManagerInterface) Stack() []build_stages.BuildStage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stack")
	ret0, _ := ret[0].([]build_stages.BuildStage)
	return ret0
}

// Stack indicates an expected call of Stack.
func (mr *MockBuildManagerInterfaceMockRecorder) Stack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stack", reflect.TypeOf((*MockBuildManagerInterface)(nil).Stack))
}

// Stage mocks base method.
func (m *MockBuildManagerInterface) Stage() build_stages.BuildStage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stage")
	ret0, _ := ret[0].(build_stages.BuildStage)
	return ret0
}

// Stage indicates an expected call of Stage.
func (mr *MockBuildManagerInterfaceMockRecorder) Stage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stage", reflect.TypeOf((*MockBuildManagerInterface)(nil).Stage))
}
