package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

var DebugToolkit bool = false

// stores the raw toolkit internals
type Toolkit struct {
	id     string

	Name   string
	Width  int
	Height int

	OnChanged func(*Toolkit)

	uiBox     *ui.Box
	uiButton  *ui.Button
	uiControl *ui.Control
	uiEntry   *ui.Entry
	uiLabel   *ui.Label
	uiSlider  *ui.Slider
	uiSpinbox *ui.Spinbox
	uiTab     *ui.Tab
	uiText    *ui.EditableCombobox
	uiWindow  *ui.Window
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
