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
		log(logInfo, "watchdog timer is alive. give me something to do.", i)
		if (Config.rootNode == nil) {
			log(logInfo, "Config.rootNode == nil", i)
		} else {
			if (logVerbose) {
				Config.rootNode.ListChildren(true)
			}
		}
		if (i == 2) {
			Config.rootNode.LoadToolkit("gocui")
		}
//		if (i == 3) {
//			Config.rootNode.LoadToolkit("andlabs")
//		}
		i += 1
		time.Sleep(watchtime * time.Second / 10)
	}
}
// https://www.reddit.com/r/golang/comments/12em87q/how_to_run_periodic_tasks/
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
*/
