package gui

import "log"
// import "fmt"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"


import "github.com/davecgh/go-spew/spew"

// functions for handling text entry boxes

func NewLabel(box *GuiBox, text string) {
	box.Append(ui.NewLabel(text), false)
}

func (n *Node) NewLabel(text string) *Node {
	// make new node here
	newNode := makeNode(n, text, 333, 334)
	newNode.Dump()

	t := toolkit.NewLabel(n.uiBox, text)
	newNode.Toolkit = t

	return newNode
}

func (b *GuiBox) GetText(name string) string {
	if (b.Window.EntryMap == nil) {
		log.Println("gui.GetText() ERROR b.Box.Window.EntryMap == nil")
		return ""
	}
	spew.Dump(b.Window.EntryMap)
	if (b.Window.EntryMap[name] == nil) {
		log.Println("gui.GetText() ERROR box.Window.EntryMap[", name, "] == nil ")
		return ""
	}
	e := b.Window.EntryMap[name]
	log.Println("gui.GetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
	log.Println("gui.GetText() END")
	return e.UiEntry.Text()
}

func (n *Node) SetText(value string) error {
	log.Println("gui.SetText() value =", value)
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
	return nil
}
