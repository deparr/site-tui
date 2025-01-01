package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) debugSwitch() (model, tea.Cmd) {
	m = m.switchPage(debugPage)
	return m, nil
}

func (m model) debugUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

func rgba2String(r, g, b, _ uint32) string {
	rb := byte((float64(r) / 65536.0) * 256)
	gb := byte((float64(g) / 65536.0) * 256)
	bb := byte((float64(b) / 65536.0) * 256)

	return fmt.Sprintf("#%02x%02x%02x", rb, gb, bb)
}

func (m model) debugView() string {
	header := "DEBUG\n\n"
	profile := "profile: " + lipgloss.ColorProfile().Name()
	dimensions := lipgloss.JoinVertical(
		lipgloss.Left,
		"content: "+fmt.Sprintf("%d x %d", m.contentWidth, m.contentHeight),
		"container: "+fmt.Sprintf("%d x %d", m.containerWidth, m.containerHeight),
		"term: "+fmt.Sprintf("%d x %d", m.termWidth, m.termHeight),
		"m.size: "+fmt.Sprintf("%d", m.size),
		"column: "+fmt.Sprintf("%d", (m.contentWidth - 8) / 2),
		"",
	)

	base := m.theme.Base()

	theme := m.theme
	colors := lipgloss.JoinVertical(
		lipgloss.Left,
		profile,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"body: ",
			rgba2String(theme.Body().RGBA()),
			" ",
			base.Foreground(theme.Body()).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"border: ",
			rgba2String(theme.Border().RGBA()),
			" ",
			base.Foreground(theme.Border()).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"highlight: ",
			rgba2String(theme.Highlight().RGBA()),
			" ",
			base.Foreground(theme.Highlight()).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"accent: ",
			rgba2String(theme.Accent().RGBA()),
			" ",
			base.Foreground(theme.Accent()).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"background: ",
			rgba2String(theme.Background().RGBA()),
			" ",
			base.Foreground(theme.Background()).Render("█"),
		),
		" ",
	)

	faint := lipgloss.JoinVertical(
		lipgloss.Left,
		"faint colors:",
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"body: ",
			rgba2String(theme.Body().RGBA()),
			" ",
			base.Foreground(theme.Body()).Faint(true).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"border: ",
			rgba2String(theme.Border().RGBA()),
			" ",
			base.Foreground(theme.Border()).Faint(true).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"highlight: ",
			rgba2String(theme.Highlight().RGBA()),
			" ",
			base.Foreground(theme.Highlight()).Faint(true).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"accent: ",
			rgba2String(theme.Accent().RGBA()),
			" ",
			base.Foreground(theme.Accent()).Faint(true).Render("█"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			"background: ",
			rgba2String(theme.Background().RGBA()),
			" ",
			base.Foreground(theme.Background()).Faint(true).Render("█"),
		),
		" ",
	)

	projects := fmt.Sprintf("pinned: %v", m.state.project.pinnedRepos)
	recent := fmt.Sprintf("recent: %v", m.state.project.recentRepos)
	projstate := fmt.Sprintf(
		"projstate\nerr: %v\nstatus: %d",
		m.state.project.err,
		m.state.project.status,
	)


	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		dimensions,
		colors,
		faint,
		"",
		projstate,
		projects,
		"",
		recent,
	)
}
