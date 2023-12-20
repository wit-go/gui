// +build windows

/*
	this is a direct compile version of andlabs/ui for windows since
	golang on windows does not yet support plugins
*/
package gui

import (
	"go.wit.com/gui/toolkit/andlabs-direct"
)

func trythis() {
	log(debugGui, "not sure what to try")
	toolkit.DebugToolkit = true
}
