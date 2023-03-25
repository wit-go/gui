package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// make new Image here
func newImage(a *toolkit.Action) {
	w := a.Widget
	log(debugToolkit, "newImage()", w.Name)

	t := andlabs[a.WhereId]
	if (t == nil) {
		log(debugToolkit, "newImage() toolkit struct == nil. name=", w.Name)
		listMap(debugToolkit)
	}
	newt := t.rawImage(w.Name)
	place(a, t, newt)
}

// make new Image using andlabs/ui
func (t *andlabsT) rawImage(title string) *andlabsT {
	var newt andlabsT
	var img *ui.Image
	newt.Name = title

	log(debugToolkit, "rawImage() create", newt.Name)

	img = ui.NewImage(16, 16)

	newt.uiImage = img
	// newt.uiControl = img

	return &newt
}
/*
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
