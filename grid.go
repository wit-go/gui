package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// Grid numbering examples (H) or (W,H)
// -----------------------
// -- (1) -- (2) -- (3) -- (X)
// -----------------------
//
//    (Y)
// ---------
// -- (1) --
// -- (2) --
// ---------
//
//    (X,Y)
// -----------------------------
// -- (1,1) -- (2,1) -- (3,1) --
// -- (1,2) -- (2,2) -- (3,2) --
// -- (1,3) --       -- (3,3) --
// -----------------------------

func (n *Node) NewGrid(name string, w int, h int) *Node {
	newNode := n.newNode(name, toolkit.Grid, func() {
		log(debugChange, "click() NewGrid not defined =", name)
	})

	a := newAction(newNode, toolkit.Add)

	a.X = w
	a.Y = h
	newNode.X = w
	newNode.Y = h
	newNode.NextW = 1
	newNode.NextH = 1

	sendAction(a, newNode, n)

	return newNode
}

// true if the grid already have a child at W,H
func (n *Node) gridCollision(w int, h int) bool {
	for _, child := range n.children {
		if ((child.AtW == w) && (child.AtH == h)) {
			return true
		}
	}
	return false
}

// increments NextW & NextH
func (n *Node) gridIncrement(w int, h int) bool {
	for _, child := range n.children {
		if ((child.AtW == w) && (child.AtH == h)) {
			return true
		}
	}
	return false
}

func (n *Node) At(w int, h int) *Node {
	if (n == nil) {
		return n
	}

	n.NextW = w
	n.NextH = h

	// TODO: check for a collision here
	if n.gridCollision(w,h) {
		// TODO: find free next w,h
	}
	return n
}

// finds the next place on the grid to place the new node 'n'
func placeGrid(a *toolkit.Action, n *Node, where *Node) {
	where.NextH += 1
	if (where.NextH > where.Y) {
		where.NextW += 1
		where.NextH = 1
	}

	a.X = where.NextW
	a.Y = where.NextH
	log(logNow, "placeGrid() (X,Y)", where.X, where.Y, " next(X,Y) =", where.NextW, where.NextH)
}
