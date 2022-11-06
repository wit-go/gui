// This is a simple example
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	gui.Main(buttonWindow)
}

// This creates a window
func buttonWindow() {
	var w, g *gui.Node
	gui.Config.Title = "Demo Plugin Window"
	gui.Config.Width = 640
	gui.Config.Height = 480

//	gui.Config.Exit = gui.StandardClose
//	gui.SetDebug(true)

	w = gui.NewWindow()
	g = w.NewGroup("buttonGroup")

	g.NewButton("hello", func () {
		log.Println("world")
	})

	g.NewButton("foo", func () {
		log.Println("bar")
	})
}
