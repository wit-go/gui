package witlog
//
// version v1.2
//
// I like things to be easy.
//
// this means all the log settings are in one place. it should allow
// things to be over-ridden externally to the library
// but still allow command line --args to pass debugging settings
//
// I also have a generic sleep() and exit() in here because it's simple
//
// Usage:
//
// log("something", foo, bar)
// var DEBUG bool = true
// log(DEBUG, "something else", someOtherVariable)  # if DEBUG == false, return doing nothing
// log(SPEW, "something else", someOtherVariable)   # this get's sent to spew.Dump(). Very useful for debugging!
//

import 	(
	"os"
	"runtime"
	"runtime/pprof"
	golog "log"
	"time"
	"reflect"
	"github.com/davecgh/go-spew/spew"
	// "net"
)

// various debugging flags
var LogNow bool = true	// useful for active development
var LogError bool = true
var LogWarn bool = false
var LogInfo bool = false
var LogVerbose bool = false
var LogSleep bool = false

var LOGOFF bool = false // turn this off, all logging stops
var debugToolkit bool = false // does spew stuff?

var Where string = "gui/log"

type Spewt struct {
	a bool
}

var SPEW Spewt

/*
	sleep()		# you know what this does? sleeps for 1 second. yep. dump. easy.
	sleep(.1)	# you know what this does? yes, it sleeps for 1/10th of a second
*/
func Sleep(a ...any) {
	if (a == nil) {
		time.Sleep(time.Second)
		return
	}

	Log(LogSleep, "sleep", a[0])

	switch a[0].(type) {
	case int:
		time.Sleep(time.Duration(a[0].(int)) * time.Second)
	case float64:
		time.Sleep(time.Duration(a[0].(float64) * 1000) * time.Millisecond)
	default:
		Log("sleep a[0], type = ", a[0], reflect.TypeOf(a[0]))
	}
}

/*
	exit()		# yep. exits. I guess everything must be fine
	exit(3)		# I guess 3 it is then
	exit("dont like apples")	# ok. I'll make a note of that
*/
func Exit(a ...any) {
	Log(LogError, "exit", a)
	//if (a) {
	//	os.Exit(a)
	//}
	os.Exit(0)
}

/*
	I've spent, am spending, too much time thinking about 'logging'. 'log', 'logrus', 'zap', whatever.
	I'm not twitter. i don't give a fuck about how many nanoseconds it takes to log. Anyway, this
	implementation is probably faster than all of those because you just set one bool to FALSE
	and it all stops.
	Sometimes I need to capture to stdout, sometimes stdout can't
	work because it doesn't exist for the user. This whole thing is a PITA. Then it's spread
	over 8 million references in every .go file. I'm tapping out and putting
	it in one place. here it is. Also, this makes having debug levels really fucking easy.
	You can define whatever level of logging you want from anywhere (command line) etc.

	log()		# doesn't do anything
	log(stuff)	# sends it to whatever log you define in a single place. here is the place
*/

func Log(a ...any) {
	if (LOGOFF) {
		return
	}

	if (a == nil) {
		return
	}
	var tbool bool
	if (reflect.TypeOf(a[0]) == reflect.TypeOf(tbool)) {
		if (a[0] == false) {
			return
		}
		a[0] = Where
	}

	if (reflect.TypeOf(a[0]) == reflect.TypeOf(SPEW)) {
		// a = a[1:]
		a[0] = Where
		if (debugToolkit) {
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(a)
			// spew.Dump(a)
		}
		return
	}

	golog.Println(a...)
}

func loggo() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	golog.Println("runtime.NumGoroutine() = ", runtime.NumGoroutine())
}

func logindent(depth int, format string, a ...interface{}) {
	var tabs string
	for i := 0; i < depth; i++ {
		tabs = tabs + format
	}

	// newFormat := tabs + strconv.Itoa(depth) + " " + format
	newFormat := tabs + format
	Log(debugToolkit, newFormat, a)
}

func SetOutput(f *os.File) {
	golog.SetOutput(f)
}
