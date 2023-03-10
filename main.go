package gui

import (
	// "embed"
	"git.wit.org/wit/gui/toolkit"
)

// Windows doesn't support plugins. How can I keep andlabs and only compile it on windows?
// https://forum.heroiclabs.com/t/setting-up-goland-to-compile-plugins-on-windows/594/5
// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

const Xaxis = 0 // stack things horizontally
const Yaxis = 1 // stack things vertically

/*
	// TODO: 2023/03/03 rethink how to get a plugin or figure out how
	// golang packages can include a binary. Pull from /usr/go/go-gui/ ?
	// may this plugin work when all other plugins fail

	// if this is in the plugin, the packages can't work with go.mod builds
	# don't do this in the plugin // go:embed /usr/lib/go-gui/toolkit/gocui.so
	# don't do this in the plugin var res embed.FS
*/

func init() {
	log("init() has been run")

	Config.counter = 0
	Config.prefix = "wit"
	Config.Width = 640
	Config.Height = 480

	// Populates the top of the binary tree
	Config.master = addNode("guiBinaryTree")

	// used to pass debugging flags to the toolkit plugins
	Config.flag = Config.master.New("flag", toolkit.Flag, nil)
}

func doGuiChan() {
	for {
		select {
		case <-Config.ActionCh1:
			log(true, "CHANNEL ACTION 1  !!!!!")
			return
		case <-Config.ActionCh2:
			log(true, "CHANNEL ACTION 2  !!!!!")
			return
		default:
			log(true, "doGuiChan() nothing")
		}
		log(true, "doGuiChan() for()")
	}
}

func InitPlugins(names []string) {
	log(debugGui, "Starting gui.Init()")

	for _, aplug := range allPlugins {
		log(debugGui, "gui.LoadToolkit() already loaded toolkit plugin =", aplug.name)
		for _, name := range names {
			if (name == aplug.name) {
				return
			}
		}
	}

	// try to load each plugin in the order passed to this function
	for _, name := range names {
		if (LoadToolkit(name)) {
			// aplug.InitOk = true
			// aplug.Init()
			return
		}
	}

	// the program didn't specify a plugin. Try to load one
	// TODO: detect the OS & user preferences to load the best one
	// TODO: commented out Init() on 02/26/2023 because I'm not sure how to run it correctly
	if (LoadToolkit("andlabs")) {
		// aplug.InitOk = true
		// aplug.Init()
		return
	}

	if (LoadToolkit("gocui")) {
		// aplug.InitOk = true
		// aplug.Init()
		return
	}

	// Should die here? TODO: need a Node to call StandardExit
	// StandardExit("golang wit/gui could not load a plugin TODO: do something to STDOUT (?)")
}

// This should not pass a function
func Main(f func()) {
	log(debugGui, "Starting gui.Main() (using gtk via andlabs/ui)")

	InitPlugins([]string{"andlabs", "gocui"})

	for _, aplug := range allPlugins {
		log(debugGui, "gui.Node.NewButton() toolkit plugin =", aplug.name)
		if (aplug.MainOk) {
			log(debugGui, "gui.Node.Main() Already Ran Main()", aplug.name)
			continue
		}
		if (aplug.Main == nil) {
			log(debugGui, "gui.Node.Main() Main == nil", aplug.name)
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
func Queue(f func()) {
	log(debugGui, "Sending function to gui.Main() (using gtk via andlabs/ui)")
	// toolkit.Queue(f)
	for _, aplug := range allPlugins {
		log(debugGui, "gui.Node.NewButton() toolkit plugin =", aplug.name)
		if (aplug.Queue == nil) {
			continue
		}
		aplug.Queue(f)
	}
}

// The window is destroyed but the application does not quit
func (n *Node) StandardClose() {
	log(debugGui, "wit/gui Standard Window Close. name =", n.Name)
	log(debugGui, "wit/gui Standard Window Close. n.Custom exit =", n.Custom)
}

// The window is destroyed and the application exits
// TODO: properly exit the plugin since Quit() doesn't do it
func StandardExit() {
	log("wit/gui Standard Window Exit. running os.Exit()")
	log("gui.Node.StandardExit() attempt to exit each toolkit plugin")
	for i, aplug := range allPlugins {
		log("gui.Node.NewButton()", i, aplug)
		if (aplug.Quit != nil) {
			aplug.Quit()
		}
	}
	exit(0)
}
