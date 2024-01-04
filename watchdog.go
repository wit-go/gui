package gui

import (
	"time"

	"go.wit.com/log"
)

var watchtime time.Duration = 100 // in tenths of seconds

/*
	This program sits here.
	If you exit here, the whole thing will os.Exit()
	TODO: use Ticker

	This goroutine can be used like a watchdog timer
*/
func Watchdog() {
	var i = 1
	for {
		log.Verbose("gui.Watchdog() is alive. give me something to do.", i)
		i += 1
		time.Sleep(watchtime * time.Second / 10)
	}
}
