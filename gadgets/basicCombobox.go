/*
	A Labeled Combobox widget:

	-----------------------------
	|            |              |
	|   Food:    |  <dropdown>  |
	|            |              |
	-----------------------------

	The user can then edit the dropdown field and type anything into it
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui"
)

type BasicCombobox struct {
	ready	bool
	name	string

	parent	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	d	*gui.Node	// dropdown widget

	value	string
	label	string

	Custom func()
}

func (d *BasicCombobox) Get() string {
	if ! d.Ready() {return ""}
	return d.value
}

// Returns true if the status is valid
func (d *BasicCombobox) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *BasicCombobox) Add(value string) {
	if ! d.Ready() {return}
	log.Println("BasicCombobox.Add() =", value)
	d.d.AddDropdownName(value)
	return
}

func (d *BasicCombobox) Set(value string) bool {
	if ! d.Ready() {return false}
	log.Println("BasicCombobox.Set() =", value)
	d.d.SetText(value)
	d.value = value
	return true
}

func NewBasicCombobox(p *gui.Node, name string) *BasicCombobox {
	d := BasicCombobox {
		parent: p,
		name: name,
		ready: false,
	}

	// various timeout settings
	d.l = p.NewLabel(name)
	d.d = p.NewCombobox("")
	d.d.Custom = func() {
		d.value = d.Get()
		log.Println("BasicCombobox.Custom() user changed value to =", d.value)
		if d.Custom != nil {
			d.Custom()
		}
	}

	d.ready = true
	return &d
}
