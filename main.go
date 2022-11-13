package gui

import (
	"log"
	"os"
	"embed"
)

// Windows doesn't support plugins. How can I keep andlabs and only compile it on windows?
// https://forum.heroiclabs.com/t/setting-up-goland-to-compile-plugins-on-windows/594/5
// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

const Xaxis = 0 // stack things horizontally
const Yaxis = 1 // stack things vertically

// may this plugin work when all other plugins fail
//go:embed toolkit/gocui.so
var res embed.FS

func init() {
	log.Println("gui.init() has been run")

	Config.counter = 0
	Config.prefix = "wit"

	// Config.Debug.Debug = true
	// Config.Debug.Node = true
	// Config.Debug.Tabs = true

	title := "guiBinaryTree"
	w     := 640
	h     := 480

	// Populates the top of the binary tree
	Config.master = addNode(title, w, h)
	if (Config.Debug.Debug) {
		Config.master.Dump()
	}
}

func Init() {
	var initBAD bool = true

	if (Config.Debug.Debug) {
		log.Println("Starting gui.Init()")
	}
	for _, aplug := range allPlugins {
		log.Println("gui.LoadToolkit() already loaded toolkit plugin =", aplug.name)
		initBAD = false
	}

	// the program didn't specify a plugin. Try to load one
	// TODO: detect the OS & user preferences to load the best one
	if (initBAD) {
		if (LoadToolkit("andlabs2")) {
			initBAD = false
		}
	}

	// andlabs2 gui failed. fall back to the terminal gui (should be compiled into the binary)
	if (initBAD) {
		if (LoadToolkit("gocui")) {
			initBAD = false
		}
	}

	// locate the shared library file
	// panic("WTF Init()")
	for _, aplug := range allPlugins {
		log.Println("gui.Node.Init() toolkit plugin =", aplug.name)
		if (aplug.InitOk) {
			log.Println("gui.Node.Init() Already Ran Init()", aplug.name)
			continue
		}
		if (aplug.Init == nil) {
			log.Println("gui.Node.Main() Init == nil", aplug.name)
			continue
		}
		aplug.InitOk = true
		aplug.Init()
	}
	// StandardExit(nil)
}

// This should not pass a function
func Main(f func()) {
	if (Config.Debug.Debug) {
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	}
	for _, aplug := range allPlugins {
		log.Println("gui.Node.NewButton() toolkit plugin =", aplug.name)
		if (aplug.MainOk) {
			log.Println("gui.Node.Main() Already Ran Main()", aplug.name)
			continue
		}
		if (aplug.Main == nil) {
			log.Println("gui.Node.Main() Main == nil", aplug.name)
			continue
		}
		aplug.MainOk = true
		aplug.Main(f)
		// f()
	}
	// toolkit.Main(f)
}

// This should never be exposed(?)

// Other goroutines must use this to access the GUI
//
// You can not acess / process the GUI thread directly from
// other goroutines. This is due to the nature of how
// Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
// For example: gui.Queue(NewWindow())
func queue(f func()) {
	log.Println("Sending function to gui.Main() (using gtk via andlabs/ui)")
	// toolkit.Queue(f)
	for _, aplug := range allPlugins {
		log.Println("gui.Node.NewButton() toolkit plugin =", aplug.name)
		if (aplug.Queue == nil) {
			continue
		}
		aplug.Queue(f)
	}
}

// The window is destroyed but the application does not quit
func StandardClose(n *Node) {
	if (Config.Debug.Debug) {
		log.Println("wit/gui Standard Window Close. name =", n.Name)
	}
}

// The window is destroyed but the application does not quit
func StandardExit(n *Node) {
	if (Config.Debug.Debug) {
		log.Println("wit/gui Standard Window Exit. running os.Exit()")
	}

	log.Println("gui.Node.StandardExit() attempt to exit each toolkit plugin")
	for i, aplug := range allPlugins {
		log.Println("gui.Node.NewButton()", i, aplug)
		if (aplug.Quit != nil) {
			aplug.Quit()
		}
	}

	os.Exit(0)
}
