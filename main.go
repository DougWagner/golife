package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage: golife [run | edit]")
	os.Exit(64)
}

func main() {
	switch len(os.Args) {
	case 1:
		runLife()
	case 2:
		if os.Args[1] == "run" {
			runLife()
		} else if os.Args[1] == "edit" {
			runEdit()
		} else {
			printUsage()
		}
	default:
		printUsage()
	}
}

func runLife() {
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

func runEdit() {
	win := initWindow()
	win.Edit()
}
