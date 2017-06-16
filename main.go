package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	setupTerminateHandler()
	win := initWindow()

	// spawner
	win.cTree.Insert(70, 25)
	win.cTree.Insert(70, 26)
	win.cTree.Insert(70, 29)
	win.cTree.Insert(71, 25)
	win.cTree.Insert(71, 28)
	win.cTree.Insert(72, 25)
	win.cTree.Insert(72, 28)
	win.cTree.Insert(72, 29)
	win.cTree.Insert(73, 27)
	win.cTree.Insert(74, 25)
	win.cTree.Insert(74, 27)
	win.cTree.Insert(74, 28)
	win.cTree.Insert(74, 29)

	// glider gun
	win.cTree.Insert(5, 9)
	win.cTree.Insert(5, 10)
	win.cTree.Insert(6, 9)
	win.cTree.Insert(6, 10)
	win.cTree.Insert(15, 9)
	win.cTree.Insert(15, 10)
	win.cTree.Insert(15, 11)
	win.cTree.Insert(16, 8)
	win.cTree.Insert(16, 12)
	win.cTree.Insert(17, 7)
	win.cTree.Insert(17, 13)
	win.cTree.Insert(18, 7)
	win.cTree.Insert(18, 13)
	win.cTree.Insert(19, 10)
	win.cTree.Insert(20, 8)
	win.cTree.Insert(20, 12)
	win.cTree.Insert(21, 9)
	win.cTree.Insert(21, 10)
	win.cTree.Insert(21, 11)
	win.cTree.Insert(22, 10)
	win.cTree.Insert(25, 7)
	win.cTree.Insert(25, 8)
	win.cTree.Insert(25, 9)
	win.cTree.Insert(26, 7)
	win.cTree.Insert(26, 8)
	win.cTree.Insert(26, 9)
	win.cTree.Insert(27, 6)
	win.cTree.Insert(27, 10)
	win.cTree.Insert(29, 5)
	win.cTree.Insert(29, 6)
	win.cTree.Insert(29, 10)
	win.cTree.Insert(29, 11)
	win.cTree.Insert(39, 7)
	win.cTree.Insert(39, 8)
	win.cTree.Insert(40, 7)
	win.cTree.Insert(40, 8)

	win.Life()
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
