package gui

import (
	"log"
	"time"
)

var watchtime time.Duration = 100 // in tenths of seconds

/*
	This program sits here.
	If you exit here, the whole thing will os.Exit()

	This goroutine can be used like a watchdog timer
*/
func Watchdog() {
	var i = 1
	for {
		log.Println("watchdog timer is alive. give me something to do.", i, "debug =", Config.Debug.Debug)
		i += 1
		time.Sleep(watchtime * time.Second / 10)
	}
}
