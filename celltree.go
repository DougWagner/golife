package main

import "fmt"

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

// Search will find a Cell within the CellTree that corresponds to
// the coordinates x and y.
func (ct *CellTree) Search(x, y int32) *Cell {
	pos := Pos(x, y)
	return ct.root.Search(pos)
}

// Insert will insert a new Cell into the CellTree. If the CellTree
// is empty, then root will be assigned to the new Cell, otherwise
// realInsert will take care of the bulk of the insertion process.
func (ct *CellTree) Insert(x, y int32) {
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
func (ct *CellTree) Remove(target *Cell) {
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
