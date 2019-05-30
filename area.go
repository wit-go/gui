package gui

import "log"
// import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func findFB(button *GuiButton) *GuiButton {
	var a *GuiButton
	for key, foo := range Data.AllButtons {
		log.Println("findFB() Data.AllButtons key, foo=", key, foo)
		if foo == button {
			log.Println("findFB() FOUND BUTTON key, foo=", key, foo)
			a = foo
		}
	}
	return a
}

func makeSplashArea(wm *GuiWindow, ah *GuiArea) {
	// make this button just to get the default font (but don't display the button)
	// There should be another way to do this (?)
	newB := CreateFontButton(wm, "AREA")

	// ah.UiAttrstr	= makeAttributedString()
	ah.UiArea	= ui.NewArea(ah)
	newB.A		= ah.UiArea
	newB.WM	= wm
	// Data.AllButtons[1].A = ah.UiArea
	// ah.Button	= &Data.AllButtons[1]
	ah.Button	= newB

	if (Data.Debug) {
		spew.Dump(ah.UiArea)
		log.Println("DEBUGGING", Data.Debug)
	} else {
		log.Println("NOT DEBUGGING AREA mhAH.Button =", ah.Button)
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
	if (Data.Debug) {
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
		mouseClick(ah.Button)
	}
}

func (ah GuiArea) MouseCrossed(a *ui.Area, left bool) {
	log.Println("GOT MouseCrossed()")
}

func (ah GuiArea) DragBroken(a *ui.Area) {
	log.Println("GOT DragBroken()")
}

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
	// splashWin.Destroy()
	// ui.Quit()
	return false
}
