// This is a simple example
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	gui.Init()
	gui.Main(helloworld)
}

// This creates a window
func helloworld() {
	var w *gui.Node
	gui.Config.Title = "helloworld golang wit/gui window"
	gui.Config.Width = 640
	gui.Config.Height = 480

	w = gui.NewWindow()
	w.NewButton("hello", func () {
		log.Println("world")
	})
}
