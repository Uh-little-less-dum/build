package mocks

import tea "github.com/charmbracelet/bubbletea"

func MockTeaProgram() *tea.Program {
	return tea.NewProgram(NewMockModel())
}
