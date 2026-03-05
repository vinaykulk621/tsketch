package app

// main model goes here

type Model struct {

	//width and height of the terminal
	width, height int

	// position of the cursor
	cursorX, cursorY int

	//terminal screen, where we will
	//render things as user moves mouse
	terminal [][]rune
}
