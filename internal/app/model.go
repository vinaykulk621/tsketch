package app

// main model
type Model struct {

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
	insertMode int = iota
	normalMode
)
