package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) contactSwitch() (model, tea.Cmd) {
	m = m.switchPage(contactPage)
	return m, nil
}

func (m model) contactUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

func (m model) contactView() string {
	base := m.theme.Base().Render
	highlight := m.theme.TextHighlight().Render

	contact := lipgloss.JoinVertical(
		lipgloss.Left,
		base("")+highlight("  " + os.Getenv("EMAIL")),
		base("󰊤")+highlight("  github/deparr"),
		base("󰌻")+highlight("  in/dparrott"),
	)

	return m.renderer.Place(
		m.contentWidth,
		m.contentHeight,
		lipgloss.Center,
		lipgloss.Center,
		contact,
	)
}
