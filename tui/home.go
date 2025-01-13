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

const (
	bird = `
                            
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

	content = `
Welcome!

Here you'll find a little bit about my current projects and, eventually, some blog posts.
I'm a recent CS grad with interested in programming languages, developer tools, terminals, graphics, and out of necessity: webdev 😢.
`
	name = `     __          _    __
 ___/ /__ __  __(_)__/ /          
/ _  / _ '/ |/ / / _  /           
\_,_/\_,_/|___/_/\_,_/     __  __ 
   ___  ___ ____________  / /_/ /_
  / _ \/ _ '/ __/ __/ _ \/ __/ __/
 / .__/\_,_/_/ /_/  \___/\__/\__/ 
/_/                               
`
)

func (m model) homeView() string {
	birdView := m.theme.TextBase().Render(bird)
	birdWidth := lipgloss.Width(birdView)
	textWidth := m.contentWidth - birdWidth - 6
	// todo formating is a little strange here
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
