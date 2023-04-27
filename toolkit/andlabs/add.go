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

func add(a toolkit.Action) {
	if (andlabs[a.WidgetId] != nil) {
		log(debugError, "add() error. can't make a widget that already exists. id =", a.WidgetId)
		actionDump(debugError, &a)
		return
	}
	if (a.WidgetType == toolkit.Root) {
		rootNode = addWidget(&a, nil)
		return
	}
	n := addWidget(&a, nil)

	p := n.parent
	switch n.WidgetType {
	case toolkit.Window:
		newWindow(n)
		return
	case toolkit.Tab:
		p.newTab(n)
		return
	case toolkit.Label:
		newLabel(&a)
		return
	case toolkit.Button:
		p.newButton(n)
		return
	case toolkit.Grid:
		newGrid(&a)
		return
	case toolkit.Checkbox:
		newCheckbox(&a)
		return
	case toolkit.Spinner:
		newSpinner(&a)
		return
	case toolkit.Slider:
		newSlider(&a)
		return
	case toolkit.Dropdown:
		newDropdown(&a)
		return
	case toolkit.Combobox:
		newCombobox(&a)
		return
	case toolkit.Textbox:
		newTextbox(&a)
		return
	case toolkit.Group:
		p.newGroup(n)
		return
	case toolkit.Box:
		p.newBox(n)
		return
	case toolkit.Image:
		newImage(&a)
		return
	default:
		log(debugError, "add() error TODO: ", n.WidgetType, n.Name)
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
		log(logInfo, "place() MAPPED", a.WidgetId, a.ParentId)
		andlabs[a.WidgetId] = newt
		newt.WidgetType = a.WidgetType
	} else {
		log(debugError, "place() DO WHAT?", a.WidgetId, a.ParentId)
		log(debugError, "place() THIS IS BAD")
	}
	log(logInfo, "place() DONE MAPPED", a.WidgetId, a.ParentId)

	if (newt.uiControl == nil) {
		log(debugError, "place() ERROR uiControl == nil", a.ParentId)
		return false
	}

	where := andlabs[a.ParentId]
	if (where == nil) {
		log(debugError, "place() ERROR where == nil", a.ParentId)
		return false
	}

	log(logInfo, "place() switch", where.WidgetType)
	switch where.WidgetType {
	case toolkit.Grid:
		log(debugGrid, "place() Grid try at Parent X,Y =", a.X, a.Y)
		newt.gridX = a.X
		newt.gridY = a.Y
		log(debugGrid, "place() Grid try at gridX,gridY", newt.gridX, newt.gridY)
		// at the very end, subtract 1 from X & Y since andlabs/ui starts counting at zero
		t.uiGrid.Append(newt.uiControl,
			newt.gridY - 1, newt.gridX - 1, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
		return true
	case toolkit.Group:
		if (t.uiBox == nil) {
			t.uiGroup.SetChild(newt.uiControl)
			log(debugGrid, "place() hack Group to use this as the box?", a.Name, a.WidgetType)
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
		log(logInfo, "place() uiBox =", t.uiBox)
		log(logInfo, "place() uiControl =", newt.uiControl)
		t.uiBox.Append(newt.uiControl, stretchy)
		t.boxC += 1
		return true
	case toolkit.Window:
		t.uiWindow.SetChild(newt.uiControl)
		return true
	default:
		log(debugError, "place() how?", a.ParentId)
	}
	return false
}
func (p *node) place(n *node) bool {
	log(logInfo, "place() START", n.WidgetType, n.Name)

	if (p.tk == nil) {
		log(logError, "p.tk == nil", p.Name, p.ParentId, p.WidgetType, p.tk)
		log(logError, "n = ", n.Name, n.ParentId, n.WidgetType, n.tk)
		panic("p.tk == nil")
	}

	log(logInfo, "place() switch", p.WidgetType)
	switch p.WidgetType {
	case toolkit.Grid:
		log(debugGrid, "place() Grid try at Parent X,Y =", n.X, n.Y)
		n.tk.gridX = n.X
		n.tk.gridY = n.Y
		log(debugGrid, "place() Grid try at gridX,gridY", n.tk.gridX, n.tk.gridY)
		// at the very end, subtract 1 from X & Y since andlabs/ui starts counting at zero
		p.tk.uiGrid.Append(n.tk.uiControl,
			n.tk.gridY - 1, n.tk.gridX - 1, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
		return true
	case toolkit.Group:
		if (p.tk.uiBox == nil) {
			p.tk.uiGroup.SetChild(n.tk.uiControl)
			log(debugGrid, "place() hack Group to use this as the box?", n.Name, n.WidgetType)
			p.tk.uiBox  = n.tk.uiBox
		} else {
			p.tk.uiBox.Append(n.tk.uiControl, stretchy)
		}
		return true
	case toolkit.Tab:
		if (p.tk.uiTab == nil) {
			log(logError, "p.tk.uiTab == nil", p.tk)
			panic("p.tk.uiTab == nil")
		}
		if (n.tk.uiControl == nil) {
			log(logError, "n.tk.uiControl == nil", n.tk)
			panic("n.tk.uiControl == nil")
		}
		p.tk.uiTab.Append(n.Text, n.tk.uiControl)
		p.tk.boxC += 1
		return true
	case toolkit.Box:
		log(logInfo, "place() uiBox =", p.tk.uiBox)
		log(logInfo, "place() uiControl =", n.tk.uiControl)
		p.tk.uiBox.Append(n.tk.uiControl, stretchy)
		p.tk.boxC += 1
		return true
	case toolkit.Window:
		p.tk.uiWindow.SetChild(n.tk.uiControl)
		return true
	default:
		log(debugError, "place() how? Parent =", p.WidgetId, p.WidgetType)
	}
	return false
}
