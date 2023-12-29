package gui

import (
	arg "github.com/alexflint/go-arg"
)

var guiArg GuiArgs

// This struct can be used with the go-arg package
type GuiArgs struct {
	Gui string `arg:"--gui" help:"Use this gui toolkit [andlabs,gocui,nocui]"`
	GuiDebug bool `arg:"--gui-debug" help:"open the GUI debugger"`
	GuiVerbose bool `arg:"--gui-verbose" help:"enable all logging"`
}

func init() {
	arg.Register(&guiArg)
}

func GetArg(a string) bool {
	return guiArg.GuiDebug
}
