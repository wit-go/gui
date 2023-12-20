// This creates a simple hello world window
package main

import 	(
	"fmt"
	arg "github.com/alexflint/go-arg"
	"go.wit.com/gui"
	log "go.wit.com/gui/log"
)


var args struct {
	Foo string
	Bar bool
	User string `arg:"env:USER"`
	Demo bool `help:"run a demo"`
	gui.GuiArgs
	log.LogArgs
}

func init() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar, args.User)

	if (args.Gui != "") {
		gui.GuiArg.Gui = args.Gui
	}
	log.Log(true, "INIT() args.GuiArg.Gui =", gui.GuiArg.Gui)

}
