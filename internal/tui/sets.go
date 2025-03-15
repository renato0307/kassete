package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type setsModel struct {
}

func newSetsModel() tea.Model {
	return setsModel{}
}

func (m setsModel) Init() tea.Cmd {
	return nil
}

func (m setsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m setsModel) View() string {
	s := "\nPress q to quit.\n"
	return s
}
