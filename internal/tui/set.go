// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type setModel struct {
}

func newSetModel() tea.Model {
	return setModel{}
}

func QuitSet() tea.Msg {
	return QuitSetMsg{}
}

type QuitSetMsg struct {
}

func (m setModel) Init() tea.Cmd {
	return nil
}

func (m setModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "esc" {
			return m, QuitSet
		}
	}

	return m, nil
}

func (m setModel) View() string {
	return "This is a set"
}
