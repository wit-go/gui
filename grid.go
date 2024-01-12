package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/widget"
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

type Grid struct {
	Width int
	Height int
}

type GridOffset struct {
	X int
	Y int
}

func (n *Node) NewGrid(name string, w int, h int) *Node {
	newNode := n.newNode(name, widget.Grid)

	newNode.W = w
	newNode.H = h
	newNode.NextW = 1
	newNode.NextH = 1

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}

	// by default, always pad grids
	newNode.Pad()
	return newNode
}

// true if the grid already have a child at W,H
func (n *Node) gridCollision() bool {
	for _, child := range n.children {
		if ((child.AtW == n.NextW) && (child.AtH == n.NextH)) {
			return true
		}
	}
	return false
}

// keeps incrementing NextW & NextH until there is not a widget
func (n *Node) gridIncrement() {
	if ! n.gridCollision() {
		return
	}

	n.NextW += 1
	if (n.NextW > n.W) {
		n.NextW = 1
		n.NextH += 1
	}

	n.gridIncrement()
}

func (n *Node) At(w int, h int) *Node {
	if (n == nil) {
		return n
	}

	n.NextW = w
	n.NextH = h

	n.gridIncrement()
	if (n.NextW != w) || (n.NextH != h) {
		log.Warn("At() (W,H)", w, h, " was moved to avoid a collision (W,H) =", n.NextW, n.NextH)
	}
	return n
}
