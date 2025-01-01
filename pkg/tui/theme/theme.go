package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	renderer *lipgloss.Renderer

	background lipgloss.TerminalColor
	body       lipgloss.TerminalColor
	border     lipgloss.TerminalColor
	highlight  lipgloss.TerminalColor
	accent     lipgloss.TerminalColor
	error      lipgloss.TerminalColor

	base lipgloss.Style
}

func BaseTheme(renderer *lipgloss.Renderer) Theme {
	base := Theme{
		renderer: renderer,
	}

	base.background = lipgloss.AdaptiveColor{Dark: "#151515", Light: "#dcdcd6"}
	base.body = lipgloss.AdaptiveColor{Dark: "#a5a8a6", Light: "#000000"}
	base.border = lipgloss.AdaptiveColor{Dark: "#696969", Light: "#545454"}
	base.highlight = lipgloss.AdaptiveColor{Dark: "#e78c45", Light: "#d05200"}
	base.accent = lipgloss.AdaptiveColor{Dark: "#e5e8e6", Light: "#151515"}
	base.error = lipgloss.AdaptiveColor{Dark: "#cc6666", Light: "#c82829"}
	base.base = lipgloss.NewStyle().Foreground(base.body)

	return base
}

func (t Theme) Background() lipgloss.TerminalColor {
	return t.background
}

func (t Theme) Body() lipgloss.TerminalColor {
	return t.body
}

func (t Theme) Base() lipgloss.Style {
	return t.base
}

func (t Theme) Border() lipgloss.TerminalColor {
	return t.border
}

func (t Theme) Highlight() lipgloss.TerminalColor {
	return t.highlight
}

func (t Theme) Accent() lipgloss.TerminalColor {
	return t.accent
}

func (t Theme) Error() lipgloss.TerminalColor {
	return t.error
}

func (t Theme) TextBase() lipgloss.Style {
	return t.base.Foreground(t.body)
}

func (t Theme) TextError() lipgloss.Style {
	return t.base.Foreground(t.error)
}

func (t Theme) TextAccent() lipgloss.Style {
	return t.base.Foreground(t.accent)
}

func (t Theme) TextHighlight() lipgloss.Style {
	return t.base.Foreground(t.highlight)
}
