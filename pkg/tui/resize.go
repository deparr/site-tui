package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) resizeView() string {
	return m.renderer.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			fmt.Sprintf("your terminal is too small (%dx%d)", m.termWidth, m.termHeight),
			"min is >= 80x30",
		),
	)
}
