package app

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// main view goes here
func (m Model) View() string {
	var b strings.Builder

	//rendering from the treminal
	for y := range m.height {
		for x := range m.width {
			// padding
			//
			// m.terminal[y][x] will be zero,
			// because zero value of [][]rune is zero
			if m.terminal[y][x] == 0 {
				b.WriteRune(' ')
			} else {
				b.WriteRune(m.terminal[y][x])
			}
		}

		//newline
		b.WriteRune('\n')

	}
	return b.String() + "\n" + taskBar(m)
}

// taskBar
func taskBar(m Model) string {
	keys := []string{"ctrl+c/quit", "shift+c/clear", m.tMode.toString()}
	return taskBarStyle.
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center).
		Width(m.width).
		Height(m.taskBarHeight).
		Render(strings.Join(keys, "\t"))
}
