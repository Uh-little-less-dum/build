package ulld_plugin

import slot_conflict "github.com/Uh-little-less-dum/build/pkg/classesKinda/slotConflict"

type PluginConflict struct {
	Plugins  []Plugin
	Conflict slot_conflict.SlotConflict
}
