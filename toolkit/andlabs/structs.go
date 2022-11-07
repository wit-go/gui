package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

var defaultBehavior bool = true

var bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
var canvas bool // if set to true, the windows are a raw canvas
var menubar bool // for windows
var stretchy bool // expand things like buttons to the maximum size
var padded bool // add space between things like buttons
var margin bool // add space around the frames of windows

var DebugToolkit bool

func setDefaultBehavior(s bool) {
	defaultBehavior = s
	if (defaultBehavior) {
		if (DebugToolkit) {
			log.Println("Setting this toolkit to use the default behavior.")
			log.Println("This is the 'guessing' part as defined by the wit/gui 'Principles'. Refer to the docs.")
		}
		stretchy = false
		padded = true
		menubar = true
		margin = true
		canvas = false
		bookshelf = true // 99% of the time, things make a vertical stack of objects

		DebugToolkit = false
	} else {
		log.Println("This toolkit is set to ignore the default behavior.")
	}
}

func SetDebugToolkit (s bool) {
	DebugToolkit = s
}

func GetDebugToolkit () bool {
	return DebugToolkit
}

// stores the raw toolkit internals
type Toolkit struct {
	id     string

	Name   string
	Width  int
	Height int

	OnChanged func(*Toolkit)
	OnExit    func(*Toolkit)

	Custom  func()

	uiBox     *ui.Box
	uiBox2    *ui.Box	// temporary hack while implementing tabs
	uiButton  *ui.Button
	uiControl *ui.Control
	uiCombobox *ui.Combobox
	uiCheckbox *ui.Checkbox
	uiEntry   *ui.Entry
	uiMultilineEntry   *ui.MultilineEntry
	uiGroup   *ui.Group
	uiLabel   *ui.Label
	uiSlider  *ui.Slider
	uiSpinbox *ui.Spinbox
	uiTab     *ui.Tab
	uiText    *ui.EditableCombobox
	uiWindow  *ui.Window
	UiWindowBad  *ui.Window

	// used as a counter to work around limitations of widgets like combobox
	// this is probably fucked up and in many ways wrong because of unsafe goroutine threading
	// but it's working for now due to the need for need for a correct interaction layer betten toolkits
	c int
	val map[int]string
	text   string
}

func (t *Toolkit) String() string {
	return t.GetText()
}

func forceDump(t *Toolkit) {
	tmp := DebugToolkit
	DebugToolkit = true
	t.Dump()
	DebugToolkit = tmp
}

func (t *Toolkit) GetText() string {
	t.Dump()
	if (DebugToolkit) {
		log.Println("gui.Toolkit.Text() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t.uiEntry != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiEntry.Text())
		}
		return t.uiEntry.Text()
	}
	if (t.uiMultilineEntry != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiMultilineEntry.Text())
		}
		text := t.uiMultilineEntry.Text()
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() text =", text)
		}
		t.text = text
		return text
	}
	if (t.uiCombobox != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.GetText() =", t.text)
		}
		return t.text
	}
	return ""
}

func (t *Toolkit) SetText(s string) bool {
	if (DebugToolkit) {
		log.Println("gui.Toolkit.Text() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t.uiEntry != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiEntry.Text)
		}
		t.uiEntry.SetText(s)
		return true
	}
	if (t.uiMultilineEntry != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiMultilineEntry.Text)
		}
		t.uiMultilineEntry.SetText(s)
		return true
	}
	return false
}

func sanity(t *Toolkit) bool {
	if (DebugToolkit) {
		log.Println("gui.Toolkit.Value() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t.uiEntry == nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiEntry.Text)
		}
		return false
	}
	return true
}

func (t *Toolkit) SetValue(i int) bool {
	log.Println("gui.Toolkit.SetValue() START")
	if (sanity(t)) {
		return false
	}
	t.Dump()
	// panic("got to toolkit.SetValue")
	return true
}

func (t *Toolkit) Value() int {
	if (DebugToolkit) {
		log.Println("gui.Toolkit.Value() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t == nil) {
		log.Println("gui.Toolkit.Value() can not get value t == nil")
		return 0
	}
	if (t.uiSlider != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiSlider.Value)
		}
		return t.uiSlider.Value()
	}
	if (t.uiSpinbox != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiSpinbox.Value)
		}
		return t.uiSpinbox.Value()
	}
	log.Println("gui.Toolkit.Value() Could not find a ui element to get a value from")
	return 0
}

func (t *Toolkit) Dump() {
	if ! DebugToolkit {
		return
	}
	log.Println("gui.Toolkit.Dump() Name  = ", t.Name, t.Width, t.Height)
	if (t.uiBox != nil) {
		log.Println("gui.Toolkit.Dump() uiBox      =", t.uiBox)
	}
	if (t.uiButton != nil) {
		log.Println("gui.Toolkit.Dump() uiButton    =", t.uiButton)
	}
	if (t.uiCombobox != nil) {
		log.Println("gui.Toolkit.Dump() uiCombobox  =", t.uiCombobox)
	}
	if (t.uiWindow != nil) {
		log.Println("gui.Toolkit.Dump() uiWindow    =", t.uiWindow)
	}
	if (t.uiTab != nil) {
		log.Println("gui.Toolkit.Dump() uiTab       =", t.uiTab)
	}
	if (t.uiGroup != nil) {
		log.Println("gui.Toolkit.Dump() uiGroup     =", t.uiGroup)
	}
	if (t.uiEntry != nil) {
		log.Println("gui.Toolkit.Dump() uiEntry     =", t.uiEntry)
	}
	if (t.uiMultilineEntry != nil) {
		log.Println("gui.Toolkit.Dump() uiMultilineEntry =", t.uiMultilineEntry)
	}
	if (t.uiSlider != nil) {
		log.Println("gui.Toolkit.Dump() uiSlider    =", t.uiSlider)
	}
	if (t.uiCheckbox != nil) {
		log.Println("gui.Toolkit.Dump() uiCheckbox  =", t.uiCheckbox)
	}
	if (t.OnExit != nil) {
		log.Println("gui.Toolkit.Dump() OnExit      =", t.OnExit)
	}
	if (t.Custom != nil) {
		log.Println("gui.Toolkit.Dump() Custom      =", t.Custom)
	}
	log.Println("gui.Toolkit.Dump() c         =", t.c)
	log.Println("gui.Toolkit.Dump() val       =", t.val)
	log.Println("gui.Toolkit.Dump() text      =", t.text)
}
