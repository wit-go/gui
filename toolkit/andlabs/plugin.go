package main

// if you include more than just this import
// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
import "git.wit.org/wit/gui/toolkit"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"

//
// This should be called ?
// Pass() ?
// This handles all interaction between the wit/gui package (what golang knows about)
// and this plugin that talks to the OS and does scary and crazy things to make
// a GUI on whatever OS or whatever GUI toolkit you might have (GTK, QT, WASM, libcurses)
//
// Once you are here, you should be in a protected goroutine created by the golang wit/gui package
//
// TODO: make sure you can't escape this goroutine
//
func Send(p *toolkit.Widget, c *toolkit.Widget) {
	log(debugPlugin, "Send() goodbye. not used anymore")
}

func Action(a *toolkit.Action) {
	if (a == nil) {
		log(debugPlugin, "Action = nil")
		return
	}
	f := func() {
		rawAction(a)
	}

	if (callback == nil) {
		if (a.Callback != nil) {
			log(debugNow, "setting Callback", a.Callback)
			callback = a.Callback
		}
	}

	// f()
	Queue(f)
}

func rawAction(a *toolkit.Action) {

	log(debugAction, "Action() START a.Type =", a.Type)
	log(debugAction, "Action() START a.S =", a.S)
	log(debugAction, "Action() START a.Widget =", a.Widget)

	log(logInfo, "Action() START a.WidgetId =", a.WidgetId, "a.WhereId =", a.WhereId)

	switch a.Type {
	case toolkit.Add:
		add(a)
	case toolkit.Show:
		a.Widget.B = true
		show(a)
	case toolkit.Hide:
		a.Widget.B = false
		show(a)
	case toolkit.Enable:
		a.Widget.B = true
		enable(a)
	case toolkit.Disable:
		a.Widget.B = false
		enable(a)
	case toolkit.Get:
		setText(a)
	case toolkit.GetText:
		switch a.Widget.Type {
		case toolkit.Textbox:
			t := andlabs[a.WidgetId]
			a.S = t.s
		}
	case toolkit.Set:
		setText(a)
	case toolkit.SetFlag:
		flag(a)
	case toolkit.SetText:
		setText(a)
	case toolkit.AddText:
		setText(a)
	case toolkit.Margin:
		pad(a)
	case toolkit.Unmargin:
		pad(a)
	case toolkit.Pad:
		pad(a)
	case toolkit.Unpad:
		pad(a)
	case toolkit.Delete:
		uiDelete(a)
	case toolkit.Flag:
		flag(a)
	case toolkit.Move:
		log(debugNow, "attempt to move() =", a.Type, a.Widget)
		move(a)
	default:
		log(debugError, "Action() Unknown =", a.Type, a.Widget)
	}
	log(debugAction, "Action() END =", a.Type, a.Widget)
}

func flag(a *toolkit.Action) {
	// log(debugFlags, "plugin Send() flag parent =", p.Name, p.Type)
	// log(debugFlags, "plugin Send() flag child  =", c.Name, c.Type)
	// log(debugFlags, "plugin Send() flag child.Action  =", c.Action)
	// log(debugFlags, "plugin Send() flag child.S  =", c.S)
	// log(debugFlags, "plugin Send() flag child.B  =", c.B)
	// log(debugFlags, "plugin Send() what to flag?")
	// should set the checkbox to this value
	switch a.S {
	case "Toolkit":
		debugToolkit = a.B
	case "Change":
		debugChange = a.B
	case "Plugin":
		debugPlugin = a.B
	case "Flags":
		debugFlags = a.B
	case "Error":
		debugError = a.B
	case "Now":
		debugNow = a.B
	case "Show":
		ShowDebug()
	default:
		log(debugError, "Can't set unknown flag", a.S)
	}
}

func setText(a *toolkit.Action) {
	t := andlabs[a.WidgetId]
	if (t == nil) {
		log(debugError, "setText error. andlabs[id] == nil", a.WidgetId)
		actionDump(debugError, a)
		return
	}
	log(debugChange, "setText() Attempt on", t.Type, "with", a.S)

	switch t.Type {
	case toolkit.Window:
		t.uiWindow.SetTitle(a.S)
	case toolkit.Tab:
	case toolkit.Group:
		t.uiGroup.SetTitle(a.S)
	case toolkit.Checkbox:
		switch a.Type {
		case toolkit.SetText:
			t.uiCheckbox.SetText(a.S)
		case toolkit.Get:
			t.tw.B = t.uiCheckbox.Checked()
		case toolkit.Set:
			// TODO: commented out while working on chan
			// t.uiCheckbox.SetChecked(a.B)
			t.tw.B = a.B
		default:
			log(debugError, "setText() unknown", a.Type, "on checkbox", t.tw.Name)
		}
	case toolkit.Textbox:
		switch a.Type {
		case toolkit.Set:
			t.uiMultilineEntry.SetText(a.S)
		case toolkit.SetText:
			t.uiMultilineEntry.SetText(a.S)
		case toolkit.Get:
			t.tw.S = t.s
		case toolkit.GetText:
			t.tw.S = t.s
		default:
			log(debugError, "setText() unknown", a.Type, "on checkbox", t.tw.Name)
		}
	case toolkit.Label:
		t.uiLabel.SetText(a.S)
	case toolkit.Button:
		t.uiButton.SetText(a.S)
	case toolkit.Slider:
		switch a.Type {
		case toolkit.Get:
			t.tw.I = t.uiSlider.Value()
		case toolkit.Set:
			t.uiSlider.SetValue(a.I)
		default:
			log(debugError, "setText() unknown", a.Type, "on checkbox", t.tw.Name)
		}
	case toolkit.Spinner:
		switch a.Type {
		case toolkit.Get:
			t.tw.I = t.uiSpinbox.Value()
		case toolkit.Set:
			t.uiSpinbox.SetValue(a.I)
		default:
			log(debugError, "setText() unknown", a.Type, "on checkbox", t.tw.Name)
		}
	case toolkit.Dropdown:
		switch a.Type {
		case toolkit.AddText:
			AddDropdownName(a)
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
					log(debugChange, "setText() Dropdown worked.", t.tw.S)
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
			t.tw.S = t.s
		case toolkit.GetText:
			t.tw.S = t.s
		default:
			log(debugError, "setText() unknown", a.Type, "on checkbox", t.tw.Name)
		}
	case toolkit.Combobox:
		switch a.Type {
		case toolkit.AddText:
			t.AddComboboxName(a.S)
		case toolkit.Set:
			t.uiEditableCombobox.SetText(a.S)
			t.s = a.S
		case toolkit.SetText:
			t.uiEditableCombobox.SetText(a.S)
			t.s = a.S
		case toolkit.Get:
			t.tw.S = t.s
		case toolkit.GetText:
			t.tw.S = t.s
		default:
			log(debugError, "setText() unknown", a.Type, "on checkbox", t.tw.Name)
		}
	default:
		log(debugError, "plugin Send() Don't know how to setText on", t.tw.Type, "yet", a.Type)
	}
}
