// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package tui

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/renato0307/kassete/internal/config"
	"github.com/renato0307/kassete/internal/logger"
)

type model struct {
	config  config.Config
	current tea.Model
	log     *slog.Logger
}

var root model

func NewRootModel(config config.Config) model {
	root = model{config: config}
	root.current = newPickSetModel()
	root.log = logger.DefaultLogger().With("module", "tui")
	return root
}

func (m model) Init() tea.Cmd {
	return m.getCurrent().Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.log.Debug("updating model")
	switch msg.(type) {
	case PickSetMsg:
		m.log.Debug("picked set")
		m.current = newSetModel()
	case QuitSetMsg:
		m.log.Debug("quit set")
		m.current = newPickSetModel()
	}
	return m.getCurrent().Update(msg)
}

func (m model) View() string {
	return m.getCurrent().View()
}

func (m model) getCurrent() tea.Model {
	if m.current == nil {
		panic("current model is nil")
	}
	return m.current
}

func getRootModel() *model {
	return &root
}

func getConfig() config.Config {
	return getRootModel().config
}

func dispatch(msg tea.Msg) (tea.Model, tea.Cmd) {
	return getRootModel(), func() tea.Msg {
		return msg
	}
}
