package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// This will create a tab in a window using direct
// calls to andlabs/ui. This can be used to bypass
// the obvuscation added in this package if it is desired
// or needed.
func (n *Node) AddDemoAndlabsUiTab(title string) {
	newNode := n.AddTab(title, makeAndlabsUiTab())
	if (Config.DebugNode) {
		newNode.Dump()
	}
	tabSetMargined(newNode.uiTab)
}

func makeAndlabsUiTab() *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(47, 100)
	slider  := ui.NewSlider(21, 100)
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

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		log.Println("test")
		test := ecbox.Text()
		log.Println("test=", test)
	})

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}
