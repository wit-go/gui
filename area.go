package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func makeSplashArea() *AreaHandler {
	// make this button just to get the default font (but don't display the button)
	// There should be another way to do this (?)
	Data.fontButton = CreateFontButton("SplashFont", "DONE")

	myAH.Attrstr	= makeAttributedString()
	Data.splashArea = ui.NewArea(myAH)
	Data.MyArea     = Data.splashArea
	myAH.Area	= Data.splashArea

	// create a 'fake' button entry for the mouse clicks
	var newmap ButtonMap
	newmap.Action	= "AREA"
	newmap.AH	= &myAH
	newmap.A	= Data.splashArea
	myAH.Button	= &newmap
	Data.AllButtons = append(Data.AllButtons, newmap)

	if (Data.Debug) {
		spew.Dump(Data.splashArea)
		log.Println("DEBUGGING", Data.Debug)
	} else {
		log.Println("NOT DEBUGGING AREA mhAH.Button =", myAH.Button)
	}
	return &myAH
}

func appendWithAttributes(newText *ui.AttributedString, what string, attrs ...ui.Attribute) {
	start := len(newText.String())
	end := start + len(what)
	newText.AppendUnattributed(what)
	for _, a := range attrs {
		newText.SetAttribute(a, start, end)
	}
}

func (ah AreaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:		myAH.Attrstr,
		DefaultFont:	Data.fontButton.Font(),
		Width:		p.AreaWidth,
		Align:		ui.DrawTextAlign(1),
	})
	p.Context.Text(tl, 0, 0)
	defer tl.Free()
}

func (ah AreaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	if (Data.Debug) {
		log.Println("GOT MouseEvent()")
		spew.Dump(me)
	}
	if (me.Down == 1) {
		log.Println("GOT MOUSE DOWN")
	}
	if (me.Up == 1) {
		log.Println("GOT MOUSE UP")
		log.Println("GOT MOUSE UP")
		log.Println("GOT MOUSE UP")
		mouseClick(myAH.Button)
	}
}

func (ah AreaHandler) MouseCrossed(a *ui.Area, left bool) {
	log.Println("GOT MouseCrossed()")
}

func (ah AreaHandler) DragBroken(a *ui.Area) {
	log.Println("GOT DragBroken()")
}

func (ah AreaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
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
