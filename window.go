package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// Window structure contains the necessary information for the
// current terminal window. x and y contain the current origin
// point (top left of terminal) for the window. size contains
// the current terminal size data. cTree is the CellTree that
// stores all currently active Cells in the window. frames is
// the number of frames that have been rendered by the window.
type Window struct {
	x, y   int8
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
	window.setupInputHandler()
	window.setupTerminateHandler()
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
			if w.cTree.Search(int8(j)+w.x, int8(i)+w.y) != nil {
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
	ticker := time.Tick(100 * time.Millisecond)
	for {
		<-ticker
		w.renderFrame()
		ncChan := make(chan *Cell, 65536)
		var ncChanCount int
		dcChan := make(chan *Cell, 65536)
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
	w.TraverseAndUpdate(c.left, dch, nch, dchCount, nchCount, ect)
	w.cTree.CheckNeighbors(c.x, c.y, dch, nch, dchCount, nchCount, ect)
	w.TraverseAndUpdate(c.right, dch, nch, dchCount, nchCount, ect)
}

// setupInputHandler enables the ability to scroll up, down,
// left, or right on the window.
func (w *Window) setupInputHandler() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// disable displaying of input characters on screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	b := make([]byte, 1)
	go func() {
		for {
			// this feels really hacky for some reason...
			os.Stdin.Read(b)
			if b[0] != 27 {
				continue
			}
			os.Stdin.Read(b)
			if b[0] != 91 {
				continue
			}
			os.Stdin.Read(b)
			switch b[0] {
			case 65: // up
				w.y--
			case 66: // down
				w.y++
			case 67: // right
				w.x++
			case 68: // left
				w.x--
			}
		}
	}()
}

// setupTerminateHandler captures the terminate signal
// and ensures that the cursor reappears when the program
// terminates.
func (w *Window) setupTerminateHandler() {
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		w.renderFrame()
		fmt.Printf("\n")
		// undo stty commands ran in setupInputHandler
		exec.Command("stty", "-F", "/dev/tty", "-cbreak").Run() // not sure if this one is needed
		exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		showCursor()
		os.Exit(0)
	}()
}
