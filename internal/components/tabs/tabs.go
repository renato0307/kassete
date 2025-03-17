// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package tabs

import (
	"log/slog"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/renato0307/kassete/internal/logger"
)

type Model struct {
	activeTab int
	content   []string
	names     []string
	title     string
	showTitle bool

	log *slog.Logger
}

func NewTabs(title string, tabs []string, tabContent []string) Model {
	return Model{
		activeTab: 0,
		content:   tabContent,
		names:     tabs,
		title:     title,
		showTitle: true,
		log:       logger.DefaultLogger().With("module", "tabs"),
	}
}

func (m Model) ShowTitle(show bool) *Model {
	m.showTitle = show
	return &m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	m.log.Debug("updating")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "right", "l", "n", "tab":
			m.activeTab = min(m.activeTab+1, len(m.names)-1)
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		}
	}

	return m, nil
}

var (
	docStyle       = lipgloss.NewStyle().Padding(1, 2, 1, 4)
	windowStyle    = lipgloss.NewStyle().Padding(2, 0).Align(lipgloss.Left)
	activeTabStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Underline(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
			Padding(0, 4, 0, 0)
	inactiveTabStyle = activeTabStyle.
				Underline(false).
				Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"})
	titleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1, 0, 1).
			Margin(0, 0, 1, 0)
)

func (m Model) View() string {
	m.log.Debug("viewing")

	var renderedTabs []string
	for i, t := range m.names {
		var style lipgloss.Style
		isActive := i == m.activeTab
		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc := strings.Builder{}
	doc.WriteString(titleStyle.Render(m.title))
	doc.WriteString("\n")
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Render(m.content[m.activeTab]))

	return docStyle.Render(doc.String())
}
