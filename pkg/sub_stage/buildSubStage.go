package sub_stage

import (
	"fmt"
	"sync"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	run_status "github.com/Uh-little-less-dum/go-utils/pkg/constants/runStatus"
	stream_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/streamIds"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type SubStage struct {
	Name            string
	id              sub_command_ids.SubCommandId
	status          run_status.RunStatus
	inProgressMsg   string
	program         *tea.Program
	onRun           func(cfg *build_config.BuildManager)
	concurrentIndex int
}

type GetSubStageFunc func(p *tea.Program) []*SubStage

func (b *SubStage) Run(cfg *build_config.BuildManager, streamId stream_ids.StreamId) tea.Cmd {
	b.status = run_status.Running
	b.onRun(cfg)
	b.status = run_status.Complete
	return b.CompleteMsg(streamId)
}

func (b SubStage) Id() sub_command_ids.SubCommandId {
	return b.id
}

type SuccessfulSubCmdMsg string

func (b SubStage) CompleteMsg(streamId stream_ids.StreamId) tea.Cmd {
	return func() tea.Msg {
		return SuccessfulSubCmdMsg(b.Name)
	}
}

func (b *SubStage) SetConcurrentIndex(idx int) *SubStage {
	b.concurrentIndex = idx
	return b
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

func (s *SubStage) HasRan() bool {
	return s.status == run_status.Complete
}

func (s *SubStage) ConcurrentIndex() (concurrentIndex int, hasIndex bool) {
	return s.concurrentIndex, s.concurrentIndex >= 0
}

func (s *SubStage) InConcurrentGroup(concurrentIndex int) bool {
	if (concurrentIndex >= 0) && (concurrentIndex == s.concurrentIndex) {
		return true
	}
	return false
}

// FIX: Passing the cfg here is redundent. Remove this.
func RunSubCommand(s *SubStage, cfg *build_config.BuildManager, streamId stream_ids.StreamId, defaultCmd tea.Cmd) tea.Cmd {
	go func() {
		defer func() {
			s.status = run_status.Complete

			// cfg2 := build_config.GetBuildManager()
			// s.program.Send(signals.SendSubStageCompleteMsg(streamId, s.id)())
			build_config.GetBuildManager().Program.Send(SuccessfulSubCmdMsg(s.Name))
		}()
		s.Run(cfg, streamId)
	}()
	// This is where you'd do i/o stuff to download and install packages. In
	// our case we're just pausing for a moment to simulate the process.
	// return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
	// 	return InstalledPkgMsg(s.Name)
	// })
	return defaultCmd
	// return m, tea.Batch(m.spinner.Spinner.Tick, m.waitForActivity(m.sub))
}

func RunConcurrently(subStages []*SubStage, cfg *build_config.BuildManager, streamId stream_ids.StreamId, waitForActivity tea.Cmd) {
	log.Fatal("This is completely untested. Come back here and handle this while on wifi.")
	var wg sync.WaitGroup
	for _, s := range subStages {
		wg.Add(1)
		go func() {
			defer wg.Done()
			RunSubCommand(s, cfg, streamId, waitForActivity)
		}()
	}
	wg.Wait()
}

func NewSubStage(name, inProgressMsg string, id sub_command_ids.SubCommandId, onRun func(cfg *build_config.BuildManager), program *tea.Program) *SubStage {
	return &SubStage{
		Name:            name,
		inProgressMsg:   inProgressMsg,
		id:              id,
		onRun:           onRun,
		status:          run_status.NotStarted,
		concurrentIndex: -1,
		program:         program,
	}
}
