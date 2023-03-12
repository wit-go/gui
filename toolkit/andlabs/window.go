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

func newWindow(w *toolkit.Widget) {
	var newt *andlabsT

	log(debugToolkit, "toolkit NewWindow", w.Name, w.Width, w.Height)

	if (w == nil) {
		log(debugToolkit, "wit/gui plugin error. widget == nil")
		return
	}
	newt = new(andlabsT)
	newt.tw = w

	// menubar bool is if the OS defined border on the window should be used
	win := ui.NewWindow(w.Name, w.Width, w.Height, menubar)
	win.SetBorderless(canvas)
	win.SetMargined(margin)
	win.OnClosing(func(*ui.Window) bool {
		newt.commonChange(newt.tw)
		return true
	})
	win.Show()
	newt.uiWindow = win
	// newt.UiWindowBad = win // deprecate this as soon as possible
	newt.Name = w.Name

	mapWidgetsToolkits(w, newt)
	return
}

func (t *andlabsT) SetWindowTitle(title string) {
	log(debugToolkit, "toolkit NewWindow", t.Name, "title", title)
	win := t.uiWindow
	if (win != nil) {
		win.SetTitle(title)
	} else {
		log(debugToolkit, "Setting the window title", title)
	}
}

func doWindow(c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newWindow(c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if (ct.uiWindow == nil) {
		log(debugError, "Window() uiWindow == nil", ct)
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	switch c.Action {
	case "Show":
		ct.uiWindow.Show()
	case "Hide":
		ct.uiWindow.Hide()
	case "Enable":
		ct.uiWindow.Enable()
	case "Disable":
		ct.uiWindow.Disable()
	case "Get":
		c.S = ct.uiWindow.Title()
	case "Set":
		ct.uiWindow.SetTitle(c.S)
	case "SetText":
		ct.uiWindow.SetTitle(c.S)
	case "SetMargin":
		ct.uiWindow.SetMargined(c.B)
	case "SetBorder":
		ct.uiWindow.SetBorderless(c.B)
	case "Delete":
		ct.uiWindow.Destroy()
	default:
		log(debugError, "Can't do", c.Action, "to a Window")
	}
}
