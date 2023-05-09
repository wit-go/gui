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

	a := newAction(n, toolkit.Add)

	a.X = w
	a.Y = h
	newNode.X = w
	newNode.Y = h
	newNode.NextX = 1
	newNode.NextY = 1

	sendAction(a, newNode, n)

	return newNode
}

// increments where the next element in the grid should go
func placeGrid(a *toolkit.Action, n *Node, where *Node) {
	where.NextY += 1
	if (where.NextY > where.Y) {
		where.NextX += 1
		where.NextY = 1
	}

	a.X = where.NextX
	a.Y = where.NextY
	log(logNow, "placeGrid() (X,Y)", where.X, where.Y, " next(X,Y) =", where.NextX, where.NextY)
}
