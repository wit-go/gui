package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

var DebugToolkit bool = false

var streachy = true
var border = true


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
	uiEntry   *ui.Entry
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

func (t *Toolkit) GetText() string {
	if (DebugToolkit) {
		log.Println("gui.Toolkit.Text() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t.uiEntry != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.Value() =", t.uiEntry.Text)
		}
		return t.uiEntry.Text()
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
	log.Println("gui.Toolkit.Dump()", t.Name, t.Width, t.Height)
	if (t.uiBox != nil) {
		log.Println("gui.Toolkit.Dump() uiBox   =", t.uiBox)
	}
	if (t.uiButton != nil) {
		log.Println("gui.Toolkit.Dump() uiButton =", t.uiButton)
	}
	if (t.uiCombobox != nil) {
		log.Println("gui.Toolkit.Dump() uiCombobox =", t.uiCombobox)
	}
	if (t.uiWindow != nil) {
		log.Println("gui.Toolkit.Dump() uiWindow =", t.uiWindow)
	}
	if (t.uiTab != nil) {
		log.Println("gui.Toolkit.Dump() uiTab =", t.uiTab)
	}
	if (t.uiGroup != nil) {
		log.Println("gui.Toolkit.Dump() uiGroup =", t.uiGroup)
	}
	if (t.uiSlider != nil) {
		log.Println("gui.Toolkit.Dump() uiSlider =", t.uiSlider)
	}
	if (t.OnExit != nil) {
		log.Println("gui.Toolkit.Dump() OnExit =", t.OnExit)
	}
	if (t.Custom != nil) {
		log.Println("gui.Toolkit.Dump() Custom =", t.Custom)
	}
	log.Println("gui.Toolkit.Dump() c =", t.c)
	log.Println("gui.Toolkit.Dump() val =", t.val)
}
