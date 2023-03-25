package main

import 	(
	"os"
	witlog "git.wit.org/wit/gui/log"
)

func log(a ...any) {
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
