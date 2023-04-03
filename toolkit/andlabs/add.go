package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func actionDump(b bool, a *toolkit.Action) {
	log(b, "actionDump() Widget.Type =", a.ActionType)
	log(b, "actionDump() Widget.S =", a.S)
	log(b, "actionDump() Widget.I =", a.I)
	log(b, "actionDump() WidgetId =", a.WidgetId)
	log(b, "actionDump() ParentId =", a.ParentId)
}

func add(a *toolkit.Action) {
	if (andlabs[a.WidgetId] != nil) {
		log(debugError, "add() error. can't make a widget that already exists. id =", a.WidgetId)
		actionDump(debugError, a)
		return
	}
	if (a.WidgetId == 0) {
		log(debugError, "add() error. w.WidgetId == 0")
		actionDump(debugError, a)
		return
	}

	// for now, window gets handled without checking where == nil)
	if (a.WidgetType == toolkit.Window) {
		doWindow(a)
		return
	}

	if (andlabs[a.ParentId] == nil) {
		// listMap(debugError) // memory corruption?
		log(debugError, "add() Widget.Name =", a.Name)
		log(debugError, "add() Widget.Type =", a.WidgetType)
		log(debugError, "ERROR add() ERROR a.Parent map to t == nil. WidgetId =", a.WidgetId, "ParentId =", a.ParentId)
		exit("can not add()")
		return
	}

	switch a.WidgetType {
	case toolkit.Window:
		doWindow(a)
		return
	case toolkit.Tab:
		newTab(a)
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
		log(debugError, "add() error TODO: ", a.WidgetType, a.Name)
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
	log(debugAction, "place() START", a.WidgetType, a.Name)

	// add the structure to the array
	if (andlabs[a.WidgetId] == nil) {
		log(logInfo, "newTab() MAPPED", a.WidgetId, a.ParentId)
		andlabs[a.WidgetId] = newt
		newt.WidgetType = a.WidgetType
	} else {
		log(debugError, "newTab() DO WHAT?", a.WidgetId, a.ParentId)
		log(debugError, "THIS IS BAD")
	}

	if (newt.uiControl == nil) {
		log(debugError, "place() ERROR uiControl == nil", a.ParentId)
		return false
	}

	where := andlabs[a.ParentId]
	if (where == nil) {
		log(debugError, "place() ERROR where == nil", a.ParentId)
		return false
	}

	switch where.WidgetType {
	case toolkit.Grid:
		log(debugGrid, "add() Grid try at Parent X,Y =", a.X, a.Y)
		newt.gridX = a.X
		newt.gridY = a.Y
		log(debugGrid, "add() Grid try at gridX,gridY", newt.gridX, newt.gridY)
		// at the very end, subtract 1 from X & Y since andlabs/ui starts counting at zero
		t.uiGrid.Append(newt.uiControl,
			newt.gridY - 1, newt.gridX - 1, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
		return true
	case toolkit.Group:
		if (t.uiBox == nil) {
			t.uiGroup.SetChild(newt.uiControl)
			log(debugGrid, "add() hack Group to use this as the box?", a.Name, a.WidgetType)
			t.uiBox  = newt.uiBox
		} else {
			t.uiBox.Append(newt.uiControl, stretchy)
		}
		return true
	case toolkit.Tab:
		t.uiTab.Append(a.Text, newt.uiControl)
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
		log(debugError, "add() how?", a.ParentId)
	}
	return false
}
