package app

import tea "github.com/charmbracelet/bubbletea"

// main update goes here
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		//setting window size
		m.height = msg.Height
		m.width = msg.Width

		return m, nil

		//capturing mouse movements
	case tea.MouseMsg:
		m.cursorX = msg.X
		m.cursorY = msg.Y

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		//exiting application
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
