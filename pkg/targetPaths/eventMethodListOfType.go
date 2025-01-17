package target_paths

import (
	event_handler_types "github.com/Uh-little-less-dum/build/pkg/buildConstants/eventTypes"
	"github.com/charmbracelet/log"
)

func (t TargetPaths) EventMethodListOfType(eventType event_handler_types.PluginEventType) string {
	switch eventType {
	case event_handler_types.OnBuild:
		return t.OnBuildMethodList()
	case event_handler_types.OnSync:
		return t.OnSyncMethodList()
	case event_handler_types.OnBackup:
		return t.OnBackupMethodList()
	case event_handler_types.OnRestore:
		return t.OnRestoreMethodList()
	}
	log.Fatal("Missing eventType in EventMethodListOfType method.")
	return t.OnSyncMethodList()
}
