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
	born   int
	died   int
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
	resetCursorLoc(0, 0)
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
}

// Life is the main loop of Conway's Game of Life.
func (w *Window) Life() {
	w.setupLifeInputHandler()
	w.setupLifeTerminateHandler()
	ticker := time.Tick(100 * time.Millisecond)
	for {
		<-ticker
		w.renderFrame()
		w.frames++
		var ncBuff []*Cell
		var dcBuff []*Cell
		ecTree := initCellTree()
		w.TraverseAndUpdate(w.cTree.root, &dcBuff, &ncBuff, ecTree)
		w.TraverseAndUpdate(ecTree.root, &dcBuff, &ncBuff, ecTree)
		for _, v := range dcBuff {
			// So the issue with removing nodes directly from the slice (previously a channel)
			// lies in the else condition of the Cell Remove method. Since it just copies the
			// values of the minRight node, the pointers stored in the slice will become invalid.
			// Maybe fix this later for minimal optimization.
			//w.cTree.Remove(v)
			w.cTree.Remove(v.x, v.y)
		}
		w.died += len(dcBuff)
		for _, v := range ncBuff {
			newCell := v
			w.cTree.Insert(newCell.x, newCell.y)
		}
		w.born += len(ncBuff)
	}
}

// TraverseAndUpdate will traverse the CellTree underneath
// the Cell c. The function will call the CheckNeighbors
// method to update the CellTree field in the window.
func (w *Window) TraverseAndUpdate(c *Cell, db, nb *[]*Cell, ect *CellTree) {
	if c == nil {
		return
	}
	w.TraverseAndUpdate(c.left, db, nb, ect)
	w.cTree.CheckNeighbors(c.x, c.y, db, nb, ect)
	w.TraverseAndUpdate(c.right, db, nb, ect)
}

// setupLifeInputHandler enables the ability to scroll up, down,
// left, or right on the window. Scrolling can use arrow keys,
// or WASD keys.
func (w *Window) setupLifeInputHandler() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// disable displaying of input characters on screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	b := make([]byte, 1)
	go func() {
		for {
			// this feels really hacky for some reason...
			os.Stdin.Read(b)
			switch b[0] {
			case 27:
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
			case 119:
				w.y--
			case 115:
				w.y++
			case 100:
				w.x++
			case 97:
				w.x--
			}
		}
	}()
}

// setupLifeTerminateHandler captures the terminate signal
// and ensures that the cursor reappears when the program
// terminates.
func (w *Window) setupLifeTerminateHandler() {
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)
	go func() {
		<-sigChan
		w.renderFrame()
		fmt.Printf("\n")
		fmt.Printf("Number of generations progressed: %v\n", w.frames)
		fmt.Printf("Number of living cells remaining: %v\n", w.cTree.count)
		fmt.Printf("Total number of cells born: %v\n", w.born)
		fmt.Printf("Total number of cells died: %v\n", w.died)
		// undo stty commands ran in setupInputHandler
		exec.Command("stty", "-F", "/dev/tty", "-cbreak").Run() // not sure if this one is needed
		exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		showCursor()
		os.Exit(0)
	}()
}
