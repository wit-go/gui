// This creates a simple hello world window
package main

import 	(
	"os"
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	f, err := os.OpenFile("/tmp/guilogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	gui.Main(initGUI)
}

// This initializes the first window
func initGUI() {
	var w *gui.Node
	gui.Config.Title = "Hello World golang wit/gui Window"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myDefaultExit

	w = gui.NewWindow()
	w.Dump()
	addDemoTab(w, "A Simple Tab Demo")
	addDemoTab(w, "A Second Tab")
}

func addDemoTab(window *gui.Node, title string) {
	var newNode, g *gui.Node

	newNode = window.NewTab(title)
        log.Println("addDemoTab() newNode.Dump")
        log.Println("addDemoTab() newNode.Dump")
        log.Println("addDemoTab() newNode.Dump")
        log.Println("addDemoTab() newNode.Dump")
	newNode.Dump()

	g = newNode.NewGroup("group 1")
        log.Println("addDemoTab() g.Dump")
        log.Println("addDemoTab() g.Dump")
        log.Println("addDemoTab() g.Dump")
        log.Println("addDemoTab() g.Dump")
	g.Dump()
	// os.Exit(0)
	dd := g.NewDropdown("demoCombo2")
	dd.AddDropdown("more 1")
	dd.AddDropdown("more 2")
	dd.AddDropdown("more 3")
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

