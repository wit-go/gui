package gui

import "strconv"
// import "fmt"

/*
	generic function to create a new node on the binary tree
*/
func (n *Node) New(title string) *Node {
	var newN *Node

	newN = addNode(title, n.Width, n.Height)

	n.Append(newN)
	newN.parent = n
	return newN
}

/*
	raw create function for a new node struct
*/
func addNode(title string, w int, h int) *Node {
	var n Node
	n.Name = title
	n.Width = w
	n.Height = h

	id := Config.prefix + strconv.Itoa(Config.counter)
	Config.counter += 1
	n.id = id
	/*
	if (Data.NodeMap[title] != nil) {
		panic(fmt.Sprintf("Duplicate window name = %s\n", title))
	} else {
		Data.NodeMap[title] = &n
	}
	*/

	return &n
}
