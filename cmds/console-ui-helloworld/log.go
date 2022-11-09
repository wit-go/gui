// This creates a simple hello world window
package main

import 	(
	"log"
	"fmt"
	"os"
	arg "github.com/alexflint/go-arg"
)


var args struct {
	Foo string
	Bar bool
	User string `arg:"env:USER"`
	Demo bool `help:"run a demo"`
}

var f *os.File
var err error

func init() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar, args.User)

	f, err = os.OpenFile("/tmp/guilogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// hmm. is there a trick here or must this be in main()
	// defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
}
