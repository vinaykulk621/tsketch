package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinaykulk621/tsketch/internal/app"
	"github.com/vinaykulk621/tsketch/internal/logger"
)

func main() {

	err := logger.Init("app.log")
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(app.New(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
