package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var mybox *ui.Box

func (n *Node) AddDemoTab(title string) {
	newNode := n.AddTab(title, makeDemoTab())
	if (Config.DebugNode) {
		newNode.Dump()
	}
	tabSetMargined(newNode.uiTab)
	newNode.Dump()
	newNode.ListChildren(false)
	addDemoGroup(newNode, "new group 1")
	addDemoGroup(newNode, "new group 2")
	
	groupNode := newNode.AddGroup("new group 3")
	groupNode.AddComboBox("testing", "foo", "man", "blah")
}

func makeDemoTab() *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("DemoEditBox")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("foo 1")
	ecbox.Append("man 2")
	ecbox.Append("bar 3")

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		log.Println("test")
		test := ecbox.Text()
		log.Println("test=", test)
	})

	vbox.Append(ecbox, false)

	return hbox
}

func addDemoGroup(n *Node, title string) {
	hbox := n.uiBox
	if (hbox == nil) {
		return
	}
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("foo 1")
	ecbox.Append("man 2")
	ecbox.Append("bar 3")

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		log.Println("test")
		test := ecbox.Text()
		log.Println("test=", test)
	})

	vbox.Append(ecbox, false)
}

func (n *Node) AddGroup(title string) *Node {
	hbox := n.uiBox
	if (hbox == nil) {
		return n
	}
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	newNode := n.AddNode(title)
	newNode.uiBox = vbox
	return newNode
}

func (n *Node) GetText(title string) string {
	if (n.uiText != nil) {
		return n.uiText.Text()
	}
	return n.Name
}

func (n *Node) AddComboBox(title string, s ...string) *Node {
	box := n.uiBox
	if (box == nil) {
		return n
	}

	ecbox := ui.NewEditableCombobox()

	for id, name := range s {
		log.Println("Adding Combobox Entry:", id, name)
		ecbox.Append(name)
	}

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		test := ecbox.Text()
		log.Println("text is now:", test)
	})

	box.Append(ecbox, false)

	newNode := n.AddNode(title)
	newNode.uiText = ecbox
	return newNode
}
