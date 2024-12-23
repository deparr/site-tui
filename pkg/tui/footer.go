package tui

import "github.com/charmbracelet/lipgloss"

type footerState struct {
	binds []footerBinding
}

type footerBinding struct {
	key    string
	action string
}

func (m model) footerView() string {
	base := m.theme.Base().Render
	bold := m.theme.TextAccent().Bold(true).Render

	table := m.theme.Base().
		Width(m.contentWidth).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(m.theme.Border()).
		PaddingBottom(1).
		Align(lipgloss.Center)

	var actions []string

	for _, bind := range m.state.footer.binds {
		actions = append(actions, bold(" "+bind.key+" ")+base(bind.action+"  "))
	}

	return table.Render(lipgloss.JoinHorizontal(lipgloss.Center, actions...))
}
