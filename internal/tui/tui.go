package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/renato0307/kassete/internal/config"
)

type model struct {
	config  config.Config
	current tea.Model
}

var rootModel model

func NewRootModel(config config.Config) model {
	rootModel = model{
		config:  config,
		current: newSetsModel(),
	}
	return rootModel
}

func (m model) Init() tea.Cmd {
	return m.getCurrent().Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
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

func RootModel() model {
	return rootModel
}

func Config() config.Config {
	return RootModel().config
}
