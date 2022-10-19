package gui

import "log"
// import "time"
import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func NewStandardWindow(title string) *Node {
	log.Println("NewStandardWindow() creating", title)

	Config.Title = title
	Config.Width = 640
	Config.Height = 480
	Config.Exit = StandardClose
	return NewWindow()
}

//
// This creates a window that shows how the toolkit works
// internally using it's raw unchanged code for the toolkit itself
// 
// This is a way to test and see if the toolkit is working at all
// right now it shows the andlabs/ui/DemoNumbersPage()
//
func DemoToolkitWindow() {
	var w, d *Node
	var tk *toolkit.Toolkit

	w = NewStandardWindow("Demo of the GUI Toolkit")

	d = w.New("demo")

	tk = toolkit.DemoNumbersPage(w.uiWindow)
	tk.OnChanged = func(t *toolkit.Toolkit) {
		log.Println("toolkit.NewSlider() value =", t.Value())
		if (d.OnChanged != nil) {
			log.Println("toolkit.Demo() running node.OnChanged")
			d.OnChanged(d)
		}
	}
	d.Toolkit = tk

	log.Println("ToolkitDemoWindow() END")
}
