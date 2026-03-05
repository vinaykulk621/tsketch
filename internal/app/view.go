package app

import (
	"strings"
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
	return b.String()
}
