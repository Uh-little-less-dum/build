package stage_pre_conflict_resolve_build

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/classesKinda/plugin"
)

func GatherPlugins(cfg *build_config.BuildManager) {
	data := []ulld_plugin.Plugin{}
	cfg.SetPlugins(data)
}
