package toolkit

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *Toolkit) MessageWindow(msg1 string, msg2 string) {
	ui.MsgBox(t.uiWindow, msg1, msg2)
}

func (t *Toolkit) ErrorWindow(msg1 string, msg2 string) {
	ui.MsgBoxError(t.uiWindow, msg1, msg2)
}

func NewWindow(title string, x int, y int) *Toolkit {
	var t Toolkit
	log.Println("toolkit NewWindow", title, x, y)
	w := ui.NewWindow(title, x, y, menubar)
	w.SetBorderless(canvas)
	w.SetMargined(margin)
	w.OnClosing(func(*ui.Window) bool {
		log.Println("ui.Window().OnExit() SHOULD ATTEMPT CALLBACK here")
		t.Dump()
		if (t.OnExit != nil) {
			log.Println("ui.Window().OnExit() ATTEMPTING toolkit.OnExit CALLBACK")
			t.OnExit(&t)
		}
		if (t.Custom != nil) {
			log.Println("ui.Window().Custom() ATTEMPTING toolkit.Custom CALLBACK")
			t.Custom()
		}
		log.Println("ui.Window().OnExit() Toolkit.OnExit is nil")
		return true
	})
	w.Show()
	t.uiWindow = w
	t.UiWindowBad = w // deprecate this as soon as possible
	return &t
}

func (t *Toolkit) SetWindowTitle(title string) {
	log.Println("toolkit NewWindow", t.Name, "title", title)
	win := t.uiWindow
	if (win != nil) {
		win.SetTitle(title)
	} else {
		log.Println("Setting the window title", title)
	}
}
