package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type blogState struct {
	//posts []apimodel.BlogPosts
	err    error
	status asyncStatus
}

func (s *blogState) update(msg asyncDoneMsg) {
}

func (m model) blogSwitch() (model, tea.Cmd) {
	m = m.switchPage(blogPage)
	return m, nil
}

func (m model) blogUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

func (m model) blogView() string {
	base := m.theme.TextBase().Render
	switch m.state.blog.status {
	case empty:
		return m.renderer.Place(
			m.contentWidth,
			m.contentHeight,
			lipgloss.Center,
			0.80,
			base("no posts yet :("),
		)

	}

	return ""
}
