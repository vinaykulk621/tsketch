package app

import (
	"github.com/charmbracelet/lipgloss"
)

// main view goes here
func (m Model) View() string {
	// return fmt.Sprintf("Clicked at:X: %d, y: %d", m.cursorX, m.cursorY)

	terminalMinWidth, terminalMinHeight := 0, 0
	terminalMaxWidth, terminalMaxHeight := m.width, m.height

	// TODO: log these to see what is going on
	posX, posY := getTerminalPos(
		m.cursorX,
		m.cursorY,
		terminalMinWidth,
		terminalMaxWidth,
		terminalMinHeight,
		terminalMaxHeight,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Position(posX),
		lipgloss.Position(posY),
		"%",
	)
}

// returns position of the cursor at a given point
func getTerminalPos(cursorX, cursorY, minX, maxX, minY, maxY int) (float32, float32) {
	// normalized = max−min
	// 				--------
	//   			value−min
	//

	//calculating the position on X axis
	posX := (float32(maxX) - float32(minX)) / (float32(cursorX) - float32(minX))

	//calculating the position on Y axis
	posY := (float32(maxY) - float32(minY)) / (float32(cursorY) - float32(minY))

	return posX, posY
}
