package types

type UlldPluginComponent interface {
	SlotKey() string
}

type UlldPlugin interface {
	Components() []*UlldPluginComponent
}

type OutputWriter interface {
	WriteOutput(cfg *BuildManager)
}
