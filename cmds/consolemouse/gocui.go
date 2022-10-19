// This creates a simple hello world window
package main

import 	(
	"log"
	"time"
	"git.wit.org/wit/gui"
)

import toolkit "git.wit.org/wit/gui/toolkit/gocui"

func configureGogui() {
	toolkit.Init()
	toolkit.OnExit(mycallback)
}

func startGogui() {
	toolkit.StartConsoleMouse()
}

func mycallback(name string) {
	log.Println("run andlabs here? name =", name)
	if (name == "andlabs") {
		go gui.Main(initGUI)
	}
	if (name == "something") {
		log.Println("add something to do here")
	}
	if (name == "DemoToolkitWindow") {
		gui.Queue( func () {
			gui.DemoToolkitWindow()
		})
	}
	if (name == "addDemoTab") {
		gui.Queue( func () {
			addDemoTab(w, "A Tab from gocui")
		})
	}
	if (name == "DebugWindow") {
			log.Println("Opening a Debug Window via the gui.Queue()")
			gui.Config.Width = 800
			gui.Config.Height = 300
			gui.Config.Exit = myExit
			gui.Queue(gui.DebugWindow)
			time.Sleep(1 * time.Second)
			gui.Queue(gui.DebugTab)
	}
	if (name == "exit") {
		myExit(nil)
	}
}
