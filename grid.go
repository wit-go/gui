package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// Grid numbering examples (H) or (W,H)
// ---------
// -- (1) --
// -- (2) --
// ---------
//
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
	// a.Width = w
	// a.Height = h
	newNode.X = w
	newNode.Y = h
	newNode.NextX = 1
	newNode.NextY = 1
	newaction(&a, newNode, n)

	return newNode
}
