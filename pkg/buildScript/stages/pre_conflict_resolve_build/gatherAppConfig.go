package stage_pre_conflict_resolve_build

import (
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	build_config "github.com/Uh-little-less-dum/go-utils/pkg/config"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
)

func GatherAppConfig(cfg *build_config.BuildConfigOpts) {
	cfg.AppConfig = app_config.NewAppConfig(cfg.TargetDir)
}

func GetAppStage() *sub_stage.SubStage {
	return sub_stage.NewSubStage("gatherAppConfig", "Gathering app config", sub_command_ids.GatherAppConfig, GatherAppConfig)
}
