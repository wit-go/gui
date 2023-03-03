package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// make new Group here
func (t *andlabsT) doAppend(newt *andlabsT, c *ui.Control) {

	if (newt.tw != nil) {
		if (newt.tw.Type == toolkit.Grid) {
			log(true, "doAppend() going to attempt uiGrid")
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
				log(true, "doAppend() on uiGrid to a uiBox")
				if (newt.Name == "output") {
					t.uiBox.Append(newt.uiGrid, true)
				} else {
				t.uiBox.Append(newt.uiGrid, stretchy)
				}
				return
			}
			log(true, "doAppend() on uiGrid failed")
			return
		}
	} else {
		log(true, "doAppend() newt.tw == nil ERROR on newt.Name =", newt.Name)
	}

	// hack to pass a group
	if (c == nil) {
		log(true, "attempting to doAppend() on a uiGroup")
		if (t.uiBox != nil) {
			if (newt.Name == "output") {
				t.uiBox.Append(newt.uiGroup, true)
			} else {
				t.uiBox.Append(newt.uiGroup, stretchy)
			}
			return
		}

		if (t.uiWindow != nil) {
			log(true, "This is a raw window without a box. probably make a box here and add the group to that")
			t.Dump(true)
			newt.Dump(true)
			t.uiBox = ui.NewHorizontalBox()
			t.uiWindow.SetChild(t.uiBox)
			log(true, "tried to make a box", t.uiBox)
			if (newt.Name == "output") {
				log(true, "tried to t.uiBox.Append(*c, true)")
				if (t.uiBox == nil) {
					log(true, "tried to t.uiBox.Append(*c, true)")
				}
				t.uiBox.Append(newt.uiGroup, true)
			} else {
				log(true, "tried to t.uiBox.Append(*c, stretchy)")
				t.uiBox.Append(newt.uiGroup, stretchy)
			}
			return
		}

		log(debugError, "NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugError, "probably could just make a box here?")
		exit("internal wit/gui error")
	}
	if (t.uiBox != nil) {
		// TODO: temporary hack to make the output textbox 'fullscreen'
		if (newt.Name == "output") {
			t.uiBox.Append(*c, true)
		} else {
			t.uiBox.Append(*c, stretchy)
		}
		return
	}
	if (t.uiWindow != nil) {
		log(true, "This is a raw window without a box. probably make a box here and add the group to that")
		t.uiBox = ui.NewHorizontalBox()
		t.uiWindow.SetChild(t.uiBox)
		log(true, "tried to make a box")
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
