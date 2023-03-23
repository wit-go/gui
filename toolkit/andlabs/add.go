package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func actionDump(b bool, a *toolkit.Action) {
	log(b, "dump() Widget.Type =", a.Type)
	log(b, "dump() Widget.S =", a.S)
	log(b, "dump() Widget.I =", a.I)
	log(b, "dump() Widget =", a.Widget)
	log(b, "dump() Where =", a.Where)
}

func add(a *toolkit.Action) {
	if (a.Widget == nil) {
		log(debugError, "add() error. w.Widget == nil")
		actionDump(debugError, a)
		return
	}
	// for now, window gets handled without checking where == nil)
	if (a.Widget.Type == toolkit.Window) {
		doWindow(a)
		return
	}

	t := mapToolkits[a.Where]
	if (t == nil) {
		// listMap(debugError) // memory corruption?
		log(debugError, "add() Widget.Name =", a.Widget.Name, a.Widget.Type)
		// log(debugError, "add() Where.Name =", a.Where.Name)
		log(debugError, "ERROR add() ERROR a.Where map to t == nil.")
		return
	}

	switch a.Widget.Type {
	case toolkit.Window:
		doWindow(a)
		return
	case toolkit.Tab:
		doTab(a)
		return
	case toolkit.Label:
		newLabel(a)
		return
	case toolkit.Button:
		newButton(a)
		return
	case toolkit.Grid:
		newGrid(a)
		return
	case toolkit.Checkbox:
		newCheckbox(a)
		return
	case toolkit.Spinner:
		newSpinner(a)
		return
	case toolkit.Slider:
		newSlider(a)
		return
	case toolkit.Dropdown:
		newDropdown(a)
		return
	case toolkit.Combobox:
		newCombobox(a)
		return
	case toolkit.Textbox:
		newTextbox(a)
		return
	case toolkit.Group:
		newGroup(a)
		return
	case toolkit.Box:
		newBox(a)
		return
	case toolkit.Image:
		newImage(a)
		return
	default:
		log(debugError, "add() error TODO: ", a.Widget.Type, a.Widget.Name)
	}
}

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

// Grid numbering examples by (X,Y)
// ---------
// -- (1) --
// -- (2) --
// ---------
//
// -----------------------------
// -- (1,1) -- (1,2) -- (1,3) --
// -- (2,1) -- (2,2) -- (2,3) --
// -----------------------------

// internally for andlabs/ui
// (x&y flipped and start at zero) 
// -----------------------------
// -- (0,0) -- (1,0) -- (1,0) --
// -- (0,1) -- (1,1) -- (1,1) --
// -----------------------------
func place(a *toolkit.Action, t *andlabsT, newt *andlabsT) bool {
	log(debugAction, "place() START", a.Widget.Type, a.Widget.Name)

	if (newt.uiControl == nil) {
		log(debugError, "place() ERROR uiControl == nil", a.Where.Type, a.Where.Name)
		return false
	}

	switch a.Where.Type {
	case toolkit.Grid:
		log(debugGrid, "add() Grid try at Where X,Y =", a.Where.X, a.Where.Y)
		newt.gridX = a.Where.X
		newt.gridY = a.Where.Y
		log(debugGrid, "add() Grid try at gridX,gridY", newt.gridX, newt.gridY)
		// at the very end, subtract 1 from X & Y since andlabs/ui starts counting at zero
		t.uiGrid.Append(newt.uiControl,
			newt.gridY - 1, newt.gridX - 1, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
		return true
	case toolkit.Group:
		if (t.uiBox == nil) {
			t.uiGroup.SetChild(newt.uiControl)
			log(debugGrid, "add() hack Group to use this as the box?", a.Widget.Name, a.Widget.Type)
			t.uiBox  = newt.uiBox
		} else {
			t.uiBox.Append(newt.uiControl, stretchy)
		}
		return true
	case toolkit.Tab:
		t.uiBox.Append(newt.uiControl, stretchy)
		t.boxC += 1
		return true
	case toolkit.Box:
		t.uiBox.Append(newt.uiControl, stretchy)
		t.boxC += 1
		return true
	case toolkit.Window:
		t.uiWindow.SetChild(newt.uiControl)
		return true
	default:
		log(debugError, "add() how?", a.Where.Type)
	}
	return false
}
