package toolkit

import (
	"log"
	"os"

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
	w := ui.NewWindow(title, x, y, false)
	w.SetBorderless(false)
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
		t.Dump()
		os.Exit(0)
		return true
	})
	w.SetMargined(true)
	w.Show()
	t.uiWindow = w
	t.UiWindowBad = w // deprecate this as soon as possible
	return &t
}
