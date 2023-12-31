// https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go
// who came up with the idea of making community tutorials. that was a good idea!

package gui

import (
	"fmt"
	"sync"
)

var debugWG *sync.WaitGroup
var debugNumberChan chan int

func (n *Node) DebugGoChannels() {
	var w, g *Node

	w = n.NewWindow("Debug GO Channels")
	w.Custom = w.StandardClose

	g = w.NewGroup("Channel stuff")

	// var debugWG sync.WaitGroup
	g.NewButton("init()", func () {
		if (debugNumberChan == nil) {
			log("making debugNumberChan channel")
			debugNumberChan = make(chan int)
		} else {
			log("debugNumberChan already made")
		}
		debugWG = new(sync.WaitGroup)
	})
	g.NewButton("go printInt(x) (read x values off the channel)", func () {
			debugWG.Add(1)
			go printInt(2, "routine1")
			debugWG.Add(1)
			go printInt(2, "routine2")
	})
	g.NewButton("sendNumber(2) (chan <- 2, 4)", func () {
		debugWG.Add(1)
		debugWG.Add(1)
		go sendNumber(2)
		go sendNumber(4)
	})
	g.NewButton("sendNumber(1) (chan <- 7)", func () {
		debugWG.Add(1)
		go sendNumber(7)
	})
	g.NewButton("send 4 numbers (chan <- int)", func () {
		log("generateNumbers(4)")
		go generateNumbers(4)
	})
	g.NewButton("debugWG.Done()", func () {
		log("ran debugWG.Done()")
		debugWG.Done()
	})
	g.NewButton("close chan", func () {
		log("close() on", debugNumberChan)
		close(debugNumberChan)
	})
	g.NewButton("print", func () {
		log("waitgroup counter is ?")
	})
}
func sendNumber(i int) {
	log("START debugNumberChan <-", i, "  (sending", i, "to channel)")
	debugNumberChan <- i
	debugWG.Wait()
	log("END   debugNumberChan sendNumber() done", i)
}

func generateNumbers(total int) {
	fmt.Printf("START generateNumbers()\n")
	for idx := 1; idx <= total; idx++ {
		log("ran debugNumberChan <= idx where idx =", idx)
		fmt.Printf("S generateNumbers() sending %d to channel\n", idx)
		debugNumberChan <- idx
		// res, err := (<-r)()
		fmt.Printf("E generateNumbers() sending %d to channel\n", idx)
	}
	debugWG.Wait()
	fmt.Printf("END   generateNumbers()\n")
}

// i equals the number of times to read values from the channel
func printInt(i int, name string) {
	tmp := 1
	log("START printInt", name, "read debugNumberChan()")
	for num := range debugNumberChan {
		log("printInt()",name, "read", num, "from channel")
		debugWG.Done()
		if (tmp == i) {
			return
		}
		tmp += 1
	}
	fmt.Printf("END printInt()", name, "read debugNumberChan\n")
}
