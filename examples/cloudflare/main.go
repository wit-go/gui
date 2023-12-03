// This is a simple example
package main

import 	(
	"os"
	"fmt"
	"log"
	"git.wit.org/wit/gui"
)

var title string = "Cloudflare DNS Control Panel"
var outfile string = "/tmp/guilogfile"
var myGui *gui.Node

var buttonCounter int = 5
var gridW int = 5
var gridH int = 3

var mainWindow, more, more2 *gui.Node

func main() {
	myGui = gui.New().Default()
	buttonWindow()

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()
}

// This creates a window
func buttonWindow() {
	var t, g *gui.Node

	log.Println("buttonWindow() START")

	mainWindow = myGui.NewWindow(title).SetText(title)
	t = mainWindow.NewTab("Cloudflare")
	g = t.NewGroup("buttons")
	g1 := t.NewGroup("buttonGroup 2")

	more = g1.NewGroup("more")
	showCloudflareCredentials(more)

	// more2 = g1.NewGrid("gridnuts", gridW, gridH)

	var domain string = os.Getenv("CLOUDFLARE_DOMAIN")
	if (domain == "") {
		domain = "example.org"
	}

	g.NewButton("Load " + domain + " DNS", func () {
		loadDNS(domain)
	})

	g.NewButton("Load 'gocui'", func () {
		// this set the xterm and mate-terminal window title. maybe works generally?
		fmt.Println("\033]0;" + title + "blah \007")
		myGui.LoadToolkit("gocui")
	})

	g.NewButton("Load 'andlabs'", func () {
		myGui.LoadToolkit("andlabs")
	})

	g.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
	})
}

func showCloudflareCredentials(box *gui.Node) {
	grid := box.NewGrid("credsGrid", 2, 4) // width = 2

	grid.NewLabel("Domain")
	grid.NewLabel(os.Getenv("CLOUDFLARE_DOMAIN"))

	grid.NewLabel("Auth Key")
	grid.NewLabel(os.Getenv("CLOUDFLARE_AUTHKEY"))

	grid.NewLabel("Email")
	grid.NewLabel(os.Getenv("CLOUDFLARE_EMAIL"))

	grid.NewLabel("URL")
	grid.NewLabel(os.Getenv("CLOUDFLARE_URL"))
}
