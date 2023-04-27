package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// make new Image using andlabs/ui
func (p *node) newImage(n *node) {
	newt := new(andlabsT)
	var img *ui.Image

	log(debugToolkit, "rawImage() create", n.Name)

	img = ui.NewImage(16, 16)

	newt.uiImage = img
	// newt.uiControl = img

	n.tk = newt
	p.place(n)
}
/*
	if (a.Name == "image") {
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
