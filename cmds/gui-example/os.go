package main

import "log"
import "strings"
import "os"
import "os/exec"
import "io/ioutil"
import "errors"
// import "bufio"

// import "github.com/davecgh/go-spew/spew"

/*
import "time"
import "runtime"
import "runtime/debug"
import "runtime/pprof"

import "git.wit.org/wit/gui"
import "git.wit.org/wit/shell"
import "github.com/gobuffalo/packr"
*/

func runSimpleCommand(s string) {
	cmd := strings.TrimSpace(s) // this is like 'chomp' in perl
	cmd  = strings.TrimSuffix(cmd, "\n") // this is like 'chomp' in perl
        cmdArgs := strings.Fields(cmd)
	runLinuxCommand(cmdArgs)
}

var geom string = "120x30+500+500"

func xterm(cmd string) {
	var tmp []string
	var argsXterm = []string{"nohup", "xterm", "-geometry", geom}
	tmp = append(argsXterm, "-hold", "-e", cmd)
	log.Println("xterm cmd=", cmd)
	go runCommand(tmp)
}

func runCommand(cmdArgs []string) {
	log.Println("runCommand() START", cmdArgs)
	process := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	// process := exec.Command("xterm", "-e", "ping localhost")
	log.Println("runCommand() process.Start()")
	process.Start()
	log.Println("runCommand() process.Wait()")
	err := process.Wait()
	lookupError(err)
	log.Println("runCommand() NEED TO CHECK THE TIME HERE TO SEE IF THIS WORKED")
	log.Println("runCommand() OTHERWISE INFORM THE USER")
	log.Println("runCommand() END", cmdArgs)
}

func lookupError(err error) {
	var (
		ee *exec.ExitError
		pe *os.PathError
	)

	if errors.As(err, &ee) {
		log.Println("ran, but non-zero exit code =", ee.ExitCode()) // ran, but non-zero exit code
	} else if errors.As(err, &pe) {
		log.Printf("os.PathError = %v", pe) // "no such file ...", "permission denied" etc.
	} else if err != nil {
		log.Printf("something really bad happened general err = %v", err) // something really bad happened!
		if exitError, ok := err.(*exec.ExitError); ok {
			log.Printf("exitError.ExitCode() is %d\n", exitError.ExitCode())
		}
	} else {
		log.Println("success! // ran without error (exit code zero)")
	}
}

func runLinuxCommand(cmdArgs []string) (string, error) {
	process := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)

	process.Stdin = os.Stdin
	process.Stderr = os.Stderr

	stdOut, err := process.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := process.Start(); err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(stdOut)
	if err != nil {
		return "", err
	}
	err = process.Wait()
	lookupError(err)

	log.Println(string(bytes))
	return string(bytes), err
}
