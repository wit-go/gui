package gui

import (
	arg "github.com/alexflint/go-arg"

	"go.wit.com/log"
)

var GUI log.LogFlag
var NODE log.LogFlag
var PLUG log.LogFlag
var INFO log.LogFlag

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

	log.Register("gui", "debugGui", &debugGui)

	INFO.B = false
	INFO.Name = "INFO"
	INFO.Subsystem = "gui"
	INFO.Desc = "Enable log.Info()"
	INFO.Register()

	GUI.B = false
	GUI.Name = "GUI"
	GUI.Subsystem = "gui"
	GUI.Desc = "basic GUI debugging"
	GUI.Register()

	NODE.B = false
	NODE.Name = "NODE"
	NODE.Subsystem = "gui"
	NODE.Desc = "basic NODE debugging"
	NODE.Register()

	PLUG.B = false
	PLUG.Name = "PLUG"
	PLUG.Subsystem = "gui"
	PLUG.Desc = "basic PLUG debugging"
	PLUG.Register()

	for _, s := range log.ListFlags() {
		log.Info("go.wit.com/gui/gui ListFlags() returned:", s)
	}
}
