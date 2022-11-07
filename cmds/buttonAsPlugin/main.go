// This is a simple example
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	// go loadPlugin(plugHello, "../../toolkit/hello.so")

	// this doesn't seem to work
	captureSTDOUT()

	// go loadPlugin("../../toolkit/gocli.so")
	gui.Main(buttonWindow)
}

// This creates a window
func buttonWindow() {
	var w, g *gui.Node
	gui.Config.Title = "Demo Plugin Window"
	gui.Config.Width = 640
	gui.Config.Height = 480

	w = gui.NewWindow()
	g = w.NewGroup("buttonGroup")

	g.NewButton("hello", func () {
		log.Println("world")
	})

	/*
	g.NewButton("LoadPlugin()", func () {
		log.Println("world")
		gui.LoadPlugin("../../toolkit/gocli.so")
	})
	*/

	g.NewButton("RunGreet()", func () {
		log.Println("world")
		go gui.RunGreet()
	})

	g.NewButton("gui.LookupJcarrButton()", func () {
		log.Println("gui.LookupJcarrButton()")
		gui.LookupJcarrButton()
	})

	g.NewButton("gui.GocuiAddButton()", func () {
		log.Println("gui.GocuiAddButton()")
		gui.GocuiAddButton("new foobar")
	})
}
