package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Window structure contains the necessary information for the
// current terminal window.
type Window struct {
	x, y   int8      // window origin coordinates
	cX, cY uint16    // cursor coordinates
	size   *winSize  // window size structure
	cTree  *CellTree // main cell tree
	frames int       // total frames executed
	born   int       // total cells born
	died   int       // total cells died
}

// initWindow will initialize a Window structure, and clear the
// currently visible terminal for use.
func initWindow() *Window {
	window := &Window{}
	window.cX = 1
	window.cY = 1
	window.size = getWinSize()
	window.cTree = initCellTree()
	buf := bytes.Buffer{}
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
	setCursorLoc(1, 1)
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
	hideCursor()
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
	stty()
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
		unstty()
		showCursor()
		os.Exit(0)
	}()
}

// Edit sets up the Window as an editor to input cells to
// use in Life.
func (w *Window) Edit() {
	if w.cTree.count != 0 {
		w.renderFrame()
	}
	showCursor()
	setCursorLoc(int(w.cX), int(w.cY))
	inputChan := make(chan bool)
	w.setupEditInputHandler(inputChan)
	w.setupEditTerminateHandler()
	//w.setupEditWinchangeHandler()
	for {
		<-inputChan
		w.renderFrame()
		setCursorLoc(int(w.cX), int(w.cY))
	}
}

// setupEditInputHandler initializes the ability to move the
// cursor and insert/remove nodes from the editor window.
func (w *Window) setupEditInputHandler(inputChan chan bool) {
	stty()
	b := make([]byte, 1)
	go func() {
		for {
			// this also feels hacky...
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
					w.decCY()
				case 66: // down
					w.incCY()
				case 67: // right
					w.incCX()
				case 68: // left
					w.decCX()
				}
			case 119:
				w.decCY()
			case 115:
				w.incCY()
			case 100:
				w.incCX()
			case 97:
				w.decCX()
			case 32:
				if w.cTree.Search(w.x+int8(w.cX)-1, w.y+int8(w.cY)-1) == nil {
					w.cTree.Insert(w.x+int8(w.cX)-1, w.y+int8(w.cY)-1)
				} else {
					w.cTree.Remove(w.x+int8(w.cX)-1, w.y+int8(w.cY)-1)
				}
			}
			inputChan <- true
		}
	}()
}

// decCY decrements the window cursor on the Y axis.
func (w *Window) decCY() {
	if w.cY == 1 {
		w.y--
	} else {
		w.cY--
	}
}

// incCY increments the window cursor on the Y axis.
func (w *Window) incCY() {
	if w.cY == w.size.row {
		w.y++
	} else {
		w.cY++
	}
}

// incCX increments the window cursor on the X axis.
func (w *Window) incCX() {
	if w.cX == w.size.col {
		w.x++
	} else {
		w.cX++
	}
}

// decCX decrements the window cursor on the X axis.
func (w *Window) decCX() {
	if w.cX == 1 {
		w.x--
	} else {
		w.cX--
	}
}

// setupEditTerminateHandler will allow the editor to
// gracefully exit on program termination.
func (w *Window) setupEditTerminateHandler() {
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)
	go func() {
		<-sigChan
		setCursorLoc(0, int(w.size.row))
		fmt.Printf("\n")
		unstty()
		os.Exit(0)
	}()
}

// setupEditWinchangeHandler doesn't work very well right now
// and should be ignored for the time being.
func (w *Window) setupEditWinchangeHandler() {
	// TODO: make this function work better
	winSigChan := make(chan os.Signal, 2)
	signal.Notify(winSigChan, syscall.SIGWINCH)
	go func() {
		<-winSigChan
		w.renderFrame()
	}()
}
