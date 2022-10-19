package gui

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (n *Node) AddBox(axis int, name string) *Node {
	newBox		:= new(GuiBox)
	newBox.Window	= n.window
	newBox.Name	= name

	if (n.box == nil) {
		n.box = newBox
	}

	// make a new box & a new node
	newNode := n.New(name)
	newNode.box = newBox
	Config.counter += 1

	var uiBox *ui.Box
	if (axis == Xaxis) {
		uiBox = ui.NewHorizontalBox()
	} else {
		uiBox = ui.NewVerticalBox()
	}
	uiBox.SetPadded(true)
	newBox.UiBox = uiBox
	newNode.uiBox = uiBox

	n.Append(newNode)
	// add(n.box, newBox)
	return newNode
}

func HorizontalBreak(box *GuiBox) {
	log.Println("VerticalSeparator added to box =", box.Name)
	tmp := ui.NewHorizontalSeparator()
	if (box == nil) {
		return
	}
	if (box.UiBox == nil) {
		return
	}
	box.UiBox.Append(tmp, false)
}

func VerticalBreak(box *GuiBox) {
	log.Println("VerticalSeparator  added to box =", box.Name)
	tmp := ui.NewVerticalSeparator()
	box.UiBox.Append(tmp, false)
}

func (n *Node) AddComboBox(title string, s ...string) *Node {
	if (n.Toolkit == nil) {
		log.Println("AddComboBox.Toolkit is nil", title, s)
		n.Dump()
		os.Exit(0)
	}
	if (n.uiBox == nil) {
		log.Println("AddComboBox.uiBox is nil", title, s)
		n.Toolkit.Dump()
		n.uiBox = n.Toolkit.GetBox()
		// os.Exit(0)
		// return n
	}
	box := n.uiBox

	newNode := n.New(title)
	ecbox := ui.NewEditableCombobox()

	for id, name := range s {
		log.Println("Adding Combobox Entry:", id, name)
		ecbox.Append(name)
	}

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		test := ecbox.Text()
		log.Println("node.Name = '" + n.Name + "' text for '" + title + "' is now: '" + test + "'")
		log.Println("need to call node.OnChanged() here")
		if (newNode.OnChanged == nil) {
			log.Println("node.OnChanged() is nil")
			log.Println("need to call node.OnChanged() here", newNode.OnChanged)
			newNode.Dump()
		} else {
			log.Println("need to call node.OnChanged() here", newNode.OnChanged)
			newNode.OnChanged(newNode)
		}
	})

	box.Append(ecbox, false)

	newNode.uiText = ecbox
	return newNode
}

func (n *Node) GetText() string {
	if (n.uiText == nil) {
		return ""
	}
	ecbox := n.uiText
	return ecbox.Text()
}
