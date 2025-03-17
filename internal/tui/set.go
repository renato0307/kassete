// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package tui

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/renato0307/kassete/internal/components/tabs"
	"github.com/renato0307/kassete/internal/config"
	"github.com/renato0307/kassete/internal/logger"
)

type setModel struct {
	tabs tabs.Model
	log  *slog.Logger
}

func newSetModel(set config.Set) tea.Model {
	names := make([]string, 0, len(set.Items))
	for _, item := range set.Items {
		names = append(names, item.Name)
	}

	tabs := tabs.NewTabs(set.Name, names, names)
	tabs.ShowTitle(false)
	return setModel{
		tabs: tabs,
		log:  logger.DefaultLogger().With("module", "set"),
	}
}

func QuitSet() tea.Msg {
	return QuitSetMsg{}
}

type QuitSetMsg struct {
}

func (m setModel) Init() tea.Cmd {
	return m.tabs.Init()
}

func (m setModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.log.Debug("updating")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			return m, QuitSet
		}
	}

	var cmd tea.Cmd
	m.tabs, cmd = m.tabs.Update(msg)
	return m, cmd
}

func (m setModel) View() string {
	m.log.Debug("viewing")
	return m.tabs.View()
}
