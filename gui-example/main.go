package main

import (
	"log"
	"os"
	"time"

	"git.wit.org/wit/gui"
)

func customExit(gw *gui.GuiWindow) {
	log.Println("Should Exit Here")
	os.Exit(0)
}

func main() {
	log.Println("starting my Control Panel")

	gui.Config.Width = 1000
	gui.Config.Height = 400
	gui.Config.Exit = customExit

	go gui.Main(initGUI)

	watchGUI()
}

func initGUI() {
	gui.NewWindow("jcarr start", 640, 480)
}

func watchGUI() {
	var i = 1
	for {
		log.Println("Waiting for customExit()", i)
		i += 1
		time.Sleep(time.Second)
		if i == 3 {
			log.Println("Sending ExampleWindow to gui.Queue()")
			gui.Queue(gui.ExampleWindow)
		}
	}
}
