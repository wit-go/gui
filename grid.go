package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// Grid numbering examples (X) or (X,Y)
// ---------
// -- (1) --
// -- (2) --
// ---------
//
// -----------------------------
// -- (1,1) -- (1,2) -- (1,3) --
// -- (2,1) -- (2,2) -- (2,3) --
// -- (3,1) --       -- (2,3) --
// -----------------------------

func (n *Node) NewGrid(name string, x int, y int) *Node {
	newNode := n.New(name, toolkit.Grid, func() {
		log(debugChange, "click() NewGrid not defined =", name)
	})

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.X = x
	a.Y = y
	newNode.X = x
	newNode.Y = y
	newNode.NextX = 1
	newNode.NextY = 1
	newaction(&a, newNode, n)

	return newNode
}

func (n *Node) NewBox(name string, b bool) *Node {
	newNode := n.New(name, toolkit.Box, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = name
	a.Text = name
	a.B = b
	newaction(&a, newNode, n)

	return newNode
}
