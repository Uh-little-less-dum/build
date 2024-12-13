package ulld_plugin

type SlotData struct {
	ParentSlot string
	SubSlot    string
}

type Plugin struct {
	PluginName string
	Slot       SlotData // Not yet implemented. Will always be null until this is populated.
}
