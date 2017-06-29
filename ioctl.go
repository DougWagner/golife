package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// winSize structure contains the size of the current
// terminal window.
type winSize struct {
	row uint16
	col uint16
	x   uint16 // unused: should be zero
	y   uint16 // unused: should be zero
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

// resetCursorLoc moves the cursor to the top left corner
// of the terminal.
func resetCursorLoc(x, y int) {
	fmt.Printf("\033[%v;%vH", x, y)
}

// hideCursor hides the termial cursor.
func hideCursor() {
	fmt.Printf("\033[?25l")
}

// showCursor shows the terminal cursor.
func showCursor() {
	fmt.Printf("\033[?25h")
}

// The following function is for debugging purposes and should
// never be called in normal execution of the program.

// debugPrintWinSize prints the window size to stdout.
func debugPrintWinSize() {
	window := getWinSize()
	fmt.Printf("col: %v\nrow: %v\nx: %v\ny: %v\n", window.col, window.row, window.x, window.y)
}
