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
	return ct.realSearch(ct.root, pos)
}

// realSearch is the recursive search for the position calculated
// by Pos(x, y) in the Search method.
func (ct *CellTree) realSearch(root *Cell, pos uint64) *Cell {
	if root == nil {
		return nil
	}
	if pos < root.Pos() {
		return ct.realSearch(root.left, pos)
	} else if pos > root.Pos() {
		return ct.realSearch(root.right, pos)
	} else {
		return root
	}
}

// Insert will insert a new Cell into the CellTree. If the CellTree
// is empty, then root will be assigned to the new Cell, otherwise
// realInsert will take care of the bulk of the insertion process.
func (ct *CellTree) Insert(x, y int32) {
	newCell := NewCell(x, y)
	if ct.root == nil { // tree is empty
		ct.root = newCell
		ct.count++
	} else {
		ct.realInsert(ct.root, newCell)
	}
}

// realInsert recursively inserts a Cell into the CellTree.
// realInsert should only be called if CellTree root is not nil.
// I'm not going to worry about privacy because this isn't a package,
// it's an application.
func (ct *CellTree) realInsert(root, c *Cell) {
	if c.Pos() < root.Pos() {
		if root.left != nil {
			ct.realInsert(root.left, c)
		} else {
			root.left = c
			c.parent = root
			ct.count++
			return
		}
	} else if c.Pos() > root.Pos() {
		if root.right != nil {
			ct.realInsert(root.right, c)
		} else {
			root.right = c
			c.parent = root
			ct.count++
			return
		}
	} else {
		panic("A node of the same value already exists")
	}
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
	ct.inOrderPrintTraverse(ct.root)
}

// preOrderPrint is the initial caller for a pre order traversal
// of the CellTree for printing.
func (ct *CellTree) preOrderPrint() {
	fmt.Printf("Count: %v\n", ct.count)
	if ct.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	ct.preOrderPrintTraverse(ct.root)
}

// postOrderPrint is the initial caller for a post order traversal
// of the CellTree for printing.
func (ct *CellTree) postOrderPrint() {
	fmt.Printf("Count: %v\n", ct.count)
	if ct.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	ct.postOrderPrintTraverse(ct.root)
}

// inOrderPrintTraverse is the actual recursive traversal of the
// CellTree to print the tree in order.
func (ct *CellTree) inOrderPrintTraverse(root *Cell) {
	if root.left != nil {
		ct.inOrderPrintTraverse(root.left)
	}
	root.PrintCell()
	if root.right != nil {
		ct.inOrderPrintTraverse(root.right)
	}
}

// preOrderPrintTraverse is the actual recursive traversal of the
// CellTree to print the tree in pre order.
func (ct *CellTree) preOrderPrintTraverse(root *Cell) {
	root.PrintCell()
	if root.left != nil {
		ct.preOrderPrintTraverse(root.left)
	}
	if root.right != nil {
		ct.preOrderPrintTraverse(root.right)
	}
}

// postOrderPrintTraverse is the actual recursive traversal of the
// CellTree to print the tree in post order.
func (ct *CellTree) postOrderPrintTraverse(root *Cell) {
	if root.left != nil {
		ct.postOrderPrintTraverse(root.left)
	}
	if root.right != nil {
		ct.postOrderPrintTraverse(root.right)
	}
	root.PrintCell()
}
