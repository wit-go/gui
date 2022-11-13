package gui

// import "git.wit.org/wit/gui/toolkit"

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
func addNode(title string, width int, height int) *Node {
	var n Node

	n.Name = title
	n.Width = width
	n.Height = height

	n.Widget.Name = title
	n.Widget.Width = width
	n.Widget.Height = height

	// no longer a string
	// id := Config.prefix + strconv.Itoa(Config.counter)
	// n.id = id

	n.id = Config.counter
	Config.counter += 1

	return &n
}
