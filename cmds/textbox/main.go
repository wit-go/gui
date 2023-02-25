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
	// GuiDebug bool `help:"open up the wit/gui Debugging Window"`
	// GuiDemo bool `help:"open the wit/gui Demo Window"`
	User string `arg:"env:USER"`
}

var args struct {
	LogOptions
	gui.GuiArgs
}

func main() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)
	log.Println("Toolkit = ", args.Toolkit)

	/*
	f, err := os.OpenFile("/tmp/guilogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
	*/

	// gui.InitPlugins([]string{"andlabs"})
	gui.Main(initGUI)
}

// This initializes the first window
func initGUI() {
	var w *gui.Node
	gui.Config.Title = "Hello World"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myDefaultExit

	w = gui.NewWindow()
	w.Dump()
	addDemoTab(w, "A Simple Tab Demo")
	addDemoTab(w, "A Second Tab")

	if (args.GuiDebug) {
		gui.DebugWindow()
	}
}

func addDemoTab(window *gui.Node, title string) {
	var newNode, g, g2, tb *gui.Node

	newNode = window.NewTab(title)
        log.Println("addDemoTab() newNode.Dump")
	newNode.Dump()

	g = newNode.NewGroup("group 1")
	dd := g.NewDropdown("demoCombo2")
	dd.AddDropdownName("more 1")
	dd.AddDropdownName("more 2")
	dd.AddDropdownName("more 3")
	dd.OnChanged = func(*gui.Node) {
		s := dd.GetText()
		tb.SetText("hello world " + args.User + "\n" + s)
	}

	g2 = newNode.NewGroup("group 2")
	tb = g2.NewTextbox("tb")
	log.Println("tb =", tb.GetText())
	tb.OnChanged = func(*gui.Node) {
		s := tb.GetText()
		log.Println("text =", s)
	}
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

