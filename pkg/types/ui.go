package types

import (
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	tea "github.com/charmbracelet/bubbletea"
)

type InternalModel interface {
	Stage() build_stages.BuildStage
	View() string
	Init() tea.Cmd
	Update(tea.Msg) (InternalModel, tea.Cmd)
}
