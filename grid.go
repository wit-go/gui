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

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = name
	a.Text = name
	a.X = w
	a.Y = h
	newNode.X = w
	newNode.Y = h
	newNode.NextX = 1
	newNode.NextY = 1

	/*
	// fix values here if they are invalid. Index starts at 1
	if (where.NextX < 1) {
		where.NextX = 1
	}
	if (where.NextY < 1) {
		where.NextY = 1
	}
	//
	a.X = where.NextX
	a.Y = where.NextY
	*/

	newaction(&a, newNode, n)

	/*
	where.NextY += 1
	if (where.NextY > where.Y) {
		where.NextX += 1
		where.NextY = 1
	}
	log(logInfo, "Action() END size (X,Y)", where.X, where.Y, "put next thing at (X,Y) =", where.NextX, where.NextY)
	*/

	return newNode
}
