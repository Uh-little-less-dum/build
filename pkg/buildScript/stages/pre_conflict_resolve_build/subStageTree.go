package stage_pre_conflict_resolve_build

import (
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
)

func GetSubStageTree() []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		sub_stage.NewSubStage("Gather App Config", "Taking a look at your config...", sub_command_ids.GatherAppConfig, GatherAppConfig),
		sub_stage.NewSubStage("Gather root package.json", "Grabbing some data from your package.json file...", sub_command_ids.GatherAppConfig, GatherRootPackageJson),
		sub_stage.NewSubStage("Gather Plugins", "Gathering your plugins...", sub_command_ids.GatherPlugins, GatherPlugins),
	}
}
