package toolkit

import "log"
import "os"
// import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t *Toolkit) NewDropdown(title string) *Toolkit {
	// make new node here
	if (DebugToolkit) {
		log.Println("gui.Toolbox.NewDropdownCombobox()", title)
	}
	var newt Toolkit

	if t.broken() {
		return nil
	}

	s := ui.NewCombobox()
	newt.uiCombobox = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

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
		newt.commonChange("Dropdown")
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
