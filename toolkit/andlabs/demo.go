package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func BlankWindow(w *ui.Window) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	w.SetChild(hbox)
	return hbox
}

func (t *Toolkit) DemoNumbersPage() {
	var w *ui.Window

	w = t.uiWindow
	t.uiBox = makeNumbersPage()
	t.uiBox.SetPadded(true)
	w.SetChild(t.uiBox)
	w.SetTitle("Internal demo of andlabs/ui toolkit")

	if (DebugToolkit) {
		log.Println("gui.Toolbox.DemoNumbersPage()")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
}

/*
func Demo(b *ui.Box) *Toolkit {
	x := 22
	y := 33

	// make new node here
	log.Println("gui.Toolbox.NewSpinbox()", x, y)
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewSpinbox() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return nil
	}
	s := ui.NewSlider(x, y)
	t.uiSlider = s
	t.uiBox = b
	t.uiBox.Append(s, false)

	s.OnChanged(func(spin *ui.Slider) {
		i := spin.Value()
		if (DebugToolkit) {
			log.Println("gui.Toolbox.ui.OnChanged() val =", i)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(t)
		}
		if (t.OnChanged != nil) {
			log.Println("gui.Toolbox.OnChanged() entered val =", i)
			t.OnChanged(&t)
		}
	})

	return &t
}
*/

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
