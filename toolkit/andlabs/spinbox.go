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
	uiLabel   *ui.Label
	uiSlider  *ui.Slider
	uiSpinbox *ui.Spinbox
	uiTab     *ui.Tab
	uiText    *ui.EditableCombobox
	uiWindow  *ui.Window
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

func NewSpinbox(b *ui.Box, name string, x int, y int) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewSpinbox()", x, y)
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewSpinbox() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return &t
	}
	spin := ui.NewSpinbox(x, y)
	t.uiSpinbox = spin
	t.uiBox = b
	t.uiBox.Append(spin, false)

	spin.OnChanged(func(spin *ui.Spinbox) {
		i := spin.Value()
		if (DebugToolkit) {
			log.Println("gui.Toolbox.ui.OnChanged() val =", i)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(t)
		}
		if (t.OnChanged != nil) {
			log.Println("gui.Toolbox.OnChanged() entered val =", i)
			t.OnChanged(&t)
		}
	})

	return &t
}
