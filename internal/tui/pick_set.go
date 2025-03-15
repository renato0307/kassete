// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package tui

import (
	"log/slog"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/renato0307/kassete/internal/config"
	"github.com/renato0307/kassete/internal/logger"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type pickSetModel struct {
	list list.Model
	log  *slog.Logger
}

type pickSetItem struct {
	set config.Set
}

func (i pickSetItem) Title() string       { return i.set.Name }
func (i pickSetItem) Description() string { return i.set.Description }
func (i pickSetItem) FilterValue() string { return i.set.Name }

func PickSet(set config.Set) tea.Cmd {
	return func() tea.Msg {
		return PickSetMsg{Set: set}
	}
}

type PickSetMsg struct {
	Set config.Set
}

func newPickSetModel() tea.Model {
	cfg := getConfig()
	log := logger.DefaultLogger().With("module", "pick_set")
	log.Debug("creating sets model", "items", len(cfg.Sets))

	items := make([]list.Item, 0, len(cfg.Sets))
	for _, set := range cfg.Sets {
		log.Debug("adding set to model", "name", set.Name)
		items = append(items, pickSetItem{set: set})
	}

	delegate := list.NewDefaultDelegate()
	delegate.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			logger.DefaultLogger().Debug("updating list after key", "key", msg.String())
			if msg.String() == "enter" {
				return PickSet(m.SelectedItem().(pickSetItem).set)
			}
		}
		return nil
	}
	list := list.New(items, delegate, 0, 0)
	list.Title = "Pick a set to inspect"
	return pickSetModel{
		list: list,
		log:  log,
	}
}

func (m pickSetModel) Init() tea.Cmd {
	return nil
}

func (m pickSetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case PickSetMsg:
		m.log.Debug("picked set", "name", msg.Set.Name)
		return dispatch(msg)
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m pickSetModel) View() string {
	return docStyle.Render(m.list.View())
}
