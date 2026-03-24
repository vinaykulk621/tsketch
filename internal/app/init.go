package app

import tea "github.com/charmbracelet/bubbletea"

// main init func
func (m Model) Init() tea.Cmd {
	return nil
}

// help mode init
func (m HelpModel) Init() tea.Cmd {
	return nil
}

// help mode init
func (m CanvaModel) Init() tea.Cmd {
	return nil
}

// create new model obj
func New() Model {
	canva := CanvaModel{tMode: mode(normalMode), taskBarHeight: 3}
	help := HelpModel{}
	return Model{screen: canvaScreen, help: help, canva: canva}
}
