/*
	A Labeled label:

	-----------------------------
	|            |              |
	|   Food:    |    Apple     |
	|            |              |
	-----------------------------
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui"
)

type Node gui.Node

type BasicLabel struct {
	p	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	v	*gui.Node	// value widget

	value	string
	label	string

	Custom func()
}

func (n *BasicLabel) Get() string {
	return n.value
}

func (n *BasicLabel) Set(value string) *BasicLabel {
	log.Println("BasicLabel.Set() =", value)
	if (n.v != nil) {
		n.v.Set(value)
	}
	n.value = value
	return n
}

func (ngui *Node) NewBasicLabel(name string) *BasicLabel {
	var n *gui.Node
	n = (*gui.Node)(ngui)
	d := BasicLabel {
		p: n,
		value: "",
	}

	// various timeout settings
	d.l = n.NewLabel(name)
	d.v = n.NewLabel("")
	d.v.Custom = func() {
		d.value = d.v.S
		log.Println("BasicLabel.Custom() user changed value to =", d.value)
	}

	return &d
}
