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

func RunSubCommand(s *SubStage, cfg *build_config.BuildManager, streamId stream_ids.StreamId, defaultCmd tea.Cmd) tea.Cmd {
	go func() {
		defer func() {
			s.status = run_status.Complete
			build_config.GetBuildManager().Program.Send(SuccessfulSubCmdMsg(s.Name))
		}()
		s.Run(cfg, streamId)
	}()
	return defaultCmd
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
