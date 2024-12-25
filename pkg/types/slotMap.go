package types

import build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"

type UlldPluginComponent interface {
	SlotKey() string
}

type UlldPlugin interface {
	Components() []*UlldPluginComponent
}

type OutputWriter interface {
	WriteOutput(cfg *build_config.BuildManager)
}
