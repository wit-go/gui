/*
	A Labeled Single Line Entry widget:

	-----------------------------
	|            |              |
	|   Food:    |  <type here> |
	|            |              |
	-----------------------------
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui"
)

type BasicEntry struct {
	parent	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	v	*gui.Node	// value widget

	value	string
	label	string

	Custom func()
}

func (n *BasicEntry) Get() string {
	return n.value
}

func (n *BasicEntry) Set(value string) *BasicEntry {
	log.Println("BasicEntry.Set() =", value)
	if (n.v != nil) {
		n.v.Set(value)
	}
	n.value = value
	return n
}

func NewBasicEntry(p *gui.Node, name string) *BasicEntry {
	d := BasicEntry {
		parent: p,
		value: "",
	}

	// various timeout settings
	d.l = p.NewLabel(name)
	d.v = p.NewEntryLine("")
	d.v.Custom = func() {
		d.value = d.v.S
		log.Println("BasicEntry.Custom() user changed value to =", d.value)
	}

	return &d
}
