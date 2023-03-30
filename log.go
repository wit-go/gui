package gui

import 	(
	witlog "git.wit.org/wit/gui/log"
)

// various debugging flags
var logNow bool = true	// useful for active development
var logError bool = true
var logWarn bool = false
var logInfo bool = false
var logVerbose bool = false

// var log interface{}

func log(a ...any) {
	witlog.Where = "wit/gui"
	witlog.Log(a...)
}

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	witlog.Exit(a...)
}

// b bool, print if true
func logindent(b bool, depth int, format string, a ...any) {
	var tabs string
	for i := 0; i < depth; i++ {
		tabs = tabs + format
	}

	// newFormat := tabs + strconv.Itoa(depth) + " " + format
	newFormat := tabs + format

	// array prepend(). Why isn't this a standard function. It should be:
	// a.prepend(debugGui, newFormat)
	a = append([]any{b, newFormat}, a...)
	witlog.Log(a...)
}
