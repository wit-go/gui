package gui

import "git.wit.org/wit/gui/toolkit"

func (parent *Node) NewButton(name string, custom func()) *Node {
	newNode := parent.newNode(name, toolkit.Button, custom)

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}

/*
// deprecate this once andlabs is refactored
func callback(i int) bool {
	log(debugError, "callback() for widget id =", i)
	n := me.rootNode.FindId(i)
	log(debugError, "callback() found node =", n)
	// running custom here means the button get's clicked twice
	if (n.Custom == nil) {
		log(debugError, "callback() = nil. SKIPPING")
		return false
	}
	n.Custom()
	return true
}
*/

// find widget by number
func (n *Node) FindId(i int) (*Node) {
	if (n == nil) {
		return nil
	}

	if (n.id == i) {
		return n
	}

	for _, child := range n.children {
		newN := child.FindId(i)
		if (newN != nil) {
			return newN
		}
	}
	return nil
}

