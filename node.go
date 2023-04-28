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

	n.children = append(n.children, newN)
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

func (n *Node) Delete(d *Node) {
	for i, child := range n.children {
		log(debugNode, "\t", i, child.id, child.Name)
		if (child.id == d.id) {
			log(debugNode, "\t\t Deleting this")
			n.children = append(n.children[:i], n.children[i+1:]...)
			return
		}
	}
	log(debugError, "did not find node to delete", d.id, d.Name)
}
