package gui

import "log"

import "github.com/davecgh/go-spew/spew"

// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"


// Get the int from the gui toolkit
// TODO: instead store the int in the Node structure? (this is probably a better idea)
// because eventually this gui package should become it's own seperate go routine and never interact from the
// gui subroutine back into the upstream application using the gui package
func (n *Node) Int() int {
	if (Config.DebugToolkit) {
		log.Println("gui.Node.Int() for node name =", n.Name)
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(n)
	}

	if (n.Toolkit == nil) {
		log.Println("gui.Node.Int() for toolkit struct = nil")
		return 0
	}

	i := n.Toolkit.Value()
	return i
}
