package main

import (
	"bytes"
	"fmt"
)

// Window structure contains the necessary information for the
// current terminal window. x and y contain the current origin
// point (top left of terminal) for the window. size contains
// the current terminal size data. cTree is the CellTree that
// stores all currently active Cells in the window. frames is
// the number of frames that have been rendered by the window.
type Window struct {
	x, y   int32
	size   *winSize
	cTree  *CellTree
	frames int
}

// initWindow will initialize a Window structure, and clear the
// currently visible terminal for use.
func initWindow() *Window {
	window := &Window{}
	window.size = getWinSize()
	window.cTree = initCellTree()
	buf := bytes.Buffer{}
	hideCursor()
	for i := uint16(0); i < window.size.row; i++ {
		for j := uint16(0); j < window.size.col; j++ {
			buf.WriteString(fmt.Sprint(" "))
		}
	}
	fmt.Print(buf.String())
	return window
}

// renderFrame prints all currently active Cells in their
// appropriate coordinate on the terminal.
func (w *Window) renderFrame() {
	resetCursorLoc()
	w.size = getWinSize()
	buf := bytes.Buffer{}
	for i := uint16(0); i < w.size.row-1; i++ {
		for j := uint16(0); j < w.size.col; j++ {
			if w.cTree.Search(int32(j)+w.x, int32(i)+w.y) != nil {
				buf.WriteString("X")
			} else {
				buf.WriteString(" ")
			}
		}
	}
	fmt.Print(buf.String())
	w.frames++
}
