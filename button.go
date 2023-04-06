package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) NewButton(name string, custom func()) *Node {
	newNode := n.New(name, toolkit.Button, custom)

	var a toolkit.Action
	a.Name = name
	a.Text = name
	a.ActionType = toolkit.Add
	// deprecate this once andlabs is refactored
	a.Callback = callback
	newaction(&a, newNode, n)

	return newNode
}

// deprecate this once andlabs is refactored
func callback(i int) bool {
	log(debugError, "callback() for widget id =", i)
	n := Config.rootNode.FindId(i)
	log(debugError, "callback() found node =", n)
	// running custom here means the button get's clicked twice
	if (n.Custom == nil) {
		log(debugError, "callback() = nil. SKIPPING")
		return false
	}
	n.Custom()
	return true
}

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

