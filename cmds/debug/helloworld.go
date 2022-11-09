// A simple helloworld window
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

// This creates a window
func helloworld() {
	var w *gui.Node
	gui.Config.Title = "helloworld golang wit/gui window"
	gui.Config.Width = 400
	gui.Config.Height = 100

	w = gui.NewWindow()
	w.NewButton("hello", func () {
		log.Println("world")
	})
}
