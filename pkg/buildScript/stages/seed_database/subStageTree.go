package stage_seed_database

import (
	seed_database_substages "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/seed_database/stages"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

func GetSubStageTree(program *tea.Program) []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		sub_stage.NewSubStage("Sync documentation indices", "Generating documentation data", sub_command_ids.SeedComponentDocs, seed_database_substages.SeedComponentDocumentationData, program),
	}
}
