package app

// main model
type Model struct {
	//screen
	screen screen

	//canva
	canva CanvaModel

	//help screen model
	help HelpModel
}

type CanvaModel struct {

	//width and height of the terminal
	width, height, taskBarHeight int

	// position of the cursor
	cursorX, cursorY int

	//terminal screen, where we will
	//render things as user moves mouse
	terminal [][]rune

	//terminal modes,
	//simial to vimotions
	tMode mode
}

// similar to vimotions
// if presses 'i' should start drawing
type mode int

const (
	insertMode mode = iota
	normalMode
)

func (m mode) toString() string {
	return [...]string{"-- INSERT --", "-- NORMAL --"}[m]
}

// screens
type screen int

const (
	canvaScreen screen = iota
	helpScreen
)

type HelpModel struct {
	width, height int
}
