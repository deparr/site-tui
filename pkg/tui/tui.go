package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deparr/portfolio/go/pkg/client"
	"github.com/deparr/portfolio/go/pkg/tui/theme"
)

type page int
type size int

const (
	splashPage page = iota
	homePage
	contactPage
	projectsPage
	experiencePage
	debugPage
)

const (
	undersized size = iota
	small
	medium
	large
	fill
)

type model struct {
	page            page
	theme           theme.Theme
	renderer        *lipgloss.Renderer
	state           state
	viewport        viewport.Model
	termWidth       int
	termHeight      int
	size            size
	contentWidth    int
	contentHeight   int
	containerWidth  int
	containerHeight int
	switched        bool
	ready           bool
	hasScroll       bool
}

type state struct {
	splash  splashState
	footer  footerState
	project *projectState
	spinner spinnerState
}

func NewModel(renderer *lipgloss.Renderer) tea.Model {
	return model{
		page:     splashPage,
		theme:    theme.BaseTheme(renderer),
		renderer: renderer,
		state: state{
			splash:  splashState{delay: false},
			project: &projectState{},
			footer: footerState{
				binds: []footerBinding{
					{key: "j/k", action: "scroll"},
					{key: "q", action: "quit"},
					{key: "f", action: "maximize"},
				},
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	checkEnv()
	client.Init()
	return tea.Batch(m.splashInit(), asyncStartProject)
}

func (m model) switchPage(newPage page) model {
	m.page = newPage
	m.switched = true
	return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height

		switch {
		case m.termWidth < 80 || m.termHeight < 30:
			m.size = undersized
			m.containerWidth = m.termWidth
			m.containerHeight = m.termHeight
		case m.termWidth < 100 && m.termHeight < 40:
			m.size = small
			m.containerWidth = 80
			m.containerHeight = 30
		case m.termWidth < 100:
			m.size = medium
			m.containerWidth = m.termWidth
			m.containerHeight = m.termHeight
		default:
			m.size = large
			m.containerWidth = min(m.termWidth, 120)
			m.containerHeight = min(msg.Height, 42)
		}

		m.contentWidth = m.containerWidth
		m = m.updateViewport()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "f":
			if m.size == fill {
				cmds = append(cmds, tea.WindowSize())
			} else {
				m.size = fill
				m.containerWidth = m.termWidth
				m.containerHeight = m.termHeight
				m.contentWidth = m.containerWidth
				m = m.updateViewport()
			}
		}
	case asyncJobMsg:
		switch msg.key {
		case projectKey:
			m.state.project.status = loading
			m.state.project.err = nil
			cmds = append(cmds, asyncUpdateProject)
		}
	case asyncDoneMsg:
		switch msg.key {
		case projectKey:
			m.state.project.update(msg)
		}
	}

	var cmd tea.Cmd
	switch m.page {
	case splashPage:
		m, cmd = m.splashUpdate(msg)
	case contactPage:
		m, cmd = m.contactUpdate(msg)
	case projectsPage:
		m, cmd = m.projectsUpdate(msg)
	}

	m, headerCmd := m.headerUpdate(msg)
	cmds = append(cmds, headerCmd)

	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	m.viewport.SetContent(m.getContent())
	m.viewport, cmd = m.viewport.Update(msg)
	if m.switched {
		m = m.updateViewport()
		m.switched = false
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) updateViewport() model {
	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())
	verticalMarginHeight := headerHeight + footerHeight

	m.contentHeight = m.containerHeight - verticalMarginHeight

	if !m.ready {
		m.viewport = viewport.New(m.contentWidth, m.contentHeight)
		m.viewport.YPosition = headerHeight
		m.viewport.HighPerformanceRendering = false
		m.viewport.KeyMap = viewport.DefaultKeyMap()
		m.ready = true
	} else {
		m.viewport.Width = m.contentWidth
		m.viewport.Height = m.contentHeight
		m.viewport.GotoTop()
	}

	m.hasScroll = m.viewport.VisibleLineCount() < m.viewport.TotalLineCount()

	return m
}

func (m model) View() string {
	if m.size == undersized {
		return m.resizeView()
	}

	switch m.page {
	case splashPage:
		return m.splashView()
	default:
		header := m.headerView()
		footer := m.footerView()

		view := m.viewport.View()
		height := m.containerHeight
		height -= lipgloss.Height(header)
		height -= lipgloss.Height(footer)

		var loc string
		if m.hasScroll {
			loc = m.locView()
		}

		boxedView := lipgloss.JoinVertical(
			lipgloss.Right,
			header,
			m.theme.Base().
				Width(m.contentWidth).
				Height(height).
				Padding(0, 1).
				Render(view),
			loc,
			footer,
		)

		return m.renderer.Place(
			m.termWidth,
			m.termHeight,
			lipgloss.Center,
			lipgloss.Center,
			m.theme.Base().
				MaxWidth(m.termWidth).
				MaxHeight(m.termHeight).
				Render(boxedView),
		)
	}
}

func (m model) locView() string {
	lines := m.viewport.TotalLineCount()
	// todo figure out why this is even possible
	if m.viewport.VisibleLineCount() >= lines {
		return ""
	}
	percent := int(m.viewport.ScrollPercent() * 100)
	loc := int(m.viewport.ScrollPercent() * float64(lines))
	var view string
	switch percent {
	case 0:
		view = "TOP"
	case 100:
		view = "BOT"
	default:
		view = fmt.Sprintf("%d%% %d/%d", percent, loc, m.viewport.TotalLineCount())
	}
	return m.theme.TextAccent().Bold(true).Render(view)
}

func (m model) getContent() string {
	content := "none :("
	switch m.page {
	case homePage:
		content = m.homeView()
	case contactPage:
		content = m.contactView()
	case projectsPage:
		content = m.projectsView()
	case experiencePage:
		content = m.experienceView()
	case debugPage:
		content = m.debugView()
	}

	return content
}
