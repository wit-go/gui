package witlog

import (
)

//
// Attempt to switch logging to syslog on linux
//

// This struct can be used with the go-arg package
type LogArgs struct {
	Log []string `arg:"--log" help:"Where to log [syslog,stdout]"`
}
