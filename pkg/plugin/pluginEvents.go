package ulld_plugin

import (
	"fmt"

	"github.com/charmbracelet/log"
)

type PluginEvents struct {
	pluginConfig *PluginConfig
}

type PluginSyncType string

const (
	OnSync    PluginSyncType = "onSync"
	OnBackup  PluginSyncType = "onBackup"
	OnRestore PluginSyncType = "onRestore"
	OnBuild   PluginSyncType = "onBuild"
)

func PluginSyncTypes() []PluginSyncType {
	return []PluginSyncType{
		OnSync,
		OnBackup,
		OnRestore,
		OnBuild,
	}
}

type PluginEventHandler struct {
	SyncType   PluginSyncType
	ImportName string
	PluginName string
	ImportPath string
}

func (p *PluginEvents) MustGetPluginName() string {
	pluginName := p.pluginConfig.data.Get("pluginName").Str
	if pluginName == "" {
		log.Fatal("Could not find a plugin name for a plugin whilie generating event handlers.")
	}
	return pluginName
}

func (p PluginEvents) OnBuildMethod(idx int) (handler PluginEventHandler, ok bool) {
	pn := p.MustGetPluginName()
	ip := p.pluginConfig.data.Get("events.onBuild").Str
	return PluginEventHandler{SyncType: OnBuild, ImportName: fmt.Sprintf("OnBuild_%d", idx), PluginName: pn, ImportPath: ip}, ip != ""
}

func (p PluginEvents) OnSyncMethod(idx int) (handler PluginEventHandler, ok bool) {
	pn := p.MustGetPluginName()
	ip := p.pluginConfig.data.Get("events.onSync").Str
	return PluginEventHandler{SyncType: OnSync, ImportName: fmt.Sprintf("OnSync_%d", idx), PluginName: pn, ImportPath: ip}, ip != ""
}

func (p PluginEvents) OnRestoreMethod(idx int) (handler PluginEventHandler, ok bool) {
	pn := p.MustGetPluginName()
	ip := p.pluginConfig.data.Get("events.onRestore").Str
	return PluginEventHandler{SyncType: OnRestore, ImportName: fmt.Sprintf("OnRestore_%d", idx), PluginName: pn, ImportPath: ip}, ip != ""

}

func (p PluginEvents) OnBackupMethod(idx int) (handler PluginEventHandler, ok bool) {
	pn := p.MustGetPluginName()
	ip := p.pluginConfig.data.Get("events.onBackup").Str
	return PluginEventHandler{SyncType: OnBackup, ImportName: fmt.Sprintf("OnBackup_%d", idx), PluginName: pn, ImportPath: ip}, ip != ""
}

func NewPluginEvents(pluginConfig *PluginConfig) *PluginEvents {
	return &PluginEvents{pluginConfig: pluginConfig}
}
