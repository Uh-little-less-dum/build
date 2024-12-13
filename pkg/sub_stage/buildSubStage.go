package sub_stage

import (
	"fmt"

	build_config "github.com/Uh-little-less-dum/go-utils/pkg/config"
	run_status "github.com/Uh-little-less-dum/go-utils/pkg/constants/runStatus"
	stream_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/streamIds"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SubStage struct {
	Name          string
	id            sub_command_ids.SubCommandId
	status        run_status.RunStatus
	inProgressMsg string
	onRun         func(cfg *build_config.BuildConfigOpts)
}

func (b *SubStage) Run(cfg *build_config.BuildConfigOpts, streamId stream_ids.StreamId) tea.Cmd {
	b.status = run_status.Running
	b.onRun(cfg)
	return b.CompleteMsg(streamId)
}

func (b SubStage) Id() sub_command_ids.SubCommandId {
	return b.id
}

func (b SubStage) CompleteMsg(streamId stream_ids.StreamId) tea.Cmd {
	return signals.SendSubCommandCompleteMsg(streamId, b.id)
}

func (b SubStage) CompleteUserMessage() string {
	checkMark := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")

	return fmt.Sprintf("%s %s", checkMark, b.Name)
}

func (b SubStage) InProgressUserMessage() string {
	// nameStyles := lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	// s := b.inProgressMsg + " "
	return b.inProgressMsg
}

func NewSubStage(name, inProgressMsg string, id sub_command_ids.SubCommandId, onRun func(cfg *build_config.BuildConfigOpts)) *SubStage {
	return &SubStage{
		Name:          name,
		inProgressMsg: inProgressMsg,
		id:            id,
		onRun:         onRun,
		status:        run_status.NotStarted,
	}
}
