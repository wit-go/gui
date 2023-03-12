// This creates a simple hello world window
package main

import 	(
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

func init() {
	arg.MustParse(&args)
	log.Println("Toolkit = ", args.Toolkit)

	if (args.GuiDebug) {
		gui.DebugWindow()
	}
	if (args.GuiVerbose) {
		gui.SetDebug(true)
	}
}
