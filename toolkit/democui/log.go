package main

import 	(
	"os"
	witlog "git.wit.org/wit/gui/log"
)

// various debugging flags
var logNow bool = true	// useful for active development
var logError bool = true
var logWarn bool = false
var logInfo bool = false
var logVerbose bool = false

func log(a ...any) {
	witlog.Where = "wit/democui"
	witlog.Log(a...)
}

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	witlog.Exit(a...)
}

func setOutput(f *os.File) {
	witlog.SetOutput(f)
}
