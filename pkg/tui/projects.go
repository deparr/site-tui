package tui

import (
	"log/slog"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	apimodel "github.com/deparr/api/pkg/model"
	"github.com/deparr/portfolio/go/pkg/client"
)

// todo consider state interface type
type projectState struct {
	pinnedRepos []apimodel.Repository
	recentRepos []apimodel.Repository
	err         error
	status      asyncStatus
}

func (s *projectState) update(msg asyncDoneMsg) {
	if msg.err != nil {
		s.status = errored
		s.err = msg.err
		return
	}
	s.pinnedRepos = msg.data["pinned"].([]apimodel.Repository)
	s.recentRepos = msg.data["recent"].([]apimodel.Repository)
	s.status = loaded
}

func (m model) projectsSwitch() (model, tea.Cmd) {
	m = m.switchPage(projectsPage)
	return m, nil
}

func (m model) projectsUpdate(msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			switch m.state.project.status {
			case errored:
				cmd = asyncStartProject
			case empty:
				cmd = asyncStartProject
			}
		}
	}

	return m, cmd
}

func (m model) projectsView() string {
	switch m.state.project.status {
	case loading:
		return lipgloss.Place(
			m.contentWidth,
			m.contentHeight,
			lipgloss.Center,
			0.80,
			"loading...",
		)
	case errored:
		error := m.theme.TextError().Render
		link := m.theme.TextFaint().Underline(true).Render
		base := m.theme.Base().Render
		return lipgloss.Place(
			m.contentWidth,
			m.contentHeight,
			lipgloss.Center,
			0.80,
			lipgloss.JoinVertical(
				lipgloss.Left,
				base("Error fetching project data:"),
				error(m.state.project.err.Error()),
				"",
				base("press 'r' to try again"),
				"",
				"",
				base("or, check the following if the issue persists:"),
				link(os.Getenv("GITHUB_URL")),
				link(os.Getenv("WEBSITE_URL")),
			),
		)
	case empty:
		return lipgloss.Place(
			m.contentWidth,
			m.contentHeight,
			lipgloss.Center,
			0.80,
			"Project state is empty somehow\n\nPress 'r' to attempt a refetch",
		)
	}

	// todo dynamically create matrices per *section*
	columnWidth := (m.contentWidth - 8) / 2
	base := m.theme.Base().Width(columnWidth).PaddingLeft(2).Render
	link := m.theme.TextFaint().Width(columnWidth).PaddingLeft(2).Render
	header := m.theme.TextAccent().Bold(true).Render
	section := m.theme.TextHighlight().Bold(true).Render

	pinnedHeader := section(titledSeparator("Github Pinned", "─", columnWidth))
	lines := []string{"", pinnedHeader, ""}
	for _, proj := range m.state.project.pinnedRepos {
		lines = append(lines,
			lipgloss.JoinVertical(
				lipgloss.Left,
				header(proj.Name),
				base(proj.Desc),
				link(proj.Url),
				"  "+colorBar(proj.Language, columnWidth-4)+"  ",
				"",
			),
		)
	}

	recentHeader := section(titledSeparator("Recently Updated", "─", columnWidth))
	recentLines := []string{"", recentHeader, ""}

	for _, proj := range m.state.project.recentRepos {
		recentLines = append(recentLines,
			lipgloss.JoinVertical(
				lipgloss.Left,
				header(proj.Name),
				base(proj.Desc),
				link(proj.Url),
				"  "+colorBar(proj.Language, columnWidth-4)+"  ",
				"",
			),
		)
	}

	pinsRendered := lipgloss.JoinVertical(lipgloss.Left, lines...)
	receRendered := lipgloss.JoinVertical(lipgloss.Left, recentLines...)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		pinsRendered,
		"      ",
		receRendered,
	)
}

func titledSeparator(header string, lineChar string, widthToFill int) string {
	return "╭ " + header + " " + strings.Repeat(lineChar, widthToFill-len(header)-4) + "╮"
}

func colorBar(langs apimodel.RepoLangs, width int) string {
	usedWidth := 0
	widths := []int{}
	for _, l := range langs {
		portioned := l.Percent * width / 100
		widths = append(widths, portioned)
		usedWidth += portioned
	}

	leftover := width - usedWidth
	if leftover != 0 {
		addToAll := leftover / len(langs)
		addToBig := leftover % len(langs)

		for i := range widths {
			widths[i] += addToAll
			if addToBig > 0 {
				widths[i] += 1
				addToBig--
			}
		}
	}

	style := lipgloss.NewStyle()
	colorBar := []string{}
	for i := range langs {
		w := widths[i]
		l := langs[i]
		color := lipgloss.Color(l.Color)
		renderedSegment := style.Foreground(color).Render(strings.Repeat("─", w))
		colorBar = append(colorBar, renderedSegment)
	}

	rendered := lipgloss.JoinHorizontal(lipgloss.Center, colorBar...)
	return rendered
}

func asyncStartProject() tea.Msg {
	return asyncJobMsg{key: projectKey}
}

func asyncUpdateProject() tea.Msg {
	m := asyncDoneMsg{key: projectKey, data: map[string]any{}}
	pinned, err := client.GetGhPinned()
	if err != nil {
		slog.Error("populating pinned projects:", "err", err)
		m.err = err
	}

	recent, err := client.GetGhRecent()
	if err != nil {
		slog.Error("populating recent projects:", "err", err)
		m.err = err
	}

	m.data["pinned"] = pinned
	m.data["recent"] = recent

	return m
}
