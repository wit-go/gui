package main

import (
	"go.wit.com/gui/toolkit"
)

func (n *node) setText(a *toolkit.Action) {
	log(debugChange, "setText() START with a.S =", a.S)
	t := n.tk
	if (t == nil) {
		log(debugError, "setText error. tk == nil", n.Name, n.WidgetId)
		actionDump(debugError, a)
		return
	}
	log(debugChange, "setText() Attempt on", n.WidgetType, "with", a.S)

	switch n.WidgetType {
	case toolkit.Window:
		t.uiWindow.SetTitle(a.S)
	case toolkit.Tab:
	case toolkit.Group:
		t.uiGroup.SetTitle(a.S)
	case toolkit.Checkbox:
		switch a.ActionType {
		case toolkit.SetText:
			t.uiCheckbox.SetText(a.S)
		case toolkit.Get:
			n.B = t.uiCheckbox.Checked()
		case toolkit.Set:
			// TODO: commented out while working on chan
			t.uiCheckbox.SetChecked(a.B)
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case toolkit.Textbox:
		switch a.ActionType {
		case toolkit.Set:
			if (t.uiEntry != nil) {
				t.uiEntry.SetText(a.S)
			}
			if (t.uiMultilineEntry != nil) {
				t.uiMultilineEntry.SetText(a.S)
			}
		case toolkit.SetText:
			if (t.uiEntry != nil) {
				t.uiEntry.SetText(a.S)
			}
			if (t.uiMultilineEntry != nil) {
				t.uiMultilineEntry.SetText(a.S)
			}
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case toolkit.Label:
		t.uiLabel.SetText(a.S)
	case toolkit.Button:
		t.uiButton.SetText(a.S)
	case toolkit.Slider:
		switch a.ActionType {
		case toolkit.Get:
			n.I = t.uiSlider.Value()
		case toolkit.Set:
			t.uiSlider.SetValue(a.I)
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case toolkit.Spinner:
		switch a.ActionType {
		case toolkit.Get:
			n.I = t.uiSpinbox.Value()
		case toolkit.Set:
			t.uiSpinbox.SetValue(a.I)
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case toolkit.Dropdown:
		switch a.ActionType {
		case toolkit.AddText:
			n.AddDropdownName(a.S)
		case toolkit.Set:
			var orig int
			var i int = -1
			var s string
			orig = t.uiCombobox.Selected()
			log(debugChange, "try to set the Dropdown to", a.S, "from", orig)
			// try to find the string
			for i, s = range t.val {
				log(debugChange, "i, s", i, s)
				if (a.S == s) {
					t.uiCombobox.SetSelected(i)
					log(debugChange, "setText() Dropdown worked.", n.S)
					return
				}
			}
			log(debugError, "setText() Dropdown did not find:", a.S)
			// if i == -1, then there are not any things in the menu to select
			if (i == -1) {
				return
			}
			// if the string was never set, then set the dropdown to the last thing added to the menu
			if (orig == -1) {
				t.uiCombobox.SetSelected(i)
			}
		case toolkit.Get:
			// t.S = t.s
		case toolkit.GetText:
			// t.S = t.s
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case toolkit.Combobox:
		switch a.ActionType {
		case toolkit.AddText:
			t.AddComboboxName(a.S)
		case toolkit.Set:
			t.uiEditableCombobox.SetText(a.S)
			n.S = a.S
		case toolkit.SetText:
			t.uiEditableCombobox.SetText(a.S)
			n.S = a.S
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	default:
		log(debugError, "plugin Send() Don't know how to setText on", n.WidgetType, "yet", a.ActionType)
	}
	log(debugChange, "setText() END with a.S =", a.S)
}
