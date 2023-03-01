package gui

import "git.wit.org/wit/gui/toolkit"

/*
	generic function to create a new node on the binary tree
*/
func (n *Node) New(title string, t toolkit.WidgetType, custom func()) *Node {
	var newN *Node

	newN = addNode(title)
	newN.widget.Type = t
	newN.widget.Action = "New"
	newN.Custom = custom

	// TODO: This should not be defined for each widget. This has to be stupid
	// or wait a second, is this where I send something to a channel?
	newN.widget.Custom = func() {
		if (newN.Custom == nil) {
			log(debugChange, "newT.Custom() == nil. Not doing anything. SEND SOMETHING TO THE CHANNEL")
			return
		}
		log(debugChange, "newT.Custom() START SEND SOMETHING TO THE CHANNEL widget.Name =", newN.widget.Name)
		// send something to the channel here????
		newN.Custom()
		log(debugChange, "newT.Custom() END   SEND SOMETHING TO THE CHANNEL widget.Name =", newN.widget.Name)
	}

	n.Append(newN)
	newN.parent = n
	return newN
}

/*
	raw create function for a new node struct
*/
func addNode(title string) *Node {
	n := new(Node)
	n.Name = title
	n.widget.Name = title

	n.id = Config.counter
	Config.counter += 1

	return n
}
