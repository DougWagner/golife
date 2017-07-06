package main

import (
	"fmt"
	"io/ioutil"
)

// CellTree structure is the main structure for the bst
// of active cells on the window grid.
type CellTree struct {
	root  *Cell
	count int
}

// initCellTree initializes a new CellTree with appropriate fields
// set to zero or nil. Not sure if I need to initialize the values
// like this, but doing it for sanity's sake.
func initCellTree() *CellTree {
	ct := &CellTree{}
	ct.root = nil
	ct.count = 0
	return ct
}

// initFromFile creates a new CellTree and initializes all of the
// nodes in the tree from a file created by the WriteToFile method.
func initFromFile(fname string) *CellTree {
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(file)%2 != 0 {
		fmt.Println("File length not even")
		return nil
	}
	ct := initCellTree()
	for i := 0; i < len(file)/2; i++ {
		ct.Insert(int8(file[i*2]), int8(file[i*2+1]))
	}
	return ct
}

// Search will find a Cell within the CellTree that corresponds to
// the coordinates x and y.
func (ct *CellTree) Search(x, y int8) *Cell {
	if ct.root == nil {
		return nil
	}
	pos := Pos(x, y)
	return ct.root.Search(pos)
}

// Insert will insert a new Cell into the CellTree. If the CellTree
// is empty, then root will be assigned to the new Cell, otherwise
// realInsert will take care of the bulk of the insertion process.
func (ct *CellTree) Insert(x, y int8) {
	newCell := NewCell(x, y)
	if ct.root == nil { // tree is empty
		ct.root = newCell
	} else {
		ct.root.Insert(newCell)
	}
	ct.count++
}

// Remove removes the target from the Cell. If the target
// cell is root, the function will account for that.
func (ct *CellTree) Remove(x, y int8) {
	target := ct.Search(x, y)
	if target == ct.root {
		tempRoot := NewCell(0, 0)
		tempRoot.left = ct.root
		ct.root.parent = tempRoot
		ct.root = tempRoot
		target.Remove()
		ct.count--
		ct.root = tempRoot.left
		if ct.root != nil {
			ct.root.parent = nil
		}
		tempRoot.left = nil // not sure if this is necessary, but whatever
	} else {
		target.Remove()
		ct.count--
	}
}

// CheckNeighbors checks the tree to see if the 8 Cells around
// x and y are active or not. If a cell is marked for death, it
// is added to dch channel. If an empty cell is marked for birth,
// it is added to nch. If an empty cell is neighboring an active
// cell, it is added to ect CellTree so we can check if empty Cells
// must be born.
func (ct *CellTree) CheckNeighbors(x, y int8, db, nb *[]*Cell, ect *CellTree) {
	xNb := []int8{x - 1, x, x + 1, x - 1, x + 1, x - 1, x, x + 1}
	yNb := []int8{y - 1, y - 1, y - 1, y, y, y + 1, y + 1, y + 1}
	var n int
	currentCell := ct.Search(x, y)
	if currentCell != nil { // cell is live, check if it must die
		for i := 0; i < len(xNb); i++ {
			if ct.Search(xNb[i], yNb[i]) != nil {
				n++
			} else {
				if ect.Search(xNb[i], yNb[i]) == nil {
					ect.Insert(xNb[i], yNb[i])
				}
			}
		}
		if n < 2 || n > 3 {
			*db = append(*db, currentCell)
		}
	} else { // cell is not live, check if it must be born
		for i := 0; i < len(xNb); i++ {
			if ct.Search(xNb[i], yNb[i]) != nil {
				n++
			}
		}
		if n == 3 {
			*nb = append(*nb, NewCell(x, y))
		}
	}
}

// WriteToFile writes the CellTree to a binary file with a pre
// order traversal. When the file is read back into a new tree
// using the initFromFile function, the tree structure should
// be the same.
func (ct *CellTree) WriteToFile(fname string) {
	if ct.count == 0 {
		return
	}
	buff := []byte{} // make buffer with minimum space to store 1 cell
	ct.root.preOrder(&buff)
	ioutil.WriteFile(fname, buff, 0664)
}

// The following methods are for debugging purposes and should
// never be called in normal execution of the program.

// inOrderPrint is the initial caller for an in order traversal
// of the CellTree for printing.
func (ct *CellTree) inOrderPrint() {
	fmt.Printf("Count: %v\n", ct.count)
	if ct.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	ct.root.inOrderPrint()
}

// preOrderPrint is the initial caller for a pre order traversal
// of the CellTree for printing.
func (ct *CellTree) preOrderPrint() {
	fmt.Printf("Count: %v\n", ct.count)
	if ct.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	ct.root.preOrderPrint()
}

// postOrderPrint is the initial caller for a post order traversal
// of the CellTree for printing.
func (ct *CellTree) postOrderPrint() {
	fmt.Printf("Count: %v\n", ct.count)
	if ct.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	ct.root.postOrderPrint()
}
