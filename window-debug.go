package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var names = make([]string, 100)

func makeWindowDebug() ui.Control {

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(22, 44)
	slider  := ui.NewSlider(22, 44)
	pbar    := ui.NewProgressBar()

	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("WindowMap")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	addName(cbox, "Window 1")
	addName(cbox, "Window 2")
	addName(cbox, "Combobox Item 3")
	vbox.Append(cbox, false)

	cbox.OnSelected(func(*ui.Combobox) {
		log.Println("test")
		test := cbox.Selected()
		log.Println("test=", test)
		log.Println("names[test] =", names[test])

//		for name := range names {
//			log.Println("gui.DumpBoxes() name: ", name)
//		}
//		if (names[test] != nil) {
//		}
	})

	for name, _ := range Data.WindowMap {
		log.Println("gui.DumpBoxes() name: ", name)
		addName(cbox, name)
	}

	return hbox
}

var x int = 0

func addName(c *ui.Combobox, s string) {
	c.Append(s)
	names[x] = s
	x = x + 1
}
