// This creates a simple hello world window
package main

import 	(
	"os"
	"log"
	"git.wit.org/wit/gui"
	arg "github.com/alexflint/go-arg"
)

type LogOptions struct {
	LogFile string
	Verbose bool
	User string `arg:"env:USER"`
}

var args struct {
	LogOptions
	gui.GuiArgs
}

func main() {
	// this implements ./cmd --gui-debug --gui-toolkit, etc
	arg.MustParse(&args)
	log.Println("Toolkit = ", args.Toolkit)

	// gui.InitPlugins([]string{"andlabs"})
	gui.Main(initGUI)
}

// This initializes the first window
func initGUI() {
	var w *gui.Node
	gui.Config.Title = "Hello World"
	gui.Config.Width = 642
	gui.Config.Height = 481
	gui.Config.Exit = myDefaultExit

	w = gui.NewWindow()
	w.Custom = func () {
		log.Println("myDefaultExit(w)")
		myDefaultExit(w)
	}
	addDemoTab(w, "A Simple Tab Demo")
	addDemoTab(w, "A Second Tab")

	if (args.GuiDebug) {
	 	gui.DebugWindow()
	}
	if (args.GuiVerbose) {
	 	gui.SetDebug(true)
	}
}

func addDemoTab(window *gui.Node, title string) {
	var newNode, g *gui.Node

	newNode = window.NewTab(title)
        log.Println("addDemoTab() newNode START")
	// newNode.Dump(true)

	g = newNode.NewGroup("group 1")

	g1 := g.NewGrid("grid 1", 2, 2)
	g1.NewLabel("less")
	dd := g1.NewDropdown("more")
	dd.AddDropdownName("more 1")
	dd.AddDropdownName("more 2")
	dd.AddDropdownName("more 3")

	// g.SetNext(3,1)
	// g1.NewLabel("label (3,1)")
	// g.SetNext(3,2)
	// g1.NewLabel("label (3,2)")

	g2 := newNode.NewGroup("group 2")
	tb := g2.NewTextbox("tb")
	log.Println("tb =", tb.GetText())
	tb.Custom = func() {
		s := tb.GetText()
		log.Println("text =", s)
	}

	dd.Custom = func() {
		s := dd.GetText()
		log.Println("hello world " + args.User + "\n" + s + "\n")
		tb.SetText("hello world " + args.User + "\n" + s + "\n")
	}
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

