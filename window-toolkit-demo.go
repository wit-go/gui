package gui

import "log"
import "time"

func NewStandardWindow(title string) *Node {
	log.Println("NewStandardWindow() creating", title)

	Config.Title = title
	Config.Width = 640
	Config.Height = 480
	Config.Exit = StandardClose
	return NewWindow()
}
func ToolkitDemoWindow() {
	var w, t, g *Node

	w = NewStandardWindow("Demo the GUI Toolkit")

	// w.DemoAndlabsUiTab("ran AddDemoAndlabsUiTab()")
	t = w.AddTab("Set time delay", nil)
	g = t.AddGroup("nanoseconds")

	s := g.NewSlider("t", 2, 80)
	s.OnChanged = func (td *Node) {
		t :=  time.Duration(s.Int())
		log.Println("ToolkitDemoWindow() OnChanged() delay =", t);
	}

	log.Println("ToolkitDemoWindow() END")
}
