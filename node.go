package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/widget"
)

/*
	generic function to create a new node on the binary tree
*/
func (n *Node) newNode(title string, t widget.WidgetType) *Node {
	var newN *Node

	newN = addNode(title)
	newN.WidgetType = t

	if n.WidgetType == widget.Grid {
		n.gridIncrement()
	}
	newN.AtW = n.NextW
	newN.AtH = n.NextH
	newN.hidden = n.hidden // by default, use the value from above

	n.children = append(n.children, newN)
	newN.parent = n
	return newN
}

/*
	raw create function for a new node struct
*/
func addNode(title string) *Node {
	n := new(Node)
	n.label = title
	n.id = me.counter
	log.Log(NODE, "addNode = widget setid =", n.id)

	me.counter += 1
	return n
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Delete(d *Node) {
	for i, child := range n.children {
		log.Log(NODE, "\t", i, child.id, child.progname)
		if (child.id == d.id) {
			log.Log(NODE, "\t\t Deleting this")
			n.children = append(n.children[:i], n.children[i+1:]...)
			return
		}
	}
	log.Warn("did not find node to delete", d.id, d.progname)
}
