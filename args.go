package gui

import (
	arg "github.com/alexflint/go-arg"

	newlog "go.wit.com/log"
)

var argGui ArgsGui

// This struct can be used with the go-arg package
type ArgsGui struct {
	GuiPlugin string `arg:"--gui" help:"Use this gui toolkit [andlabs,gocui,nocui]"`
	GuiVerbose bool `arg:"--gui-verbose" help:"enable all logging"`
}

func init() {
	arg.Register(&argGui)

	newlog.Register("gui", "debugGui", &debugGui)

	for _, s := range newlog.ListFlags() {
		newlog.Info("go.wit.com/gui/gui ListFlags() returned:", s)
	}
}

// returns the toolkit
func ArgToolkit() string {
	return argGui.GuiPlugin
}
