// This is a simple example
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

func main() {
	helloworld()
	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
}

// This creates a window
func helloworld() {
	myGui := gui.New().Default()
	myWindow := myGui.NewWindow("helloworld golang wit/gui window")

	myWindow.NewButton("hello", func () {
		log.Println("world")
	})
}
