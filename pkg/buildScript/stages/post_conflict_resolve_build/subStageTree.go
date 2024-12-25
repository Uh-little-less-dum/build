package stage_post_conflict_resolve_build

import (
	post_conflict_resolve_stages "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/post_conflict_resolve_build/stages"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

// RESUME: Every one of these methods is completely un-implemented. Need to handle this first thing before getting back online and able to wrap this up.
func GetSubStageTree(program *tea.Program) []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		sub_stage.NewSubStage("Write plugin output", "Generating plugin outputs", sub_command_ids.WritePluginOutput, post_conflict_resolve_stages.WritePluginOutput, program),
		sub_stage.NewSubStage("Generated personalized paths", "Creating paths based on your configuration", sub_command_ids.WriteNoteTypeOutputs, post_conflict_resolve_stages.WriteNoteTypeData, program),
		sub_stage.NewSubStage("Write additional parsers", "Writing additional parsers", sub_command_ids.WriteAddtionalParsers, post_conflict_resolve_stages.WriteAdditionalParserLists, program),
		sub_stage.NewSubStage("Write embeddable components", "Writing embeddable components", sub_command_ids.WriteComponentMap, post_conflict_resolve_stages.WriteComponentMap, program),
		// Begin event handler stuffs
		sub_stage.NewSubStage("Write onSync event handlers", "Writing onSync event handlers", sub_command_ids.WriteOnSyncEventHandlers, post_conflict_resolve_stages.WriteOnSyncHandlers, program),
		sub_stage.NewSubStage("Write onRestore event handlers", "Writing onRestore event handlers", sub_command_ids.WriteOnRestoreEventHandlers, post_conflict_resolve_stages.WriteOnRestoreHandlers, program),
		sub_stage.NewSubStage("Write onBackup event handlers", "Writing onBackup event handlers", sub_command_ids.WriteOnBackupEventHandlers, post_conflict_resolve_stages.WriteOnBackupHandlers, program),
		sub_stage.NewSubStage("Write onBuild event handlers", "Writing onBuild event handlers", sub_command_ids.WriteOnBuildEventHandlers, post_conflict_resolve_stages.WriteOnBuildHandlers, program),
		// End event handler stuffs
		sub_stage.NewSubStage("Generate database", "Generating database tables", sub_command_ids.GenerateDatabase, post_conflict_resolve_stages.GenerateDatabase, program),
		sub_stage.NewSubStage("Write build output data", "Writing buildi output data", sub_command_ids.WriteBuildDataOutput, post_conflict_resolve_stages.WriteBuildStaticData, program),
	}
}
