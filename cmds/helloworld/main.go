// This creates a simple hello world window
package main

import 	(
	"os"
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	gui.Main(initGUI)
}

// This initializes the first window
func initGUI() {
	gui.Config.Title = "Hello World golang wit/gui Window"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myDefaultExit

	node1 := gui.NewWindow()
	addDemoTab(node1, "A Simple Tab Demo")
	addDemoTab(node1, "A Second Tab")
}

func addDemoTab(n *gui.Node, title string) {
	newNode := n.AddTab(title, nil)

	g := newNode.NewGroup("group 1")
	// g.Dump()
	g.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

