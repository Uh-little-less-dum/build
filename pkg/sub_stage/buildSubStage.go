package sub_stage

import (
	"fmt"
	"math/rand"
	"time"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
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
	onRun         func(cfg *build_config.BuildManager)
}

func (b *SubStage) Run(cfg *build_config.BuildManager, streamId stream_ids.StreamId) tea.Cmd {
	b.status = run_status.Running
	b.onRun(cfg)
	b.status = run_status.Complete
	return b.CompleteMsg(streamId)
}

func (b SubStage) Id() sub_command_ids.SubCommandId {
	return b.id
}

type InstalledPkgMsg string

func (b SubStage) CompleteMsg(streamId stream_ids.StreamId) tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return InstalledPkgMsg(b.Name)
	})
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

func DownloadAndInstall(s *SubStage) tea.Cmd {
	// This is where you'd do i/o stuff to download and install packages. In
	// our case we're just pausing for a moment to simulate the process.
	d := time.Millisecond * time.Duration(rand.Intn(10000)) //nolint:gosec
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return InstalledPkgMsg(s.Name)
	})
}

func NewSubStage(name, inProgressMsg string, id sub_command_ids.SubCommandId, onRun func(cfg *build_config.BuildManager)) *SubStage {
	return &SubStage{
		Name:          name,
		inProgressMsg: inProgressMsg,
		id:            id,
		onRun:         onRun,
		status:        run_status.NotStarted,
	}
}
