package gui

import "log"
import "reflect"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"
// import "github.com/davecgh/go-spew/spew"


// This is the default mouse click handler
// Every mouse click that hasn't been assigned to
// something specific will fall into this routine
// By default, all it runs is the call back to
// the main program that is using this library
//
// This routine MUST be here as this is how the andlabs/ui works
// This is the raw routine passed to every button in andlabs libui / ui
//
// There is a []GuiButton which has all the buttons. We search
// for the button and then call the function below
//

func (n *Node) AddButton(name string, custom func(*Node)) *Node {
	if (n.uiBox == nil) {
		log.Println("gui.Node.AppendButton() filed node.UiBox == nil")
		return n
	}
	button := ui.NewButton(name)
	log.Println("reflect.TypeOF(uiBox) =", reflect.TypeOf(n.uiBox))
	log.Println("reflect.TypeOF(uiButton) =", reflect.TypeOf(button))
	// true == expand, false == make normal size button
	n.uiBox.Append(button, false)
	n.uiButton = button

	newNode := n.makeNode(name, 888, 888 + Config.counter)
	newNode.uiButton = button
	newNode.custom = custom

	button.OnClicked(func(*ui.Button) {
		log.Println("gui.AppendButton() Button Clicked. Running custom()")
		custom(newNode)
	})
	// panic("AppendButton")
	// time.Sleep(3 * time.Second)
	return newNode
}

func CreateFontButton(n *Node, action string) *GuiButton {
        // create a 'fake' button entry for the mouse clicks
	var newGB	GuiButton
	newGB.Name	= "FONT"
	newGB.FB	= ui.NewFontButton()
	// newGB.Box	= n.box
	Data.AllButtons	= append(Data.AllButtons, &newGB)

	newGB.FB.OnChanged(func (*ui.FontButton) {
		log.Println("FontButton.OnChanged() START mouseClick(&newBM)", newGB)
		if (Data.MouseClick != nil) {
			Data.MouseClick(&newGB)
		}
	})
	return &newGB
}

func CreateColorButton(n *Node, custom func(*GuiButton), name string, values interface {}) *GuiButton {
        // create a 'fake' button entry for the mouse clicks
	var newCB	GuiButton
	newCB.Name	= name
	newCB.CB	= ui.NewColorButton()
	// newCB.Box	= n.box
	newCB.Custom	= custom
	newCB.Values	= values

	Data.AllButtons	= append(Data.AllButtons, &newCB)

	newCB.CB.OnChanged(func (*ui.ColorButton) {
		log.Println("ColorButton.OnChanged() START Color Button Click")
		r, g, b, a := newCB.CB.Color()
		log.Println("ColorButton.OnChanged() Color() =", r, g, b, a)
		if (newCB.Custom != nil) {
			newCB.Custom(&newCB)
		} else if (Data.MouseClick != nil) {
			Data.MouseClick(&newCB)
		}
	})
	n.uiBox.Append(newCB.CB, false)
	return &newCB
}
