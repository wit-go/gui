package gui

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

	// no longer a string
	// id := Config.prefix + strconv.Itoa(Config.counter)
	// n.id = id

	n.id = Config.counter
	Config.counter += 1

	return &n
}
