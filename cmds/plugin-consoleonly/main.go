// This is a simple example
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	gui.InitPlugins([]string{"gocui"})
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
	w.NewButton("Flags", func () {
		log.Println("the debugging flags window")
		w.DebugFlags(false)
	})
	w.NewButton("Widgets", func () {
		w.DebugWidgets(false)
		log.Println("debug the widgets window")
	})
}
