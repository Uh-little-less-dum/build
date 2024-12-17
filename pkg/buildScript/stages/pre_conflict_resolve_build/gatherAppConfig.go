package stage_pre_conflict_resolve_build

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
)

func GatherAppConfig(cfg *build_config.BuildManager) {
	cfg.SetAppConfig(app_config.NewAppConfig(cfg.AppConfigPath))
}
