package stage_post_conflict_resolve_build

import (
	post_conflict_resolve_stages "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/post_conflict_resolve_build/stages"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

func GetSubStageTree(program *tea.Program) []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		// Config based output
		sub_stage.NewSubStage("Generated personalized paths", "Creating paths based on your configuration", sub_command_ids.WriteNoteTypeOutputs, post_conflict_resolve_stages.WriteNoteTypeData, program),
		// Plugin based output
		sub_stage.NewSubStage("Write style sheets", "Generating style sheets from your provided config", sub_command_ids.WriteScssOutput, post_conflict_resolve_stages.WriteStyleSheets, program),
		sub_stage.NewSubStage("Write plugin generate output", "Generating plugin pages", sub_command_ids.WritePluginOutput, post_conflict_resolve_stages.WritePluginOutput, program),
		sub_stage.NewSubStage("Write additional parsers", "Writing additional parsers", sub_command_ids.WriteAdditionalParsers, post_conflict_resolve_stages.WriteAdditionalParserLists, program),
		// Consolidated plugin output (accumulative across all plugins)
		sub_stage.NewSubStage("Write embeddable components", "Writing embeddable components", sub_command_ids.WriteComponentMap, post_conflict_resolve_stages.WriteComponentMap, program),
		sub_stage.NewSubStage("Write event handlers", "Writing event handlers", sub_command_ids.WriteEventHandlers, post_conflict_resolve_stages.WriteEventHandlers, program),
		// Database stuff
		sub_stage.NewSubStage("Create database schema", "Generating database schema files", sub_command_ids.WriteDatabaseSchemaFile, post_conflict_resolve_stages.WriteDatabaseSchemaFile, program),
		sub_stage.NewSubStage("Generate database", "Generating database tables", sub_command_ids.GenerateDatabase, post_conflict_resolve_stages.GenerateDatabase, program),
		sub_stage.NewSubStage("Write build output data", "Writing build output data", sub_command_ids.WriteBuildDataOutput, post_conflict_resolve_stages.WriteBuildStaticData, program),
	}
}
