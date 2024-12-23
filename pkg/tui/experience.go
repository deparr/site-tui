package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deparr/portfolio/go/pkg/resource"
)

func (m model) experienceSwitch() (model, tea.Cmd) {
	m = m.switchPage(experiencePage)
	return m, nil
}

func (m model) experienceUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

func (m model) experienceView() string {
	lines := []string{""}
	title := m.theme.TextAccent().Render
	baseFaint := m.theme.Base().Faint(true).Italic(true)
	base := m.theme.Base().Render

	for _, job := range resource.Experience {
		lines = append(lines,
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					title(job.Title),
					baseFaint.PaddingLeft(
						m.contentWidth - len(job.Title) - len(job.Start) - len(job.End),
					).
					Render(job.Start+"—"+job.End),
				),
				baseFaint.Render(job.Company),
				base(job.Desc),
			),
		)	
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
