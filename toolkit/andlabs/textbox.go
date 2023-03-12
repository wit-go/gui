package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func newTextbox(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "NewTexbox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		log(debugError, "newTextbox() listMap()")
		log(debugError, "FFFFFFFFFFFF listMap()")
		log(debugError, "FFFFFFFFFFFF listMap()")
	}

	// t.NewTextbox(w)
// func (t *andlabsT) NewTextbox(w *toolkit.Widget) *andlabsT {
	var newt *andlabsT
	newt = new(andlabsT)

	log(debugToolkit, "gui.Toolkit.NewTextbox()", w.Name)
	if t.broken() {
		return
	}

	c := ui.NewNonWrappingMultilineEntry()
	newt.uiMultilineEntry = c

	newt.uiBox = t.uiBox
	newt.Name = w.Name
	newt.tw = w
	if (defaultBehavior) {
		t.uiBox.Append(c, true)
	} else {
		t.uiBox.Append(c, stretchy)
	}

	/*
	// don't bother with "images" on andlabs/ui
	"image"
	"bytes"
	_ "image/png"
	"image/draw"

	if (w.Name == "image") {
		log(true, "NewTextbox() trying to add a new image")
		i := ui.NewImage(16, 16)
		img, _, err := image.Decode(bytes.NewReader(rawImage))
		if err != nil {
			panic(err)
		}
		nr, ok := img.(*image.RGBA)
		if !ok {
			i2 := image.NewRGBA(img.Bounds())
			draw.Draw(i2, i2.Bounds(), img, img.Bounds().Min, draw.Src)
			nr = i2
		}
		i.Append(nr)
		t.uiBox.Append(i, true)

		var img *ui.Image
		var icon []byte
		var imgA image.Image

		icon, _ = res.ReadFile("resources/ping6.working.png")
		// imgA, _, err := image.Decode(bytes.NewReader(b))
		imgA, _, _ = image.Decode(icon)
		img.Append(imgA)
		img.Append(icon)
	}
	*/

	c.OnChanged(func(spin *ui.MultilineEntry) {
		w.S = newt.uiMultilineEntry.Text()
		// this is still dangerous
		// newt.commonChange(newt.tw)
		log(debugChange, "Not yet safe to trigger on change for ui.MultilineEntry")
	})
	mapWidgetsToolkits(w, newt)
}


func doTextbox(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newTextbox(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Textbox() ct.broken", ct)
		return
	}
	if (ct.uiMultilineEntry == nil) {
		log(debugError, "Textbox() uiMultilineEntry == nil", ct)
		return
	}
	// the dns control panel isn't crashing anymore (?)
	Queue(ct.doSimpleAction)
}

func (t *andlabsT) doSimpleAction() {
	if (t.tw == nil) {
		log(true, "doSimpleAction() got an empty widget")
		log(true, "THIS SHOULD NEVER HAPPEN")
		panic("crap. panic. widget == nil")
	}
	log(debugChange, "Going to attempt:", t.tw.Action)
	switch  t.tw.Action {
	case "Enable":
		if (t.uiEntry != nil) {
			t.uiEntry.Enable()
		} else if (t.uiMultilineEntry != nil) {
			t.uiMultilineEntry.Enable()
		} else {
			log(debugError, "don't know what to enable", t.Name)
		}
	case "Disable":
		if (t.uiEntry != nil) {
			t.uiEntry.Disable()
		} else if (t.uiMultilineEntry != nil) {
			t.uiMultilineEntry.Disable()
		} else {
			log(debugError, "don't know what to disable", t.Name)
		}
	case "Show":
		t.uiMultilineEntry.Show()
	case "Hide":
		t.uiMultilineEntry.Hide()
	case "SetText":
		t.uiMultilineEntry.SetText(t.tw.S)
	case "Set":
		t.uiMultilineEntry.SetText(t.tw.S)
	default:
		log(debugError, "Can't do", t.tw.Action, "to a Textbox")
	}
}
