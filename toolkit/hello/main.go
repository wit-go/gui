// This creates a simple hello world window
package main

import 	(
	"os"
	"log"
	"git.wit.org/wit/gui"
)

func Init() {
	gui.Main(myGUI)
}

// This initializes the first window
func myGUI() {
	var w *gui.Node
	gui.Config.Title = "Hello World golang wit/gui Window"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myExit

	w = gui.NewWindow()
	addHelloWorld(w, "A Simple Tab")
}

func addHelloWorld(window *gui.Node, title string) {
	var newNode, g, tb *gui.Node

	newNode = window.NewTab(title)

	g = newNode.NewGroup("hello")
	tb = g.NewTextbox("hello world box") // when debugging, this string will be used
	tb.OnChanged = func(*gui.Node) {
		s := tb.GetText()
		log.Println("text box =", s)
	}
	tb.SetText("world")
}

func myExit(n *gui.Node) {
        log.Println("exit() here")
	os.Exit(0)
}

