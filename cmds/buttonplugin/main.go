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

func main() {
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
	var w, t, g, more, more2 *gui.Node
	gui.Config.Title = title
	gui.Config.Width = 640
	gui.Config.Height = 480

	w = gui.NewWindow()
	t = w.NewTab("buttonTab")
	g = t.NewGroup("buttonGroup")
	g1 := t.NewGroup("buttonGroup 2")
	more = g1.NewGroup("more")
	g1.NewButton("hello2", func () {
		log.Println("world2")
	})
	more2 = g1.NewGrid("gridnuts", 3, 3)

	more2.NewLabel("more2")

	g.NewButton("this app is useful for plugin debuggin", func () {
	})
	g.NewLabel("STDOUT is set to: " + outfile)

	g.NewButton("hello", func () {
		log.Println("world")
	})

	g.NewButton("Load 'gocui'", func () {
		// this set the xterm and mate-terminal window title. maybe works generally?
		fmt.Println("\033]0;" + title + "blah \007")
		gui.StartS("gocui")
	})

	g.NewButton("Redraw 'gocui'", func () {
		fmt.Println("\033]0;" + title + "blah2 \007")
		gui.Redraw("gocui")
	})

	g.NewButton("NewButton(more)", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "foobar " + strconv.Itoa(counter)
		counter += 1
		more.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewButton(more2)", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "foobar " + strconv.Itoa(counter)
		counter += 1
		more2.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewButton(more2 d)", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "d" + strconv.Itoa(counter)
		counter += 1
		more2.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewGroup()", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "neat " + strconv.Itoa(counter)
		counter += 1
		more.NewGroup(name)
	})

	g.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
	})
}
