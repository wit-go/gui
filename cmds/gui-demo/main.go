package main

import (
	"log"
	"os"
	"time"

	"git.wit.org/wit/gui"
)

// This initializes the first window
//
// Then starts a goroutine to demonstrate how to
// inject things into the GUI
func main() {
	log.Println("Starting my Control Panel")

	go gui.Main(initGUI)

	watchGUI()
}

// This initializes the first window
func initGUI() {
	gui.ToolkitDemoWindow()
}

// This demonstrates how to properly interact with the GUI
// You can not involke the GUI from external goroutines in most cases.
func watchGUI() {
	var i = 1
	for {
		log.Println("Waiting", i, "seconds")
		i += 1
		time.Sleep(1 * time.Second)
		if i == 4 {
			log.Println("Opening a Debug Window via the gui.Queue()")
			gui.Config.Width = 800
			gui.Config.Height = 300
			gui.Config.Exit = myDebugExit
			gui.Queue(gui.DebugWindow)
		}
	}
}

func myExit(n *gui.Node) {
	log.Println()
	log.Println("Entered myExit() on node.Name =", n.Name)
	log.Println()
	os.Exit(0)
}

func myDebugExit(n *gui.Node) {
	log.Println("Entered myDebugExit() on node.Name =", n.Name)
	log.Println("Don't actually os.Exit()")
}
