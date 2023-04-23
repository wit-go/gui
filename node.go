package gui

import "git.wit.org/wit/gui/toolkit"

/*
	generic function to create a new node on the binary tree
*/
func (n *Node) newNode(title string, t toolkit.WidgetType, custom func()) *Node {
	var newN *Node

	newN = addNode(title)
	newN.WidgetType = t
	newN.Custom = custom

	// TODO: This should not be defined for each widget. This has to be stupid
	// or wait a second, is this where I send something to a channel?
	newN.Custom = func() {
		log(debugChange, "Trying to find Window Close. widgetType =", newN.WidgetType)
		if (newN.WidgetType == toolkit.Window) {
			log(debugChange, "Need to delete newN here")
			n.Delete(newN)
		}
		if (newN.Custom == nil) {
			log(debugChange, "newT.Custom() == nil. Not doing anything. SEND SOMETHING TO THE CHANNEL")
			return
		}
		log(debugChange, "newT.Custom() START SEND SOMETHING TO THE CHANNEL node =", newN.Name)
		// send something to the channel here????
		newN.Custom()
		log(debugChange, "newT.Custom() END   SEND SOMETHING TO THE CHANNEL node =", newN.Name)
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
	n.Text = title
	n.id = Config.counter
	log(debugNode, "addNode = widget setid =", n.id)

	Config.counter += 1
	return n
}

func (n *Node) Parent() *Node {
	return n.parent
}

/*
func (n *Node) Window() *Node {
	return n.parent
}
*/

func (n *Node) Append(child *Node) {
	n.children = append(n.children, child)
	if (debugNode) {
		log(debugNode, "child node:")
		child.Dump(debugDump)
		log(debugNode, "parent node:")
		n.Dump(debugDump)
	}
}

func (n *Node) Delete(d *Node) {
	for i, child := range n.children {
		log(debugNode, "\t", i, child.id, child.Width, child.Height, child.Name)
		if (child.id == d.id) {
			log(debugNode, "\t\t Deleting this")
			n.children = append(n.children[:i], n.children[i+1:]...)
			return
		}
	}
	log(debugError, "did not find node to delete", d.id, d.Width, d.Height, d.Name)
}
