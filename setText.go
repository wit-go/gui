package gui

// Common actions for widgets like 'Enable' or 'Hide'

import (
	"go.wit.com/log"
	"go.wit.com/gui/widget"
)

func (n *Node) SetText(text string) *Node {
	if ! n.Ready() { return n }
	log.Log(CHANGE, "SetText() value =", text)

	n.value = text

	if ! n.hidden {
		a := newAction(n, widget.SetText)
		a.Value = n.value
		sendAction(a)
	}
	return n
}

/*
func convertString(val any) string {
	switch v := val.(type) {
	case bool:
		n.B = val.(bool)
	case string:
		n.label = val.(string)
		n.S = val.(string)
	case int:
		n.I = val.(int)
	default:
		log.Error(errors.New("Set() unknown type"), "v =", v)
	}
}
*/


func (n *Node) Set(val any) {
	log.Log(CHANGE, "Set() value =", val)

	n.value = val
	/*
	n.value = val

	switch v := val.(type) {
	case bool:
		n.B = val.(bool)
	case string:
		n.label = val.(string)
		n.S = val.(string)
	case int:
		n.I = val.(int)
	default:
		log.Error(errors.New("Set() unknown type"), "v =", v)
	}
	*/

	if ! n.hidden {
		a := newAction(n, widget.Set)
		a.Value = n.value
		sendAction(a)
	}
}
