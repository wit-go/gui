package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
)

type greeting string


// func main() {
func (g greeting) Greet() {
	fmt.Println("Hello Universe")
	Init()
	// ToolkitMain()
}

// this is exported
var Greeter greeting

func AddGroup(name string) {
	log.Println("addGroup()", name)
	currentY = 2
	currentX += groupSize + 6
}

func AddButton(name string) error {
	t := len(name)
	v, err := baseGui.SetView(name, currentX, currentY, currentX+t+3, currentY+2, 0)
	if err == nil {
		return err
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		return err
	}

	v.Wrap = true
	fmt.Fprintln(v, " " + name)
	fmt.Fprintln(v, strings.Repeat("foo\n", 2))

	if _, err := baseGui.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	currentY += 3
	if (groupSize < len(views)) {
		groupSize = len(views)
	}
	return nil
}
