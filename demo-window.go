package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var mybox *ui.Box

func (n *Node) AddGroup(title string) *Node {
	if (n == nil) {
		return nil
	}
	hbox := n.uiBox
	if (hbox == nil) {
		return n
	}
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, Config.Stretchy)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	newNode := n.AddNode(title)
	newNode.uiBox = vbox
	return newNode
}

func (n *Node) MakeBasicControlsPage(title string) *Node {
	if (n == nil) {
		return nil
	}
	origbox := n.uiBox
	if (origbox == nil) {
		return n
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	hbox.Append(ui.NewButton("Button"), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", ui.NewMultilineEntry(), true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	origbox.Append(vbox, false)
	newNode := n.AddNode(title)
	newNode.uiBox = vbox
	return newNode
}

func (n *Node) MakeGroupEdit(title string) *Node {
	n.Dump()

	group := ui.NewGroup(title)
	group.SetMargined(true)
	n.uiBox.Append(group, Config.Stretchy)

	entrybox := ui.NewNonWrappingMultilineEntry()

	group.SetChild(entrybox)

	log.Println("entrybox =", entrybox)
	n.uiMultilineEntry = entrybox
	newNode := n.AddNode(title)
	newNode.uiMultilineEntry = entrybox
	newNode.uiGroup = group
	return newNode
}
