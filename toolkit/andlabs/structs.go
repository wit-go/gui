package main

import "git.wit.org/wit/gui/toolkit"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// stores the raw toolkit internals
type andlabsT struct {
	id     string

	Name   string
	Width  int
	Height int
	tw	*toolkit.Widget

	OnChanged func(*andlabsT)
	OnExit    func(*andlabsT)

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
