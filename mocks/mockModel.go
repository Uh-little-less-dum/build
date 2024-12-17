package mocks

import tea "github.com/charmbracelet/bubbletea"

type MockModel struct {
}

func (m MockModel) Init() tea.Cmd {
	return nil
}

func (m MockModel) View() string {
	return ""
}

func (m MockModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func NewMockModel() *MockModel {
	return &MockModel{}
}
