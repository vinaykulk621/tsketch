package app

import (
	"github.com/charmbracelet/lipgloss"
)

// main view goes here
func (m Model) View() string {
	return lipgloss.NewStyle().
		PaddingLeft(m.cursorX).
		PaddingTop(m.cursorY).
		Render("%")
}
