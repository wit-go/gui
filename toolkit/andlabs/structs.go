package main

import "git.wit.org/wit/gui/toolkit"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// var andlabs map[int]*andlabsT
// var callback func(int) bool
var callback chan toolkit.Action

type node struct {
	parent	*node
	children []*node

	WidgetId	int	// widget ID
	WidgetType	toolkit.WidgetType
	ParentId	int	// parent ID

	Name   string
	Text   string

	// This is how the values are passed back and forth
	// values from things like checkboxes & dropdown's
	B	bool
	I	int
	S	string

	A	any // switch to this or deprecate this? pros/cons?

	// This is used for things like a slider(0,100)
	X      int
	Y      int

	// the internal plugin toolkit structure
	tk *andlabsT
}

// stores the raw toolkit internals
type andlabsT struct {
	// wId	int	// widget ID
	// WidgetType	toolkit.WidgetType

	Name   string
	Text   string
	// Type   toolkit.WidgetType
	Width  int
	Height int

	// tw	*toolkit.Widget
	parent	*andlabsT
	children []*andlabsT

	uiControl ui.Control

	uiBox     *ui.Box
	uiButton  *ui.Button
	uiCombobox *ui.Combobox
	uiCheckbox *ui.Checkbox
	uiEntry   *ui.Entry
	uiGroup   *ui.Group
	uiLabel   *ui.Label
	uiSlider  *ui.Slider
	uiSpinbox *ui.Spinbox
	uiTab     *ui.Tab
	uiWindow  *ui.Window
	uiMultilineEntry   *ui.MultilineEntry
	uiEditableCombobox    *ui.EditableCombobox
	uiGrid    *ui.Grid
	uiImage  *ui.Image
	gridX	int
	gridY	int

	// used as a counter to work around limitations of widgets like combobox
	// this is probably fucked up and in many ways wrong because of unsafe goroutine threading
	// but it's working for now due to the need for need for a correct interaction layer betten toolkits
	c int
	i int
	b bool
	s string
	val map[int]string

	// andlabs/ui only accesses widget id numbers
	boxC int	// how many things on in a box
	boxW map[int]int // find a widget in a box
	text   string
}
