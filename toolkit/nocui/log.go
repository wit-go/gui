package main

import 	(
	"io"
//	"fmt"
//	"strings"
	witlog "git.wit.org/wit/gui/log"
)

// various debugging flags
var logNow bool = true	// useful for active development
var logError bool = true
var logWarn bool = false
var logInfo bool = false
var logVerbose bool = false

var outputS []string

func log(a ...any) {
	witlog.Where = "wit/gocui"
	witlog.Log(a...)
}

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	witlog.Exit(a...)
}

/*
func newLog(a ...any) {
	s := fmt.Sprint(a...)
	tmp := strings.Split(s, "\n")
	outputS = append(outputS, tmp...)
	if (len(outputS) > 50) {
		outputS = outputS[10:]
	}
	if (me.baseGui != nil) {
		v, _ := me.baseGui.View("msg")
		if (v != nil) {
			v.Clear()
			fmt.Fprintln(v, strings.Join(outputS, "\n"))
		}
	}
}
*/

func setOutput(w io.Writer) {
	if (w == nil) {
		return
	}
	witlog.SetOutput(w)
	// witlog.SetToolkitOutput(newLog)
}
