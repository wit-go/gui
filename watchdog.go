package gui

import (
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
		log(logNow, "watchdog timer is alive. give me something to do.", i)
		if (Config.rootNode == nil) {
			log(logInfo, "Config.rootNode == nil", i)
		} else {
			if (logVerbose) {
				Config.rootNode.ListChildren(true)
			}
		}
		i += 1
		time.Sleep(watchtime * time.Second / 10)
	}
}
