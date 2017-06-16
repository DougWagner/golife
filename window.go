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
	for i := uint16(0); i < w.size.row; i++ {
		for j := uint16(0); j < w.size.col; j++ {
			if w.cTree.Search(int32(j)+w.x, int32(i)+w.y) != nil {
				buf.WriteString("O")
			} else {
				buf.WriteString(" ")
			}
		}
		if i < w.size.row-1 {
			buf.WriteString("\n")
		}
	}
	fmt.Print(buf.String())
	w.frames++
}

// Life is the main loop of Conway's Game of Life.
func (w *Window) Life() {
	for {
		w.renderFrame()
		ncChan := make(chan *Cell, 10000)
		var ncChanCount int
		dcChan := make(chan *Cell, 10000)
		var dcChanCount int
		ecTree := initCellTree()
		w.TraverseAndUpdate(w.cTree.root, dcChan, ncChan, &dcChanCount, &ncChanCount, ecTree)
		w.TraverseAndUpdate(ecTree.root, dcChan, ncChan, &dcChanCount, &ncChanCount, ecTree)
		for i := 0; i < dcChanCount; i++ {
			// for some reason removing the cell directly from the channel breaks the tree...
			// i don't think the pointers in the channel are updating correctly as things are
			// removed, so i had to modify the CellTree remove method to include a search...
			// try to fix later.
			removedCell := <-dcChan
			w.cTree.Remove(removedCell.x, removedCell.y)
		}
		for i := 0; i < ncChanCount; i++ {
			newCell := <-ncChan
			w.cTree.Insert(newCell.x, newCell.y)
		}
	}
}

// TraverseAndUpdate will traverse the CellTree underneath
// the Cell c. The function will call the CheckNeighbors
// method to update the CellTree field in the window.
func (w *Window) TraverseAndUpdate(c *Cell, dch, nch chan *Cell, dchCount, nchCount *int, ect *CellTree) {
	if c == nil {
		return
	}
	if c.left != nil {
		w.TraverseAndUpdate(c.left, dch, nch, dchCount, nchCount, ect)
	}
	w.cTree.CheckNeighbors(c.x, c.y, dch, nch, dchCount, nchCount, ect)
	if c.right != nil {
		w.TraverseAndUpdate(c.right, dch, nch, dchCount, nchCount, ect)
	}
}
