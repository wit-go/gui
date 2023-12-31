// This is a simple example
package main

import 	(
	"log"
	"strconv"
	"go.wit.com/gui"
	arg "github.com/alexflint/go-arg"
)

var title string = "Demo Plugin Window"
var myGui *gui.Node

var buttonCounter int = 5
var gridW int = 5
var gridH int = 3

func init() {
	arg.MustParse()
}

func main() {
	// This will turn on all debugging
	// gui.SetDebug(true)

	myGui = gui.New().Default()
	buttonWindow()

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()
}

// This creates a window
func buttonWindow() {
	var w, t, g, more, more2 *gui.Node

	log.Println("buttonWindow() START")

	w = myGui.NewWindow(title).SetText("Nueva Ventana de Botones")
	t = w.NewTab("buttonTab is this thing").Pad()
	g = t.NewGroup("buttonGroup").Pad()
	g1 := t.NewGroup("buttonGroup 2").Pad()
	more = g1.NewGroup("more").Pad()
	g1.NewButton("hello2", func () {
		log.Println("world2")
	})
	more2 = g1.NewGrid("gridnuts", gridW, gridH).Pad()

	more2.NewLabel("more2")

	g.NewButton("hello", func () {
		log.Println("world")
	})

	g.NewButton("NewButton(more)", func () {
		name := "foobar " + strconv.Itoa(buttonCounter)
		log.Println("NewButton(more) Adding button", name)
		buttonCounter += 1
		more.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewButton(more2)", func () {
		name := "foobar " + strconv.Itoa(buttonCounter)
		log.Println("NewButton(more2) Adding button", name)
		buttonCounter += 1
		more2.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewButton(more2 d)", func () {
		name := "d" + strconv.Itoa(buttonCounter)
		log.Println("NewButton(more2 d) Adding button", name)
		buttonCounter += 1
		more2.NewButton(name, func () {
			log.Println("Got all the way to main() name =", name)
		})
	})

	g.NewButton("NewGroup()", func () {
		name := "neat " + strconv.Itoa(buttonCounter)
		log.Println("NewGroup() Adding button", name)
		buttonCounter += 1
		more.NewGroup(name).Pad()
	})
}
