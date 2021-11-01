package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func makeGenericArea(n *Node, newText *ui.AttributedString, custom func(*GuiButton)) {
	// make this button just to get the default font (but don't display the button)
	// There should be another way to do this (?)
	var newB *GuiButton
	newB		= CreateFontButton(n, "AREA")
	newB.Custom	= custom

/*
	gw := n.window
	// initialize the GuiArea{}
	gw.Area.Button		= newB
*/
	area			:= new(GuiArea)

	// gw.Area.Box		= gb
	n.uiAttrstr		= newText
	n.uiArea		= ui.NewArea(area)

	if (Config.Debug) {
		spew.Dump(n.uiArea)
		log.Println("DEBUGGING", Config.Debug)
	} else {
		log.Println("NOT DEBUGGING AREA mhAH.Button =", n.uiButton)
	}
}

func AreaAppendText(newText *ui.AttributedString, what string, attrs ...ui.Attribute) {
	start := len(newText.String())
	end := start + len(what)
	newText.AppendUnattributed(what)
	for _, a := range attrs {
		newText.SetAttribute(a, start, end)
	}
}

func appendWithAttributes(newText *ui.AttributedString, what string, attrs ...ui.Attribute) {
	start := len(newText.String())
	end := start + len(what)
	newText.AppendUnattributed(what)
	for _, a := range attrs {
		newText.SetAttribute(a, start, end)
	}
}

func (ah GuiArea) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:		ah.UiAttrstr,
		DefaultFont:	ah.Button.FB.Font(),
		Width:		p.AreaWidth,
		Align:		ui.DrawTextAlign(1),
	})
	p.Context.Text(tl, 0, 0)
	defer tl.Free()
}

func (ah GuiArea) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	if (Config.Debug) {
		log.Println("GOT MouseEvent() ah.Button =", ah.Button)
		spew.Dump(me)
	}
	if (me.Down == 1) {
		log.Println("GOT MOUSE DOWN")
		log.Println("GOT MOUSE DOWN ah.Button =", ah.Button)
		log.Println("GOT MOUSE UP ah.Button.FB =", ah.Button.FB)
	}
	if (me.Up == 1) {
		log.Println("GOT MOUSE UP")
		log.Println("GOT MOUSE UP ah.Button =", ah.Button)
		log.Println("GOT MOUSE UP ah.Button.FB =", ah.Button.FB)
		if (ah.Button.Custom != nil) {
			ah.Button.Custom(ah.Button)
		} else if (Data.MouseClick != nil) {
			Data.MouseClick(ah.Button)
		}
	}
}

func (ah GuiArea) MouseCrossed(a *ui.Area, left bool) {
	log.Println("GOT MouseCrossed()")
}

func (ah GuiArea) DragBroken(a *ui.Area) {
	log.Println("GOT DragBroken()")
}

// TODO: fix KeyEvents
func (ah GuiArea) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	log.Println("GOT KeyEvent()")
	if (ke.Key == 10) {
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
	}
	if (ke.Key == 32) {
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
	}
	spew.Dump(ke)
	return false
}

func (n *Node) ShowTextBox(newText *ui.AttributedString, custom func(*GuiButton), name string) {
	log.Println("ShowTextBox() START")

	gw := n.Window
	if (gw == nil) {
		log.Println("ShowTextBox() ERROR gw = nil")
		return
	}
	log.Println("ShowTextBox() START gw =", gw)

	// TODO: allow padded & axis here
	n.uiBox.SetPadded(true)

	// add(gw.BoxMap["MAINBOX"], newbox)

	makeGenericArea(n, newText, custom)
	n.uiBox.Append(n.area.UiArea, true)
}
