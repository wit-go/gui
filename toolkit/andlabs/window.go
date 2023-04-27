package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) MessageWindow(msg1 string, msg2 string) {
	ui.MsgBox(t.uiWindow, msg1, msg2)
}

func (t *andlabsT) ErrorWindow(msg1 string, msg2 string) {
	ui.MsgBoxError(t.uiWindow, msg1, msg2)
}

func newWindow(n *node) {
	var newt *andlabsT

	newt = new(andlabsT)

	// menubar bool is if the OS defined border on the window should be used
	win := ui.NewWindow(n.Name, n.X, n.Y, menubar)
	win.SetBorderless(canvas)
	win.SetMargined(margin)
	win.OnClosing(func(*ui.Window) bool {
		n.doUserEvent()
		return true
	})
	newt.uiWindow = win
	newt.uiControl = win

	n.tk = newt
	win.Show()
	return
}

func (n *node) SetWindowTitle(title string) {
	log(debugToolkit, "toolkit NewWindow", n.Text, "title", title)
	win := n.tk.uiWindow
	if (win == nil) {
		log(debugError, "Error: no window", n.WidgetId)
	} else {
		win.SetTitle(title)
		log(debugToolkit, "Setting the window title", title)
	}
}
