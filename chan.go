package gui

// channel communication to the plugins
// https://github.com/sourcegraph/conc
// https://www.reddit.com/r/golang/comments/11x1oek/hello_gophers_show_me_your_concurrent_code/

import (
	// "regexp"
	// "go.wit.com/gui/toolkit"
	"sync"
	"runtime"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/stream"
	"github.com/sourcegraph/conc/panics"
)

// this should never exit
// TODO: clean up all this poorly named code
func makeConc() {
	var wg conc.WaitGroup
	defer wg.Wait()

	startTheThing(&wg)
	log(debugError, "panic?")
	sleep(2)
	log(debugError, "panic? after sleep(5)")
}

func startTheThing(wg *conc.WaitGroup) {
	f := func() {
		log(debugError, "startTheThing() == about to panic now")
		panic("test conc.WaitGroup")
	}
	wg.Go(func() {
		ExampleCatcher(f)
	})
}

func ExampleCatcher(f func()) {
	var pc panics.Catcher
	i := 0
	pc.Try(func() { i += 1 })
	pc.Try(f)
	pc.Try(func() { i += 1 })

	recovered := pc.Recovered()

	log(debugError, "panic.Recovered():", recovered.Value.(string))
	frames := runtime.CallersFrames(recovered.Callers)
	for {
		frame, more := frames.Next()
		log(debugError, "\t", frame.Function)

		if !more {
			break
		}
	}
}

func mapStream(
    in chan int,
    out chan int,
    f func(int) int,
) {
    tasks := make(chan func())
    taskResults := make(chan chan int)

    // Worker goroutines
    var workerWg sync.WaitGroup
    for i := 0; i < 10; i++ {
        workerWg.Add(1)
        go func() {
            defer workerWg.Done()
            for task := range tasks {
                task()
            }
        }()
    }

    // Ordered reader goroutines
    var readerWg sync.WaitGroup
    readerWg.Add(1)
    go func() {
        defer readerWg.Done()
        for result := range taskResults {
            item := <-result
            out <- item
        }
    }()

    // Feed the workers with tasks
    for elem := range in {
        resultCh := make(chan int, 1)
        taskResults <- resultCh
        tasks <- func() {
            resultCh <- f(elem)
        }
    }

    // We've exhausted input.
    // Wait for everything to finish
    close(tasks)
    workerWg.Wait()
    close(taskResults)
    readerWg.Wait()
}

func mapStream2(
    in chan int,
    out chan int,
    f func(int) int,
) {
    s := stream.New().WithMaxGoroutines(10)
    for elem := range in {
        elem := elem
        s.Go(func() stream.Callback {
            res := f(elem)
            return func() { out <- res }
        })
    }
    s.Wait()
}
