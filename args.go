package gui

import (
	"go.wit.com/arg"
	"go.wit.com/log"
)

var INFO log.LogFlag
var NOW log.LogFlag
var GUI log.LogFlag
var NODE log.LogFlag
var PLUG log.LogFlag
var CHANGE log.LogFlag

var argGui ArgsGui

// This struct can be used with the go-arg package
type ArgsGui struct {
	GuiPlugin string `arg:"--gui" help:"Use this gui toolkit [andlabs,gocui,nocui]"`
	GuiVerbose bool `arg:"--gui-verbose" help:"enable all logging"`
}

// returns the toolkit
func ArgToolkit() string {
	return argGui.GuiPlugin
}

func init() {
	arg.Register(&argGui)

	full := "go.wit.com/gui/gui"
	short := "gui"

	NOW.NewFlag("NOW", true,  full, short, "temp debugging stuff")
	INFO.NewFlag("INFO", false, full, short, "General Info")
	GUI.NewFlag("GUI", false, full, short, "basic GUI internals")
	NODE.NewFlag("NODE", false, full, short, "binary tree debugging")
	PLUG.NewFlag("PLUG", false, full, short, "basic PLUG debuggging")
	CHANGE.NewFlag("CHANGE", false, full, short, "user changed something")
}
