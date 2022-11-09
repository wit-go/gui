// This is a simple example
package main

import 	(
	"log"
	"strconv"
	"git.wit.org/wit/gui"
)

func main() {
	// this doesn't seem to work
	captureSTDOUT()

	gui.Main(buttonWindow)
}

var counter int = 10

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

	g.NewButton("RunGreet()", func () {
		log.Println("world")
		go gui.RunGreet()
	})

	g.NewButton("gui.LookupJcarrButton()", func () {
		log.Println("gui.LookupJcarrButton()")
		gui.LookupJcarrButton()
	})

	g.NewButton("new foobar 2", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "foobar " + strconv.Itoa(counter)
		counter += 1
		g.NewButton(name, nil)
	})
}
