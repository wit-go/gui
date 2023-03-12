// This is a simple example
package main

import 	(
	"fmt"
	"log"
	"time"
	"strconv"
	"git.wit.org/wit/gui"
)

var title string = "Demo Plugin Window"
var outfile string = "/tmp/guilogfile"

// this is broken. delete this

func main() {
	// this set the xterm and mate-terminal window title. maybe works generally?
	fmt.Println("\033]0;" + title + "\007")
	// time.Sleep(5 * time.Second)
	// var w *gui.Node

	// this doesn't seem to work
	// captureSTDOUT()

	// gui.LoadToolkit("default")
	// panic("WTF gocui not happening")
	// gui.LoadToolkit("gocui")
//	gui.Init()

	// buttonWindow()
	go gui.Main(func () {
		log.Println("START Main f()")
		buttonWindow()
		/*
		log.Println("END NewWindow()")
		log.Println("START NewGroup()")
		g := w.NewGroup("new Group 22")
		log.Println("END NewGroup()")
		g.NewButton("asdjkl", func () {
			log.Println("world")
		})
		*/
		log.Println("END Main f()")
		// gui.StandardExit(nil)
	})
	log.Println("Main() END")
	time.Sleep(1 * time.Second)
	gui.Watchdog()
	gui.StandardExit()
}

var counter int = 5

// This creates a window
func buttonWindow() {
	var w, g *gui.Node
	gui.Config.Title = title
	gui.Config.Width = 640
	gui.Config.Height = 480

	w = gui.NewWindow()
	g = w.NewGroup("buttonGroup")

	g.NewButton("this app is old", func () {
	})
	g.NewLabel("STDOUT is set to: " + outfile)

	g.NewButton("hello", func () {
		log.Println("world")
	})

	g.NewButton("NewButton()", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "foobar " + strconv.Itoa(counter)
		counter += 1
		g.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewGroup()", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "neat " + strconv.Itoa(counter)
		counter += 1
		g.NewGroup(name)
	})

	g.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
	})

	g.NewButton("LoadToolkit(andlabs)", func () {
		gui.LoadToolkit("andlabs")
	})

	g.NewButton("LoadToolkit(gocui)", func () {
		gui.LoadToolkit("gocui")
	})

	g.NewButton("Init()", func () {
		log.Println("gui.Init() is deprecated(?)")
		//gui.Init()
	})

	g.NewButton("Main()", func () {
		go gui.Main(func () {
			w := gui.NewWindow()
			w.NewGroup("buttonGroup")
		})
	})
}
