package main

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

/*
	This is a code example taken directly from the toolkit andlabs/ui

	This code is here to double check that the toolkit itself still works
	the same way. This is intended as a sanity check.
*/

func BlankWindow(w *ui.Window) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	w.SetChild(hbox)
	return hbox
}

func (t *guiWidget) DemoNumbersPage() {
	var w *ui.Window

	log(debugToolkit, "Starting wit/gui toolkit andlabs/ui DemoNumbersPage()")

	w = t.uiWindow
	t.uiBox = makeNumbersPage()
	t.uiBox.SetPadded(true)
	w.SetChild(t.uiBox)
	w.SetTitle("Internal demo of andlabs/ui toolkit")
}

func makeNumbersPage() *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)
	pbar := ui.NewProgressBar()
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

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}
