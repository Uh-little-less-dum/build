package stage_pre_conflict_resolve_build

import build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"

func GatherPluginConflicts(cfg *build_config.BuildManager) {
	cfg.GatherPluginConflicts()
}
