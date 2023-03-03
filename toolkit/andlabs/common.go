package main

import (
	"git.wit.org/wit/gui/toolkit"
	"github.com/davecgh/go-spew/spew"
)

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func init() {
	// Can you pass values to a plugin init() ? Otherwise, there is no way to safely print
	// log(debugToolkit, "gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)
}

func (t *andlabsT) commonChange(tw *toolkit.Widget) {
	log(debugChange, "commonChange() START widget   =", t.Name, t.Type)
	if (tw == nil) {
		log(true, "commonChange() What the fuck. there is no widget t.tw == nil")
		return
	}
	if (tw.Custom == nil) {
		log(debugChange, "commonChange() END    Widget.Custom() = nil", t.tw.Name, t.tw.Type)
		return
	}
	tw.Custom()
	log(debugChange, "commonChange() END   Widget.Custom()", t.tw.Name, t.tw.Type)
}

// does some sanity checks on the internal structs of the binary tree
// TODO: probably this should not panic unless it's running in devel mode (?)
// TODO: redo this now that WidgetType is used and send() is used to package plugins
func (t *andlabsT) broken() bool {
	if (t.parent != nil) {
		return false
	}
	if (t.uiBox == nil) {
		if (t.uiWindow != nil) {
			log(debugToolkit, "UiBox == nil. This is an empty window. Try to add a box")
			t.newBox()
			return false
		}
		log(true, "UiBox == nil. I can't add a widget without a place to put it")
		// log(debugToolkit, "probably could just make a box here?")
		// corruption or something horrible?
		t.Dump(true)
		panic("wit/gui toolkit/andlabs func broken() invalid goroutine access into this toolkit?")
		panic("wit/gui toolkit/andlabs func broken() this probably should not cause the app to panic here (?)")
		return true
	}
	if (t.uiWindow == nil) {
		log(debugToolkit, "UiWindow == nil. I can't add a widget without a place to put it (IGNORING FOR NOW)")
		t.Dump(debugToolkit)
		return false
	}
	return false
}
func broken(w *toolkit.Widget) bool {
	if (w == nil) {
		log(true, "widget == nil. I can't do anything widget")
		return true
	}
	return false
}

func dump(p *toolkit.Widget, c *toolkit.Widget, b bool) {
	log(b, "Parent:")
	pt := mapToolkits[p]
	if (pt == nil) {
		log(b, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	pt.Dump(b)

	log(b, "Child:")
	ct := mapToolkits[c]
	if (ct == nil) {
		log(b, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	ct.Dump(b)
}

func setMarginNew(w *toolkit.Widget, b bool) {
	wt := mapToolkits[w]
	log(true, "START setMarginNew", w.Name)
	if (wt == nil) {
		return
	}
	if (wt.uiGroup != nil) {
		log(true, "uiGroup.SetMargined(true)")
		wt.uiGroup.SetMargined(b)
	}
	if (wt.uiTab != nil) {
		i := wt.uiTab.NumPages()
		log(true, "tab.NumPages() =", i)
		for i > 0 {
			i -= 1
			log(true, "uiTab.SetMargined(true) for i =", i)
			wt.uiTab.SetMargined(i, b)
		}
	} else {
		log(true, "no uitab")
	}
	if (wt.uiWindow != nil) {
		log(true, "uiWindow.SetMargined(true)")
		wt.uiWindow.SetMargined(b)
	}
	log(true, "END   setMarginNew", w.Name)
}

func setMargin(p *toolkit.Widget, c *toolkit.Widget, b bool) {
	log(true, "Starting to implement SetMargin here")
	dump(p, c, true)

	setMarginNew(c, b)
	setMarginNew(p, b)
}

func (t *andlabsT) String() string {
	return t.GetText()
}

func (t *andlabsT) GetText() string {
	log(debugToolkit, "GetText() Enter debugToolkit=", debugToolkit)
	if (t.uiEntry != nil) {
		log(debugToolkit, "uiEntry.Text() =", t.uiEntry.Text())
		return t.uiEntry.Text()
	}
	if (t.uiMultilineEntry != nil) {
		log(debugToolkit, "uiMultilineEntry.Text() =", t.uiMultilineEntry.Text())
		text := t.uiMultilineEntry.Text()
		log(debugToolkit, "uiMultilineEntry.Text() =", text)
		t.text = text
		return text
	}
	if (t.uiCombobox != nil) {
		log(debugToolkit, "uiCombobox() =", t.text)
		return t.text
	}
	return ""
}

func (t *andlabsT) SetText(s string) bool {
	log(debugToolkit, "Text() SetText() Enter")
	if (t.uiEntry != nil) {
		log(debugToolkit, "Value() =", t.uiEntry.Text)
		t.uiEntry.SetText(s)
		return true
	}
	if (t.uiMultilineEntry != nil) {
		log(debugToolkit, "Value() =", t.uiMultilineEntry.Text)
		t.uiMultilineEntry.SetText(s)
		return true
	}
	return false
}

func sanity(t *andlabsT) bool {
	if (debugToolkit) {
		log(debugToolkit, "Value() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t.uiEntry == nil) {
		log(debugToolkit, "Value() =", t.uiEntry.Text)
		return false
	}
	return true
}

func (t *andlabsT) SetValue(i int) bool {
	log(debugToolkit, "SetValue() START")
	if (sanity(t)) {
		return false
	}
	t.Dump(debugToolkit)
	// panic("got to toolkit.SetValue")
	return true
}

func (t *andlabsT) Value() int {
	if (debugToolkit) {
		log(debugToolkit, "Value() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t == nil) {
		log(debugToolkit, "Value() can not get value t == nil")
		return 0
	}
	if (t.uiSlider != nil) {
		log(debugToolkit, "Value() =", t.uiSlider.Value)
		return t.uiSlider.Value()
	}
	if (t.uiSpinbox != nil) {
		log(debugToolkit, "Value() =", t.uiSpinbox.Value)
		return t.uiSpinbox.Value()
	}
	log(debugToolkit, "Value() Could not find a ui element to get a value from")
	return 0
}
