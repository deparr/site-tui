package tui

import (
	"log/slog"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	apimodel "github.com/deparr/api/pkg/model"
	"github.com/deparr/portfolio/go/pkg/client"
)

var pinnedRepos []apimodel.Repository
var recentRepos []apimodel.Repository

func (m model) projectsSwitch() (model, tea.Cmd) {
	m = m.switchPage(projectsPage)
	return m, nil
}

func (m model) projectsUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

func (m model) projectsView() string {
	// TODO make the ui aware of client errors and show that instead of only `loading...`
	if pinnedRepos == nil {
		return lipgloss.PlaceHorizontal(15, lipgloss.Center, lipgloss.PlaceVertical(1, lipgloss.Center, "loading..."))
	}

	// this is not useful at all, don't know what i was thinking
	// if !client.ApiIsHealthy() {
	// 	return lipgloss.JoinVertical(
	// 		lipgloss.Center,
	// 		"",
	// 		base("Hmm, looks like something went wrong on the backend."),
	// 		base("I won't be able to  display recent projects here, try going to the website:"),
	// 		link(os.Getenv("WEBSITE_URL")),
	// 		base("or check out my github"),
	// 		link(os.Getenv("GITHUB_URL")),
	// 	)
	// }

	base := m.theme.Base().Width(m.contentWidth).PaddingLeft(4).Render
	link := m.theme.Base().Faint(true).Width(m.contentWidth).PaddingLeft(4).Render
	lines := []string{""}
	header := m.theme.TextAccent().Bold(true).Render
	for _, proj := range pinnedRepos {
		lines = append(lines,
			lipgloss.JoinVertical(
				lipgloss.Left,
				header(proj.Name),
				base(proj.Desc),
				link(proj.Url),
				"    "+colorBar(proj.Language, m.contentWidth-8)+"    ",
				"",
			),
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
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
		color := lipgloss.AdaptiveColor{Dark: l.Color, Light: l.Color}
		renderedSegment := style.Foreground(color).Render(strings.Repeat("─", w))
		colorBar = append(colorBar, renderedSegment)
	}

	rendered := lipgloss.JoinHorizontal(lipgloss.Center, colorBar...)
	return rendered
}

func (m model) populateProjects() {
	newPinned, err := client.GetRepos("pinned")
	pinnedRepos = newPinned

	if err != nil {
		slog.Error("populating projects:", "err", err)
	}

	// if we're somehow at the projects page before theyre loaded
	// force a reswitch to show them
	if m.page == projectsPage {
		m.projectsSwitch()
	}
}
