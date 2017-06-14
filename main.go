package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	setupTerminateHandler()
	win := initWindow()

	// H
	win.cTree.Insert(4, 4)
	win.cTree.Insert(4, 5)
	win.cTree.Insert(4, 6)
	win.cTree.Insert(4, 7)
	win.cTree.Insert(4, 8)
	win.cTree.Insert(5, 6)
	win.cTree.Insert(6, 6)
	win.cTree.Insert(7, 4)
	win.cTree.Insert(7, 5)
	win.cTree.Insert(7, 6)
	win.cTree.Insert(7, 7)
	win.cTree.Insert(7, 8)

	// E
	win.cTree.Insert(9, 4)
	win.cTree.Insert(9, 5)
	win.cTree.Insert(9, 6)
	win.cTree.Insert(9, 7)
	win.cTree.Insert(9, 8)
	win.cTree.Insert(10, 4)
	win.cTree.Insert(10, 6)
	win.cTree.Insert(10, 8)
	win.cTree.Insert(11, 4)
	win.cTree.Insert(11, 6)
	win.cTree.Insert(11, 8)

	// L
	win.cTree.Insert(13, 4)
	win.cTree.Insert(13, 5)
	win.cTree.Insert(13, 6)
	win.cTree.Insert(13, 7)
	win.cTree.Insert(13, 8)
	win.cTree.Insert(14, 8)
	win.cTree.Insert(15, 8)

	// L
	win.cTree.Insert(17, 4)
	win.cTree.Insert(17, 5)
	win.cTree.Insert(17, 6)
	win.cTree.Insert(17, 7)
	win.cTree.Insert(17, 8)
	win.cTree.Insert(18, 8)
	win.cTree.Insert(19, 8)

	// O
	win.cTree.Insert(21, 5)
	win.cTree.Insert(21, 6)
	win.cTree.Insert(21, 7)
	win.cTree.Insert(22, 4)
	win.cTree.Insert(22, 8)
	win.cTree.Insert(23, 4)
	win.cTree.Insert(23, 8)
	win.cTree.Insert(24, 5)
	win.cTree.Insert(24, 6)
	win.cTree.Insert(24, 7)

	// W
	win.cTree.Insert(28, 4)
	win.cTree.Insert(28, 5)
	win.cTree.Insert(28, 6)
	win.cTree.Insert(28, 7)
	win.cTree.Insert(29, 8)
	win.cTree.Insert(30, 7)
	win.cTree.Insert(31, 8)
	win.cTree.Insert(32, 7)
	win.cTree.Insert(32, 6)
	win.cTree.Insert(32, 5)
	win.cTree.Insert(32, 4)

	// O
	win.cTree.Insert(34, 5)
	win.cTree.Insert(34, 6)
	win.cTree.Insert(34, 7)
	win.cTree.Insert(35, 4)
	win.cTree.Insert(35, 8)
	win.cTree.Insert(36, 4)
	win.cTree.Insert(36, 8)
	win.cTree.Insert(37, 5)
	win.cTree.Insert(37, 6)
	win.cTree.Insert(37, 7)

	// R
	win.cTree.Insert(39, 4)
	win.cTree.Insert(39, 5)
	win.cTree.Insert(39, 6)
	win.cTree.Insert(39, 7)
	win.cTree.Insert(39, 8)
	win.cTree.Insert(40, 4)
	win.cTree.Insert(40, 6)
	win.cTree.Insert(41, 4)
	win.cTree.Insert(41, 6)
	win.cTree.Insert(41, 7)
	win.cTree.Insert(42, 5)
	win.cTree.Insert(42, 8)

	// L
	win.cTree.Insert(44, 4)
	win.cTree.Insert(44, 5)
	win.cTree.Insert(44, 6)
	win.cTree.Insert(44, 7)
	win.cTree.Insert(44, 8)
	win.cTree.Insert(45, 8)
	win.cTree.Insert(46, 8)

	// D
	win.cTree.Insert(48, 4)
	win.cTree.Insert(48, 5)
	win.cTree.Insert(48, 6)
	win.cTree.Insert(48, 7)
	win.cTree.Insert(48, 8)
	win.cTree.Insert(49, 4)
	win.cTree.Insert(49, 8)
	win.cTree.Insert(50, 4)
	win.cTree.Insert(50, 8)
	win.cTree.Insert(51, 5)
	win.cTree.Insert(51, 6)
	win.cTree.Insert(51, 7)

	// !
	win.cTree.Insert(53, 4)
	win.cTree.Insert(53, 5)
	win.cTree.Insert(53, 6)
	win.cTree.Insert(53, 8)

	// so far all this will do is draw HELLO WORLD! on the terminal
	for {
		win.renderFrame()
	}
}

// setupTerminateHandler captures the terminate signal
// and ensures that the cursor reappears when the program
// terminates.
func setupTerminateHandler() {
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		resetCursorLoc()
		showCursor()
		os.Exit(0)
	}()
}
