package app

import tea "github.com/charmbracelet/bubbletea"

// main init func
func (m Model) Init() tea.Cmd {
	return nil
}

// create new model obj
func New() Model {
	return Model{}
}
