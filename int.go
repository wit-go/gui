package gui

// import "github.com/davecgh/go-spew/spew"


/*
	Get the int from the gui toolkit
	because eventually this gui package should become it's own seperate go routine and never interact from the
	gui subroutine back into the upstream application using the gui package

	TODO: instead store the int in the Node structure? (this is probably a better idea)
	because technically every interaction with the toolkit has to go through the Queue() goroutine.
	Is it "has to go" or "should go"? Probably it makes sense to strictly inforce it. No "callback" functions. IPC only (go channels)
*/
func (n *Node) Int() int {
	log(debugToolkit, "gui.Node.Int() for node name =", n.Name)
	log(debugToolkit, SPEW, n)

	// FIXME: this needs to be redone
	// i := n.toolkit.Value()
	i := 3333
	return i
}

// which name to use?
func (n *Node) Value() int {
	return n.Int()
}

func (n *Node) SetValue(i int) {
	log(debugGui, "gui.SetValue() START")
	n.Dump()
	// FIXME: this needs to be redone
	// n.toolkit.SetValue(i)
}
