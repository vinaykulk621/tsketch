package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinaykulk621/tsketch/internal/app"
)

func main() {
	p := tea.NewProgram(app.New(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
