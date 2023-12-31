package gui

import (
	arg "github.com/alexflint/go-arg"

	newlog "go.wit.com/log"
)

var argGui ArgsGui

// This struct can be used with the go-arg package
type ArgsGui struct {
	GuiDebug bool `arg:"--gui-debug" help:"open the GUI debugger"`
	GuiPlugin string `arg:"--gui" help:"Use this gui toolkit [andlabs,gocui,nocui]"`
	GuiVerbose bool `arg:"--gui-verbose" help:"enable all logging"`
}

func init() {
	arg.Register(&argGui)

	newlog.Register("gui", "debugGui", &debugGui)

	for _, s := range newlog.ListAll() {
		newlog.Info("go.wit.org/gui ListAll() returned:", s)
	}
}

// returns the toolkit
func ArgToolkit() string {
	return argGui.GuiPlugin
}

// returns true if --gui-debug was passed from the command line
func ArgDebug() bool {
	return argGui.GuiDebug
}
