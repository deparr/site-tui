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
		case "e":
			return m.experienceSwitch()
		case "d":
			return m.debugSwitch()
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) headerView() string {
	base := m.theme.Base().Render
	highlight := m.theme.TextHighlight().Bold(true).Render
	accent := m.theme.TextAccent().Render
	logoBase := m.theme.TextAccent().Bold(true).Italic(true).Render
	logoAccent := m.theme.TextHighlight().Bold(true).Italic(true).Render

	logo := logoAccent("@") + logoBase("dp")
	home := accent("h") + base(" home")
	projects := accent("p") + base(" projects")
	experience := accent("e") + base(" experience")
	contact := accent("c") + base(" contact")

	switch m.page {
	case homePage:
		home = accent("h") + highlight(" home")
	case contactPage:
		contact = accent("c") + highlight(" contact")
	case projectsPage:
		projects = accent("p") + highlight(" projects")
	case experiencePage:
		experience = accent("e") + highlight(" experience")
	}

	tabs := []string{
		logo,
		home,
		projects,
		experience,
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
