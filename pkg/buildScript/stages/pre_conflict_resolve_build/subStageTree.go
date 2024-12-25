package stage_pre_conflict_resolve_build

import (
	stage_write_npmrc "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/pre_conflict_resolve_build/writeNpmrc"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

func GetSubStageTree(program *tea.Program) []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		sub_stage.NewSubStage("Gather App Config", "Taking a look at your config", sub_command_ids.GatherAppConfig, GatherAppConfig, program),
		sub_stage.NewSubStage("Gather root package.json", "Grabbing some data from your package.json file", sub_command_ids.GatherRootPackageJson, GatherRootPackageJson, program),
		sub_stage.NewSubStage("Gather Plugins", "Gathering your plugins", sub_command_ids.GatherPlugins, GatherPlugins, program),
		sub_stage.NewSubStage("Writing package manager configuration", "Writing package manager configuration", sub_command_ids.WriteNpmrc, stage_write_npmrc.WriteNpmrc, program),
		sub_stage.NewSubStage("Install Dependencies", "Installing dependencies. This might take a bit.", sub_command_ids.InstallDependencies, InstallDependencies, program),
		sub_stage.NewSubStage("Gather plugin conflicts", "Checking for any conflicts amongst your plugins", sub_command_ids.GatherPluginConflicts, GatherPluginConflicts, program),
	}
}
