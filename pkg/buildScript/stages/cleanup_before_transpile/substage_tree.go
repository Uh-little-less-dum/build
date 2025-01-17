package stage_cleanup_before_transpile

import (
	cleanup_before_transpile_stages "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/cleanup_before_transpile/stages"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

func GetSubStageTree(program *tea.Program) []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		sub_stage.NewSubStage("Cleaning up git", "Cleaning up version control", sub_command_ids.CleanupGit, cleanup_before_transpile_stages.CleanupGit, program),
	}
}
