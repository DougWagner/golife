package main

import (
	"fmt"
	"math"
)

// Pos returns the linear index of the coordinates x and y.
// x and y need to be cast to uint8 before uint16 so
// the value doesn't become incorrect if x or y are
// negative.
func Pos(x, y int8) uint16 {
	return (math.MaxUint8+1)*uint16(uint8(y)) + uint16(uint8(x))
}

// Cell structure represents an active location on the
// window grid. Each Cell represents a node in the CellTree
// data structure.
type Cell struct {
	x, y   int8
	left   *Cell
	right  *Cell
	parent *Cell
}

// NewCell initializes a new Cell with x and y coordinates
// and sets the pointers to nil. Pointers will be updated
// during the insertion and removal phase of the tree.
func NewCell(x, y int8) *Cell {
	newCell := &Cell{x, y, nil, nil, nil}
	return newCell
}

// Pos returns the linear index of the cell on the grid.
func (c *Cell) Pos() uint16 {
	return Pos(c.x, c.y)
}

// Search is the recursive search for the position calculated
// by Pos(x, y).
func (c *Cell) Search(pos uint16) *Cell {
	if pos < c.Pos() {
		if c.left != nil {
			return c.left.Search(pos)
		}
		return nil
	} else if pos > c.Pos() {
		if c.right != nil {
			return c.right.Search(pos)
		}
		return nil
	} else {
		return c
	}
}

// Insert recursively inserts a Cell into the CellTree underneath
// the calling Cell.
func (c *Cell) Insert(nc *Cell) {
	if nc.Pos() < c.Pos() {
		if c.left != nil {
			c.left.Insert(nc)
		} else {
			c.left = nc
			nc.parent = c
			return
		}
	} else if nc.Pos() > c.Pos() {
		if c.right != nil {
			c.right.Insert(nc)
		} else {
			c.right = nc
			nc.parent = c
			return
		}
	} else {
		panic("A node of the same value already exists")
	}
}

// MinChild will return the minimum value node in the tree
// beneath the Cell.
func (c *Cell) MinChild() *Cell {
	if c.left == nil {
		return c
	}
	return c.left.MinChild()
}

// Remove will unlink the Cell from the CellTree.
func (c *Cell) Remove() {
	if c.left == nil && c.right == nil {
		if c.parent.left == c {
			c.parent.left = nil
			c.parent = nil
		} else if c.parent.right == c {
			c.parent.right = nil
			c.parent = nil
		} else {
			panic("There is a major issue with your tree and you should feel bad")
		}
	} else if c.right != nil && c.left == nil {
		if c.parent.left == c {
			c.right.parent = c.parent
			c.parent.left = c.right
		} else if c.parent.right == c {
			c.right.parent = c.parent
			c.parent.right = c.right
		} else {
			panic("There is a major issue with your tree and you should feel bad")
		}
	} else if c.left != nil && c.right == nil {
		if c.parent.left == c {
			c.left.parent = c.parent
			c.parent.left = c.left
		} else if c.parent.right == c {
			c.left.parent = c.parent
			c.parent.right = c.left
		} else {
			panic("There is a major issue with your tree and you should feel bad")
		}
	} else {
		minRight := c.right.MinChild()
		c.x = minRight.x
		c.y = minRight.y
		minRight.Remove()
	}
}

// preOrder is a pre order traversal that appends the coordinates
// to a byte slice.
func (c *Cell) preOrder(buff *[]byte) {
	*buff = append(*buff, byte(c.x), byte(c.y))
	if c.left != nil {
		c.left.preOrder(buff)
	}
	if c.right != nil {
		c.right.preOrder(buff)
	}
}

// The following methods are for debugging purposes and should
// never be called in normal execution of the program.

// PrintCell prints the cell data to stdout
func (c *Cell) PrintCell() {
	if c.parent != nil {
		fmt.Printf("x: %v\ty: %v\t Pos: %v\tParent: x: %v, y: %v\n", c.x, c.y, c.Pos(), c.parent.x, c.parent.y)
	} else {
		fmt.Printf("x: %v\ty: %v\t Pos: %v\tParent: nil\n", c.x, c.y, c.Pos())
	}
}

// inOrderPrint is a recursive traversal of the CellTree
// underneath the Cell c to print the tree in order.
func (c *Cell) inOrderPrint() {
	if c.left != nil {
		c.left.inOrderPrint()
	}
	c.PrintCell()
	if c.right != nil {
		c.right.inOrderPrint()
	}
}

// preOrderPrint is a recursive traversal of the CellTree
// underneath the Cell c to print the tree in pre order.
func (c *Cell) preOrderPrint() {
	c.PrintCell()
	if c.left != nil {
		c.left.preOrderPrint()
	}
	if c.right != nil {
		c.right.preOrderPrint()
	}
}

// postOrderPrint is a recursive traversal of the CellTree
// underneath the Cell c to print the tree in post order.
func (c *Cell) postOrderPrint() {
	if c.left != nil {
		c.left.postOrderPrint()
	}
	if c.right != nil {
		c.right.postOrderPrint()
	}
	c.PrintCell()
}
