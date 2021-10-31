package main

import "log"

import "git.wit.org/wit/gui"

/*
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"
*/

func addDemoTab(n *gui.Node, title string) {
	newNode := n.AddTab(title, nil)
	if (gui.Config.Debug) {
		newNode.Dump()
	}
	newNode.ListChildren(false)

	groupNode1 := newNode.AddGroup("group 1")
	groupNode1.AddComboBox("demoCombo1", "foo", "bar", "stuff")
	groupNode1.AddComboBox("demoCombo3", "foo 3", "bar", "stuff")

	groupNode1.Dump()
	/*
	b := groupNode1.FindBox()
	b.Dump()
	*/
	// n1, b1 := addButton(groupNode1, "Data.ListChildren(false)")
/*
	b1.OnClicked(func(*ui.Button) {
		gui.Data.ListChildren(false)
	})
*/

	//n2, b2 := addButton(groupNode1, "dumpBox(window)")
	newNode2 := groupNode1.AppendButton("foo 3 " + "AppendButton()", func(groupNode1 *gui.Node) {
		log.Println("Dumping groupNode1")
		groupNode1.Dump()
	})
	newNode2.Dump()
/*
	b2.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("names[x] =", names[x])
		x.Dump(names[x])
	})
	n2.Dump()
*/

	groupNode2 := newNode.AddGroup("group 2")
	groupNode2.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
}

/*
func addButton(n *gui.Node, name string) (*gui.Node, *ui.Button) {
	// val    := &myButtonInfo{}
	button := ui.NewButton(name)
	// val.uiC = button

	button.OnClicked(func(*ui.Button) {
		log.Println("Should do something here")
	})

	// n.Append(button, false)
	newNode := n.AppendButton(name + "AppendButton", func() {
		log.Println("Should do something here also")
	})
	return newNode, button
}
*/

/*
type myButtonInfo struct {
	Custom		func (*gui.GuiButton)
	ADD		func (*gui.GuiButton)
	Name		string
	Action		string
	Node		*gui.Node
}

func newMakeButton(n *gui.Node, name string, action string, custom func(*gui.GuiButton)) *gui.Node {
	val          := &myButtonInfo{}
	val.Custom   = custom
	val.Name     = name
	val.Node     = n
	// val.Action   = action
	return n.CreateButton(custom, name, val)
}
*/
