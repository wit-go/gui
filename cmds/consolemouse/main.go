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

func addDemoTab(w *gui.Node, title string) {
	var newNode, g *gui.Node

	newNode = w.NewTab(title)

	g = newNode.NewGroup("group 1")

	dd := g.NewDropdown("demoCombo2")
	dd.AddDropdown("more 1")
	dd.AddDropdown("less 2")
	dd.AddDropdown("foo 3")
}

func myExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	toolkit.Exit()
	os.Exit(0)
}
