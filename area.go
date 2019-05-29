package gui

import "log"
import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func findFB(button *ButtonMap) *ButtonMap {
	var a *ButtonMap
	for key, foo := range Data.AllButtons {
		log.Println("defaultButtonClick() Data.AllButtons =", key, foo)
		// if Data.AllButtons[key] == *button {
		if &foo == button {
			a = &foo
		}
	}
	return a
}

func makeSplashArea() *AreaHandler {
	// make this button just to get the default font (but don't display the button)
	// There should be another way to do this (?)
	newB		:= CreateFontButton("AREA")

	time.Sleep(200 * time.Millisecond)
	tmp := findFB(newB)
	log.Println("makeSplashArea() newB =", newB)
	log.Println("makeSplashArea() newB.AH =", newB.AH)
	log.Println("makeSplashArea() newB =",	newB)
	newB.AH		= &myAH
	// log.Println("makeSplashArea() newB.AH =", newB.AH)
	log.Println("makeSplashArea() newB =", newB)

	time.Sleep(200 * time.Millisecond)
	tmp = findFB(newB)
	log.Println("makeSplashArea() tmp =", tmp, "newB", newB)

	myAH.Attrstr	= makeAttributedString()
	myAH.Area	= ui.NewArea(myAH)
	newB.A		= myAH.Area
	myAH.FontButton = newB.FB
	myAH.Button	= newB

	if (Data.Debug) {
		spew.Dump(myAH.Area)
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
		String:		ah.Attrstr,
		DefaultFont:	ah.FontButton.Font(),
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
		mouseClick(ah.Button)
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
