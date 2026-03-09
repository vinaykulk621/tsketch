package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinaykulk621/tsketch/internal/logger"
)

// main update goes here
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// log everything
	logger.Log.Info("update message", "msg", msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		//setting window size
		m.height = msg.Height
		m.width = msg.Width

		//initializing the terminal grid
		m = terminalWindow(m)

		//capturing mouse movements
	case tea.MouseMsg:
		m.cursorX = msg.X
		m.cursorY = msg.Y

		//updating the terminal grid
		m.terminal[m.cursorY][m.cursorX] = '%'

	case tea.KeyMsg:
		switch msg.String() {
		//exiting application
		case "ctrl+c", "q":
			return m, tea.Quit

		//clear terminal
		//"shift+c"
		case "C":
			m = terminalWindow(m)
			return m, nil
		}
	}
	return m, nil
}

// initialize/clear terminal window
func terminalWindow(m Model) Model {
	m.terminal = nil
	m.terminal = make([][]rune, m.height)
	for i := range m.terminal {
		m.terminal[i] = make([]rune, m.width)
	}
	return m
}
