package gui

import "log"
import "errors"

// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

// functions for handling text related GUI elements

func (n *Node) NewLabel(text string) *Node {
	// make new node here
	newNode := n.New(text)
	newNode.Dump()

	t := n.Toolkit.NewLabel(text)
	newNode.Toolkit = t

	return newNode
}

func (n *Node) SetText(value string) error {
	log.Println("gui.SetText() value =", value)
	if (n.Toolkit != nil) {
		n.Toolkit.SetText(value)
		return nil
	}
	if (n.uiText != nil) {
		n.uiText.SetText(value)
		return nil
	}
	if (n.uiButton != nil) {
		n.uiButton.SetText(value)
		return nil
	}
	if (n.uiWindow != nil) {
		n.uiWindow.SetTitle(value)
		return nil
	}
	return errors.New("nothing found for gui.Node.SetText()")
}
