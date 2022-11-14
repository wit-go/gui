package main

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) MessageWindow(msg1 string, msg2 string) {
	ui.MsgBox(t.uiWindow, msg1, msg2)
}

func (t *andlabsT) ErrorWindow(msg1 string, msg2 string) {
	ui.MsgBoxError(t.uiWindow, msg1, msg2)
}

func NewWindow(w *toolkit.Widget) {
	var t *andlabsT

	if (DebugToolkit) {
		log.Println("toolkit NewWindow", w.Name, w.Width, w.Height)
	}

	if (w == nil) {
		log.Println("wit/gui plugin error. widget == nil")
		return
	}
	t = new(andlabsT)
	// t = NewWindow2(w.Name, w.Width, w.Height)

// func NewWindow2(title string, x int, y int) *andlabsT {
	// menubar bool is if the OS defined border on the window should be used
	win := ui.NewWindow(w.Name, w.Width, w.Height, menubar)
	win.SetBorderless(canvas)
	win.SetMargined(margin)
	win.OnClosing(func(*ui.Window) bool {
		if (DebugToolkit) {
			log.Println("ui.Window().OnExit() SHOULD ATTEMPT CALLBACK here")
			t.Dump()
		}
		if (w.Custom != nil) {
			w.Custom()
			return true
		}
		if (w.Event != nil) {
			w.Event(w)
			return true
		}
		if (DebugToolkit) {
			log.Println("andlabs.ui.Window().OnClosing() was not defined")
		}
		return false
	})
	win.Show()
	t.uiWindow = win
	t.UiWindowBad = win // deprecate this as soon as possible
	t.Name = w.Name

	mapWidgetsToolkits(w, t)
	return
}

func (t *andlabsT) SetWindowTitle(title string) {
	if (DebugToolkit) {
		log.Println("toolkit NewWindow", t.Name, "title", title)
	}
	win := t.uiWindow
	if (win != nil) {
		win.SetTitle(title)
	} else {
		if (DebugToolkit) {
			log.Println("Setting the window title", title)
		}
	}
}
