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
var myGui *gui.Node

var buttonCounter int = 5

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
	gui.SetDebug(true)
	myGui = gui.Start()
	time.Sleep(1 * time.Second)
	buttonWindow()
	log.Println("Main() END")
	time.Sleep(1 * time.Second)
	// gui.StartS("gocui")
	// gui.Redraw("gocui")
	gui.Watchdog()
	gui.StandardExit()
}

// This creates a window
func buttonWindow() {
	var w, t, g, more, more2 *gui.Node

	log.Println("buttonWindow()")
	log.Println("buttonWindow()")
	log.Println("buttonWindow()")

	w = myGui.NewWindow2(title).SetText("Nueva Ventana de Botones")
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
		gui.Redraw("gocui")
	})

	g.NewButton("Load 'andlabs'", func () {
		gui.StartS("andlabs")
	})

	g.NewButton("NewButton(more)", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "foobar " + strconv.Itoa(buttonCounter)
		buttonCounter += 1
		more.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewButton(more2)", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "foobar " + strconv.Itoa(buttonCounter)
		buttonCounter += 1
		more2.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewButton(more2 d)", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "d" + strconv.Itoa(buttonCounter)
		buttonCounter += 1
		more2.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewGroup()", func () {
		log.Println("new foobar 2. Adding button 'foobar 3'")
		name := "neat " + strconv.Itoa(buttonCounter)
		buttonCounter += 1
		more.NewGroup(name)
	})

	g.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
	})
}
