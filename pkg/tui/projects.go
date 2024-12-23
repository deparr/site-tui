package tui

import (
	"context"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/deparr/portfolio/go/pkg/resource"
	"github.com/google/go-github/v67/github"
)

type language struct {
	name    string
	color   string
	percent int
}

type Project struct {
	name     string
	repoUrl  string
	language []language
	stars    int
	desc     string
}

const (
	defaultLangColor = "#a5a8a6"
)

var (
	toget = []string{
		"chip8",
		"cs330",
		"portfolio",
	}
	projects []Project
	loadedProj   = false
)

func (m model) projectsSwitch() (model, tea.Cmd) {
	m = m.switchPage(projectsPage)
	return m, nil
}

func (m model) projectsUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

func (m model) projectsView() string {
	if projects == nil || !loadedProj {
		return lipgloss.PlaceVertical(1, lipgloss.Center, "loading...")
	}

	lines := []string{""}
	header := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Width(m.contentWidth).PaddingLeft(4).Render
	link := m.theme.Base().Faint(true).Width(m.contentWidth).PaddingLeft(4).Render
	for _, proj := range projects {
		lines = append(lines,
			lipgloss.JoinVertical(
				lipgloss.Left,
				header(proj.name),
				base(proj.desc),
				link(proj.repoUrl),
				"    "+colorBar(proj.language, m.contentWidth-8)+"    ",
				"",
			),
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func colorBar(langs []language, width int) string {
	usedWidth := 0
	widths := []int{}
	for _, l := range langs {
		portionedWidth := max(width*l.percent/100, 1)
		widths = append(widths, portionedWidth)
		usedWidth += portionedWidth
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
		langColor := defaultLangColor
		if l.color != "" {
			langColor = l.color
		}

		color := lipgloss.AdaptiveColor{Dark: langColor, Light: langColor}
		renderedSegment := style.Foreground(color).Render(strings.Repeat("─", w))
		colorBar = append(colorBar, renderedSegment)
	}

	rendered := lipgloss.JoinHorizontal(lipgloss.Center, colorBar...)
	return rendered
}

func populateProjects() {
	client := github.NewClient(nil)
	for _, repo := range toget {
		repoData, _, err := client.Repositories.Get(context.Background(), "deparr", repo)
		if err != nil {
			continue
		}
		langData, _, err := client.Repositories.ListLanguages(context.Background(), "deparr", repo)
		langs := []language{}
		if err == nil {
			sizeTotal := 0
			for l, s := range langData {
				langs = append(langs, language{name: l, color: resource.LanguageColors[l], percent: s})
			}
			slices.SortFunc(langs, func(a language, b language) int {
				return b.percent - a.percent
			})
			if len(langs) > 5 {
				langs = langs[:5]
			}
			for _, l := range langs {
				sizeTotal += l.percent
			}
			for i := range langs {
				langs[i].percent = langs[i].percent * 100 / sizeTotal
			}
		} else {
			langs = append(langs, language{name: *repoData.Language, color: resource.LanguageColors[*repoData.Language]})
		}

		projects = append(projects, Project{
			name:     *repoData.Name,
			repoUrl:  *repoData.URL,
			language: langs,
			stars:    *repoData.StargazersCount,
			desc:     *repoData.Description,
		})
	}

	client.Client().CloseIdleConnections()
	client = nil
	loadedProj = true
}
