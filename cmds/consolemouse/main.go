// This creates a simple hello world window
package main

import 	(
	"os"
	"log"
//	"time"
	"git.wit.org/wit/gui"
)

import toolkit "git.wit.org/wit/gui/toolkit/gocui"

var w *gui.Node

func main() {
	go gui.Main(initGUI)

	configureGogui()
	startGogui()
}

// This initializes the first window
func initGUI() {
	gui.Config.Title = "Hello World golang wit/gui Window"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myExit

	w = gui.NewWindow()
	w.Dump()
	addDemoTab(w, "A Simple Tab Demo")
	addDemoTab(w, "A Second Tab")
}

func addDemoTab(window *gui.Node, title string) {
	var newNode, g *gui.Node

	newNode = window.AddTab(title, nil)
        log.Println("addDemoTab() newNode.Dump")
        log.Println("addDemoTab() newNode.Dump")
        log.Println("addDemoTab() newNode.Dump")
        log.Println("addDemoTab() newNode.Dump")
	newNode.Dump()
	newNode.Toolkit.Dump()

	g = newNode.NewGroup("group 1")
        log.Println("addDemoTab() g.Dump")
        log.Println("addDemoTab() g.Dump")
        log.Println("addDemoTab() g.Dump")
        log.Println("addDemoTab() g.Dump")
	g.Dump()
	g.Toolkit.Dump()
	// myExit(nil)
	g.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
}

func myExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	toolkit.Exit()
	os.Exit(0)
}
