package toolkit

import "log"
import "os"
// import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func (pt *Toolkit) NewDropdown(title string) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewDropdownCombobox()")
	var newt Toolkit

	if (pt.uiBox == nil) {
		log.Println("gui.ToolboxNode.NewDropdown() node.UiBox == nil. I can't add a range UI element without a place to put it")
		os.Exit(0)
		return nil
	}

	s := ui.NewCombobox()
	newt.uiCombobox = s
	newt.uiBox = pt.uiBox
	pt.uiBox.Append(s, false)

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	s.OnSelected(func(spin *ui.Combobox) {
		i := spin.Selected()
		if (newt.val == nil) {
			log.Println("make map didn't work")
			os.Exit(0)
		}
		newt.text = newt.val[i]
		val := newt.text
		log.Println("gui.Toolbox.ui.Dropdown.OnChanged() val =", i, val)
		if (DebugToolkit) {
			log.Println("gui.Toolbox.ui.OnChanged() val =", i, val)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(newt)
		}
		if (newt.OnChanged != nil) {
			log.Println("gui.Toolbox.OnChanged() trying to run toolkit.OnChanged() entered val =", i, val)
			newt.OnChanged(&newt)
			return
		}
		if (newt.Custom != nil) {
			log.Println("gui.Toolbox.OnChanged() Running toolkit.Custom()", i, val)
			newt.Custom()
			return
		}
		log.Println("gui.Toolbox.Dropdown.OnChanged() ENDED without finding any callback", i, val)
	})

	return &newt
}

func (t *Toolkit) AddDropdown(title string) {
	t.uiCombobox.Append(title)
	if (t.val == nil) {
		log.Println("make map didn't work")
		return
	}
	t.val[t.c] = title
	t.c = t.c + 1
}

func (t Toolkit) SetDropdown(i int) {
	t.uiCombobox.SetSelected(i)
}
