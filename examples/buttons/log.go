// This creates a simple hello world window
package main

import 	(
	"fmt"
	arg "github.com/alexflint/go-arg"
	"go.wit.com/gui"
	log "go.wit.com/gui/log"
)


var args struct {
	Foo string
	Bar bool
	User string `arg:"env:USER"`
	Demo bool `help:"run a demo"`
	gui.GuiArgs
	log.LogArgs
}

/*
var f1 *os.File
var f2 *os.File
var err error
*/

func init() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar, args.User)

	if (args.Gui != "") {
		gui.GuiArg.Gui = args.Gui
	}
	log.Log(true, "INIT() args.GuiArg.Gui =", gui.GuiArg.Gui)

/*
	log.Println()
	log.Println("STDOUT is now at /tmp/guilogfile")
	log.Println("STDOUT is now at /tmp/guilogfile")
	log.Println()
	f1, err = os.OpenFile(outfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// hmm. is there a trick here or must this be in main()
	// defer f.Close()

	log.SetOutput(f1)
	log.Println("This is a test log entry")
*/
}

/*
func captureSTDOUT() {
	f2, _ = os.OpenFile("/tmp/my.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	multiWriter := io.MultiWriter(os.Stderr, f2)
	rd, wr, err := os.Pipe()
	if err != nil {
	    os.Exit(1)
	}

	// overwrite os.Stdout
	os.Stderr = wr

	go func() {
	    scanner := bufio.NewScanner(rd)
	    for scanner.Scan() {
	        stdoutLine := scanner.Text()
	        multiWriter.Write([]byte(stdoutLine + "\n"))
	    }
	}()

	fmt.Println("foobar")

	// hacky sleep to ensure the go routine can write before program exits
	time.Sleep(time.Second)
}
*/
