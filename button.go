package gui

import "log"
import "reflect"
// import "image/color"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"
// import "github.com/davecgh/go-spew/spew"


// TODO: bring this generic mouse click function back
//
// This is the default mouse click handler
// Every mouse click that hasn't been assigned to
// something specific will fall into this routine
// By default, all it runs is the call back to
// the main program that is using this library
//
// This routine MUST be here as this is how the andlabs/ui works
// This is the raw routine passed to every button in andlabs libui / ui
//
//

func (n *Node) AddButton(name string, custom func(*Node)) *Node {
	if (n.uiBox == nil) {
		log.Println("gui.Node.AppendButton() filed node.UiBox == nil")
		return n
	}
	button := ui.NewButton(name)
	if (Config.Debug) {
		log.Println("reflect.TypeOF(uiBox) =", reflect.TypeOf(n.uiBox))
		log.Println("reflect.TypeOF(uiButton) =", reflect.TypeOf(button))
	}
	// true == expand, false == make normal size button
	n.uiBox.Append(button, Config.Stretchy)
	n.uiButton = button

	newNode := n.makeNode(name, 888, 888 + Config.counter)
	newNode.uiButton = button
	newNode.custom = custom

	button.OnClicked(func(*ui.Button) {
		log.Println("gui.AppendButton() Button Clicked. Running custom()")
		custom(newNode)
	})
	return newNode
}

func (n *Node) CreateFontButton(action string) *Node {
	n.uiFontButton	= ui.NewFontButton()

	n.uiFontButton.OnChanged(func (*ui.FontButton) {
		log.Println("FontButton.OnChanged() START")
		n.Dump()
	})
	n.uiBox.Append(n.uiFontButton, Config.Stretchy)
	return n
}

func (n *Node) CreateColorButton(custom func(*Node), name string, values interface {}) *Node {
        // create a 'fake' button entry for the mouse clicks
	n.uiColorButton	= ui.NewColorButton()
	n.custom	= custom
	n.values	= values

	n.uiColorButton.OnChanged(func (*ui.ColorButton) {
		log.Println("ColorButton.OnChanged() START Color Button Click")
		rgba := n.Color
		r, g, b, a := rgba.R, rgba.G, rgba.B, rgba.A
		log.Println("ColorButton.OnChanged() Color() =", r, g, b, a)
		if (n.custom != nil) {
			n.custom(n)
		} else if (Data.MouseClick != nil) {
			Data.MouseClick(n)
		}
	})
	n.uiBox.Append(n.uiColorButton, Config.Stretchy)
	return n
}
