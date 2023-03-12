package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// move all the append code here
func (t *andlabsT) doAppend(kind toolkit.WidgetType, newt *andlabsT, c *ui.Control) {
	if (kind == toolkit.Grid) {
		log(debugToolkit, "doAppend() attempt to append a ui.Control into a uiGrid")
		// hack to add shit to a grid
		button1 := ui.NewButton("a(0,2)")
		newt.uiGrid.Append(button1,
			0, 2, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
		button2 := ui.NewButton("a(1,2)")
		newt.uiGrid.Append(button2,
			1, 2, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
	
		if (t.uiBox != nil) {
			log(debugToolkit, "doAppend() on uiGrid to a uiBox")
			if (newt.Name == "output") {
				t.uiBox.Append(newt.uiGrid, true)
			} else {
				t.uiBox.Append(newt.uiGrid, stretchy)
			}
			return
		}
		log(debugToolkit, "doAppend() on uiGrid failed")
		return
	}

	if (kind == toolkit.Group) {
		log(debugToolkit, "doAppend() attempt a uiGroup")
		if (t.uiBox != nil) {
			if (newt.Name == "output") {
				t.uiBox.Append(newt.uiGroup, true)
			} else {
				t.uiBox.Append(newt.uiGroup, stretchy)
			}
			return
		}

		if (t.uiWindow != nil) {
			log(debugToolkit, "This is a raw window without a box. probably make a box here and add the group to that")
			newt.Dump(debugToolkit)
			t.uiBox = ui.NewHorizontalBox()
			t.uiWindow.SetChild(t.uiBox)
			log(debugToolkit, "tried to make a box", t.uiBox)
			if (newt.Name == "output") {
				log(debugToolkit, "tried to t.uiBox.Append(*c, true)")
				if (t.uiBox == nil) {
					log(debugToolkit, "tried to t.uiBox.Append(*c, true)")
				}
				t.uiBox.Append(newt.uiGroup, true)
			} else {
				log(debugToolkit, "tried to t.uiBox.Append(*c, stretchy)")
				t.uiBox.Append(newt.uiGroup, stretchy)
			}
			return
		}

		log(debugError, "NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugError, "probably could just make a box here?")
		exit("internal wit/gui error")
	}

	if (kind == toolkit.Textbox) {
		if (t.uiBox == nil) {
			log(debugError, "NewTextbox() node.UiBox == nil. I can't add a range UI element without a place to put it")
			log(debugError, "probably could just make a box here?")
			exit("internal wit/gui error")
		}
		// TODO: temporary hack to make the output textbox 'fullscreen'
		if (newt.Name == "output") {
			t.uiBox.Append(*c, true)
		} else {
			t.uiBox.Append(*c, stretchy)
		}
		return
	}

	if (t.uiWindow != nil) {
		log(debugToolkit, "This is a raw window without a box. probably make a box here and add the group to that")
		t.uiBox = ui.NewHorizontalBox()
		t.uiWindow.SetChild(t.uiBox)
		log(debugToolkit, "tried to make a box")
		if (newt.Name == "output") {
			t.uiBox.Append(*c, true)
		} else {
			t.uiBox.Append(*c, stretchy)
		}
		return
	}

	log(debugError, "NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
	log(debugError, "probably could just make a box here?")
	exit("internal wit/gui error")
}
