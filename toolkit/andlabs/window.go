package main

import (
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

func newWindow(a *toolkit.Action) {
	w := a.Widget
	var newt *andlabsT

	// log(debugToolkit, "toolkit NewWindow", w.Name, w.Width, w.Height)

	if (w == nil) {
		log(debugToolkit, "wit/gui plugin error. widget == nil")
		return
	}
	newt = new(andlabsT)
	newt.tw = w
	newt.Type = toolkit.Window
	newt.wId = a.WidgetId

	// menubar bool is if the OS defined border on the window should be used
	win := ui.NewWindow(w.Name, a.Width, a.Height, menubar)
	win.SetBorderless(canvas)
	win.SetMargined(margin)
	win.OnClosing(func(*ui.Window) bool {
		newt.commonChange(newt.tw, a.WidgetId)
		return true
	})
	win.Show()
	newt.uiWindow = win
	newt.uiControl = win
	// newt.UiWindowBad = win // deprecate this as soon as possible
	newt.Name = w.Name

	andlabs[a.WidgetId] = newt
	return
}

func (t *andlabsT) SetWindowTitle(title string) {
	log(debugToolkit, "toolkit NewWindow", t.Name, "title", title)
	win := t.uiWindow
	if (win == nil) {
		log(debugError, "Error: no window", t.wId)
	} else {
		win.SetTitle(title)
		log(debugToolkit, "Setting the window title", title)
	}
}

func doWindow(a *toolkit.Action) {
	newWindow(a)
}