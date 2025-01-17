package event_handler_types

type PluginEventType string

const (
    OnBuild    PluginEventType = "onBuild"
    OnSync    PluginEventType = "onSync"
    OnBackup    PluginEventType = "onBackup"
    OnRestore    PluginEventType = "onRestore"
)

func PluginSyncTypes() []PluginEventType {
	return []PluginEventType{
        OnBuild,
        OnSync,
        OnBackup,
        OnRestore,
         }
}

