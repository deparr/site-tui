package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) homeSwitch() (model, tea.Cmd) {
	m = m.switchPage(homePage)
	return m, nil
}

func (m model) homeUpdate(msg tea.Msg) (model, tea.Cmd) {
	_ = msg
	return m, nil
}

const bird = `
                            
  ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀  
  ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⠛⠀⠀⠛⣤⠀⠀⠀⠀⠀⠀⠀⠀  
  ⠀⠀⠀⠀⠀⠀⠀⠀⣤⠛⠀⠀⠀⠛⠀⠀⣿⣤⠀⠀⠀⠀⠀⠀  
  ⠀⠀⠀⠀⠀⠀⣤⠛⠀⣤⠛⠛⠀⠀⣤⠛⠀⠀⠀⠀⠀⠀⠀⣤  
  ⠀⠀⠀⠀⠀⣤⠛⠛⠛⠀⣤⣤⣤⠛⠀⠀⠀⠀⣤⣤⣤⠛⠛⠀  
  ⠀⠀⠀⣤⣿⣤⣤⣤⣤⠛⠛⣤⠀⣤⣤⣿⠛⠛⠛⠀⠀⠀⠀⠀  
  ⠀⠀⠛⠀⣤⣤⣤⣤⣿⣤⣤⠛⣿⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤  
  ⣤⠛⠛⠛⠛⠀⠀⠀⠀⠛⠋⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⠛⠛⠀  
  ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⠛⠛⠀⠀⠀⠀⠀  
  ⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⣤⣤⠛⠛⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀  
  ⣤⣤⣤⣤⣤⣿⠛⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀   
`

const content = `
Welcome! Here you can find a little bit about me and what I get up to.

I'm a recent CS grad with special interests in programming languages, developer tools, terminals, and out of necessity: webdev 😢.
`
const name = `     __          _    __
 ___/ /__ __  __(_)__/ /          
/ _  / _ '/ |/ / / _  /           
\_,_/\_,_/|___/_/\_,_/     __  __ 
   ___  ___ ____________  / /_/ /_
  / _ \/ _ '/ __/ __/ _ \/ __/ __/
 / .__/\_,_/_/ /_/  \___/\__/\__/ 
/_/                               
`

func (m model) homeView() string {
	olines := []string{}
	for range 40 {
		olines = append(olines, "LINE")
	}

	birdView := m.theme.Base().Bold(true).Render(bird)
	birdWidth := lipgloss.Width(birdView)
	textWidth := m.contentWidth - birdWidth - 6
	return lipgloss.JoinHorizontal(lipgloss.Top,
		birdView,
		m.theme.Base().Width(3).Render(),
		lipgloss.JoinVertical(lipgloss.Left,
			m.theme.TextHighlight().Width(textWidth).Render(name),
			m.theme.Base().Width(textWidth).Render(content),
			"",
		),
	)
}
