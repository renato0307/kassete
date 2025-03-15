// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

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
	cfg := getConfig()
	s := "Pick a set:\n"
	for _, set := range cfg.Sets {
		s += set.Name + "\n"
	}
	s = "\nPress q to quit.\n"
	return s
}
