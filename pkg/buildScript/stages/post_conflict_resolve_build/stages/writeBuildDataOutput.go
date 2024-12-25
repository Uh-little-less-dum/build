package post_conflict_resolve_stages

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	build_static_data "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/buildStaticData"
)

// Writes the builidStaticData.json file.
func WriteBuildStaticData(cfg *build_config.BuildManager) {
	b := build_static_data.NewBuildStaticData(cfg)
	b.WriteOutput()
}
