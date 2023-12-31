/*
	A Labeled Dropdown widget:

	-----------------------------
	|            |              |
	|   Food:    |  <dropdown>  |
	|            |              |
	-----------------------------

	This being a 'Basic Dropdown', the dropdown names must be unique
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui"
)

type BasicDropdown struct {
	ready	bool
	name	string

	parent	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	d	*gui.Node	// dropdown widget

	value	string
	label	string

	Custom func()
}

func (d *BasicDropdown) Get() string {
	if ! d.Ready() {return ""}
	return d.d.GetText()
}

// Returns true if the status is valid
func (d *BasicDropdown) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *BasicDropdown) Add(value string) {
	if ! d.Ready() {return}
	log.Println("BasicDropdown.Set() =", value)
	d.d.AddDropdownName(value)
	return
}

func (d *BasicDropdown) Set(value string) bool {
	if ! d.Ready() {return false}
	log.Println("BasicDropdown.Set() =", value)
	d.l.SetText(value)
	d.value = value
	return true
}

func NewBasicDropdown(p *gui.Node, name string) *BasicDropdown {
	d := BasicDropdown {
		parent: p,
		name: name,
		ready: false,
	}

	// various timeout settings
	d.l = p.NewLabel(name)
	d.d = p.NewDropdown("")
	d.d.Custom = func() {
		d.value = d.Get()
		log.Println("BasicDropdown.Custom() user changed value to =", d.value)
		if d.Custom != nil {
			d.Custom()
		}
	}

	d.ready = true
	return &d
}
