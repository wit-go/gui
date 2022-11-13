package gui

import "log"
// import "time"
// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

//
// This creates a window that shows how this package works
// 
func DemoWindow() {
	var w, t1 *Node
	log.Println("DemoWindow() START")

	w = NewStandardWindow("Demo of WIT/GUI")

	t1 = w.DebugTab("WIT GUI Debug Tab t1")
	t1.DebugTab("WIT GUI Debug Tab t2")

	log.Println("DemoWindow() END")
}

func NewStandardWindow(title string) *Node {
	log.Println("NewStandardWindow() creating", title)

	Config.Title = title
	Config.Width = 640
	Config.Height = 480
	Config.Exit = StandardClose
	return NewWindow()
}
