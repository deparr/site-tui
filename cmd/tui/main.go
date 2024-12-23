package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deparr/portfolio/go/pkg/tui"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	renderer := lipgloss.DefaultRenderer()
	prog := tea.NewProgram(tui.NewModel(renderer))
	if _, err := prog.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
