package stage_pre_conflict_resolve_build

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
)

func GatherPlugins(cfg *build_config.BuildManager) {
	data := cfg.AppConfig().GatherPlugins(cfg.Paths)
	cfg.SetPlugins(data)
}
