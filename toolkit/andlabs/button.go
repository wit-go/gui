package toolkit

import "log"
// import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// make new Group here
func (t Toolkit) NewButton(name string) *Toolkit {
	var newt Toolkit
	var b *ui.Button

	if t.broken() {
		return nil
	}

	if (DebugToolkit) {
		log.Println("gui.Toolbox.NewButton() create", name)
	}
	b = ui.NewButton(name)
	newt.uiButton = b

	b.OnClicked(func(*ui.Button) {
		log.Println("TODO: IN TOOLKIT GOROUTINE. SHOULD LEAVE HERE VIA channels. button name =", name)
		t.Dump()
		newt.Dump()
		if (DebugToolkit) {
			log.Println("wit/gui/toolkit NewButton() Should do something here")
		}
		if (newt.Custom == nil) {
			if (DebugToolkit) {
				log.Println("wit/gui/toolkit NewButton() toolkit.Custom == nil")
			}
		} else {
			if (DebugToolkit) {
				log.Println("wit/gui/toolkit NewButton() toolkit.Custom() START")
			}
			newt.Custom()
			return
			if (DebugToolkit) {
				log.Println("wit/gui/toolkit NewButton() toolkit.Custom() END")
			}
		}
		if (t.Custom == nil) {
			if (DebugToolkit) {
				log.Println("wit/gui/toolkit NewButton() parent toolkit.Custom == nil")
			}
		} else {
			if (DebugToolkit) {
				log.Println("wit/gui/toolkit NewButton() running parent toolkit.Custom() START (IS THIS A BAD IDEA?)")
			}
			t.Custom()
			return
			if (DebugToolkit) {
				log.Println("wit/gui/toolkit NewButton() running parent toolkit.Custom() END   (IS THIS A BAD IDEA?)")
			}
		}
		log.Println("TODO: LEFT TOOLKIT GOROUTINE WITH NOTHING TO DO button name =", name)
	})

	if (DebugToolkit) {
		log.Println("gui.Toolbox.NewButton() about to append to Box parent t:", name)
		t.Dump()
		log.Println("gui.Toolbox.NewButton() about to append to Box new t:", name)
		newt.Dump()
	}
	if (t.uiBox != nil) {
		t.uiBox.Append(b, stretchy)
	} else if (t.uiWindow != nil) {
		t.uiWindow.SetChild(b)
	} else {
		log.Println("ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		log.Println("ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		return &t
	}

	return &newt
}
