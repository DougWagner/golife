package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"unsafe"
)

// winSize structure contains the size of the current
// terminal window.
type winSize struct {
	row     uint16
	col     uint16
	unusedX uint16 // unused: should be zero
	unusedY uint16 // unused: should be zero
}

// getWinSize uses the ioctl syscall to return a winSize
// structure with the current terminal size.
func getWinSize() *winSize {
	win := &winSize{}
	ret, _, err := syscall.Syscall(uintptr(syscall.SYS_IOCTL), uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(win)))
	if int(ret) == -1 {
		panic(err)
	}
	return win
}

// setCursorLoc moves the cursor to the top left corner
// of the terminal.
func setCursorLoc(x, y int) {
	fmt.Printf("\033[%v;%vH", y, x)
}

// hideCursor hides the termial cursor.
func hideCursor() {
	fmt.Printf("\033[?25l")
}

// showCursor shows the terminal cursor.
func showCursor() {
	fmt.Printf("\033[?25h")
}

func stty() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// disable displaying of input characters on screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func unstty() {
	// undo stty commands ran in stty()
	exec.Command("stty", "-F", "/dev/tty", "-cbreak").Run() // not sure if this one is needed
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

// The following function is for debugging purposes and should
// never be called in normal execution of the program.

// debugPrintWinSize prints the window size to stdout.
func debugPrintWinSize() {
	window := getWinSize()
	fmt.Printf("col: %v\nrow: %v\nx: %v\ny: %v\n", window.col, window.row, window.unusedX, window.unusedY)
}
