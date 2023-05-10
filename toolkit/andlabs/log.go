package main

import 	(
	witlog "git.wit.org/wit/gui/log"
)

// various debugging flags
var logNow bool = true	// useful for active development
var logError bool = true
var logWarn bool = true
var logInfo bool = false
var logVerbose bool = false

func log(a ...any) {
	witlog.Where = "wit/gui/andlabs"
	witlog.Log(a...)
}

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	witlog.Exit(a...)
}
