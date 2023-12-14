// This is a simple example
package main

import 	(
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"git.wit.org/wit/gui"
)

var title string = "Cloudflare DNS Control Panel"
var outfile string = "/tmp/guilogfile"
var configfile string = ".config/wit/cloudflare"
var myGui *gui.Node

var buttonCounter int = 5
var gridW int = 5
var gridH int = 3

var mainWindow, more, more2 *gui.Node

func main() {
	config = make(map[string]*configT)
	readConfig()
	myGui = gui.New().Default()
	makeCloudflareWindow()

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()
}

// This creates a window
func makeCloudflareWindow() {
	var t *gui.Node

	log.Println("buttonWindow() START")

	mainWindow = myGui.NewWindow(title).SetText(title)

	// this tab has the master cloudflare API credentials
	makeConfigTab(mainWindow)

	t = mainWindow.NewTab("Zones")
	g1 := t.NewGroup("zones")

	// make dropdown list of zones
	zonedrop = g1.NewDropdown("zone")
	zonedrop.AddText("example.org")
	for d, _ := range config {
		zonedrop.AddText(d)
	}

	zonedrop.Custom = func () {
		domain := zonedrop.S
		log.Println("custom dropdown() zone (domain name) =", zonedrop.Name, domain)
		if (config[domain] == nil) {
			log.Println("custom dropdown() config[domain] = nil for domain =", domain)
			domainWidget.SetText(domain)
			zoneWidget.SetText("")
			authWidget.SetText("")
			emailWidget.SetText("")
		} else {
			log.Println("custom dropdown() a =", domain, config[domain].zoneID, config[domain].auth, config[domain].email)
			domainWidget.SetText(config[domain].domain)
			zoneWidget.SetText(config[domain].zoneID)
			authWidget.SetText(config[domain].auth)
			emailWidget.SetText(config[domain].email)
		}
	}

	more = g1.NewGroup("data")
	showCloudflareCredentials(more)

	makeDebugTab(mainWindow)
}

func makeConfigTab(window *gui.Node) {
	t := window.NewTab("Get Zones")
	vb := t.NewBox("vBox", false)
	g1 := vb.NewGroup("Cloudflare API Config")

	g1.NewLabel("If you have an API key with access to list all of /n your zone files, enter it here. \n \n Alternatively, you can set the enviroment variables: \n env $CLOUDFLARE_AUTHKEY \n env $CLOUDFLARE_EMAIL \n env $CLOUDFLARE_URL \n")

	// make grid to display credentials
	grid := g1.NewGrid("credsGrid", 2, 4) // width = 2

	grid.NewLabel("Auth Key")
	aw := grid.NewEntryLine(os.Getenv("CLOUDFLARE_AUTHKEY"))
	aw.SetText(os.Getenv("CLOUDFLARE_AUTHKEY"))

	grid.NewLabel("Email")
	ew := grid.NewEntryLine(os.Getenv("CLOUDFLARE_EMAIL"))
	ew.SetText(os.Getenv("CLOUDFLARE_EMAIL"))

	var url string = "https://api.cloudflare.com/client/v4/zones/"
	grid.NewLabel("Cloudflare API")
	grid.NewLabel(url)

	grid.Pad()

	vb.NewButton("getZones()", func () {
		log.Println("getZones()")
		getZones(aw.S, ew.S)
	})

	t.Pad()
	t.Margin()
	vb.Pad()
	vb.Margin()
	g1.Pad()
	g1.Margin()
}

func makeDebugTab(window *gui.Node) {
	t2 := window.NewTab("debug")
	g := t2.NewGroup("debug")
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

	g.NewButton("List all Widgets", func () {
		myGui.ListChildren(true)
	})
	g.NewButton("Dump all Widgets", func () {
		myGui.Dump()
	})
}

func showCloudflareCredentials(box *gui.Node) {
	// make grid to display credentials
	grid := box.NewGrid("credsGrid", 2, 4) // width = 2

	grid.NewLabel("Domain")
	domainWidget = grid.NewEntryLine(os.Getenv("CLOUDFLARE_DOMAIN"))

	grid.NewLabel("Zone ID")
	zoneWidget = grid.NewEntryLine(os.Getenv("CLOUDFLARE_ZONEID"))

	grid.NewLabel("Auth Key")
	authWidget = grid.NewEntryLine(os.Getenv("CLOUDFLARE_AUTHKEY"))

	grid.NewLabel("Email")
	emailWidget = grid.NewEntryLine(os.Getenv("CLOUDFLARE_EMAIL"))

	var url string = "https://api.cloudflare.com/client/v4/zones/"
	grid.NewLabel("Cloudflare API")
	grid.NewLabel(url)

	grid.Pad()

	saveButton = box.NewButton("Save to config", func () {
	})
	saveButton.Disable()

	loadButton = box.NewButton("Load Cloudflare DNS zonefile", func () {
		var domain configT
		domain.domain = domainWidget.S
		domain.zoneID = zoneWidget.S
		domain.auth = authWidget.S
		domain.email = emailWidget.S
		loadDNS(&domain)
	})
}

func readConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("searchPaths() error. exiting here?")
	}
	filename := homeDir + "/" + configfile
	log.Println("filename =", filename)

	readFileLineByLine(filename)
	// os.Exit(0)
}

// readFileLineByLine opens a file and reads through each line.
func readFileLineByLine(filename string) error {
	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Println("readFileLineByLine() =", filename)

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(file)

	// Read through each line using scanner.
	for scanner.Scan() {
		var newc *configT
		newc = new(configT)

		line := scanner.Text()
		parts := strings.Fields(line)

		if (len(parts) < 4) {
			log.Println("readFileLineByLine() SKIP =", parts)
			continue
		}

		newc.domain = parts[0]
		newc.zoneID = parts[1]
		newc.auth = parts[2]
		newc.email = parts[3]

		config[parts[0]] = newc
		log.Println("readFileLineByLine() =", newc.domain, newc.zoneID, newc.auth, newc.email)
	}

	// Check for errors during Scan.
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
