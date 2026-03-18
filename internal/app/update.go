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
		//when we introduce taskBar, our drawable pane gets pushed upwards (3 units, as defined as taskBarHeight var)
		// which in turn makes the drawing experience worse,
		// thus following formula
		terminalCanvaSize := 2 * m.taskBarHeight

		//setting window size
		m.height = msg.Height - terminalCanvaSize
		m.width = msg.Width

		//initializing the terminal grid
		m = terminalWindow(m)

		//capturing mouse movements
	case tea.MouseMsg:
		m.cursorX = msg.X
		m.cursorY = msg.Y

		// updating the terminal grid
		// we only need to check on y-axis as,
		// taskBar gets in the way, which could break the coordinates
		// system of our terminal canvas
		if m.tMode == mode(insertMode) && m.cursorY < m.height {
			m.terminal[m.cursorY][m.cursorX] = '%'
		}

	case tea.KeyMsg:
		switch msg.String() {
		//exiting application
		case "ctrl+c", "q":
			if m.tMode == mode(insertMode) {
				m.tMode = mode(normalMode)
				return m, nil
			}
			return m, tea.Quit

			//enable/disable Insert tMode
		case "i", "I":
			if m.tMode == mode(insertMode) {
				m.tMode = mode(normalMode)
			} else {
				m.tMode = mode(insertMode)
			}

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
