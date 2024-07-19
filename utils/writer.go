package utils

import (
	tm "github.com/buger/goterm"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
}

// WriteText writes text to the screen at the current cursor position
func WriteText(text string) {
	tm.Print(text)
	tm.Flush()
}

// WriteLn writes text to the screen at the current cursor position and moves to the next line
func WriteLn(text string) {
	tm.Println(text)
	tm.Flush()
}

// WriteAt writes text to the screen at the specified position
func WriteAt(text string, x, y int) {
	tm.MoveCursor(x, y)
	tm.Print(text)
	tm.Flush()
}
