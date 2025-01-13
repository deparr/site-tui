package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m model) headerUpdate(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			return m.homeSwitch()
		case "c":
			return m.contactSwitch()
		case "p":
			return m.projectsSwitch()
		case "b":
			return m.blogSwitch()
		case "d":
			return m.debugSwitch()
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) headerView() string {
	base := m.theme.TextBase().Render
	baseUnder := m.theme.TextBase().Underline(true).Render
	highlightSel := m.theme.TextHighlight().Underline(true).Bold(true).Render
	highlight := m.theme.TextHighlight().Bold(true).Render
	logoBase := m.theme.TextAccent().Bold(true).Italic(true).Render
	logoHighlight := m.theme.TextHighlight().Bold(true).Italic(true).Render

	logo := logoHighlight("@") + logoBase("dp")
	home := baseUnder("h") + base("ome")
	projects := baseUnder("p") + base("rojects")
	blog := baseUnder("b") + base("log")
	contact := baseUnder("c") + base("ontact")

	switch m.page {
	case homePage:
		home = highlightSel("h") + highlight("ome")
	case contactPage:
		contact = highlightSel("c") + highlight("ontact")
	case projectsPage:
		projects = highlightSel("p") + highlight("rojects")
	case blogPage:
		blog = highlightSel("b") + highlight("log")
	}

	tabs := []string{
		logo,
		home,
		projects,
		blog,
		contact,
	}

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(m.renderer.NewStyle().Foreground(m.theme.Border())).
		Row(tabs...).
		Width(m.contentWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			return m.theme.Base().
				Padding(0, 1).
				AlignHorizontal(lipgloss.Center)
		}).
		Render()

}
