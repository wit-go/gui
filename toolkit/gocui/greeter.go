package main

import (
	"log"

	// "errors"
	// "fmt"
	// "strings"
	// "github.com/awesome-gocui/gocui"
)

type greeting string

// stores the raw toolkit internals
type toolkit struct {
	id     string
	Name   string

	OnChanged	func(toolkit)
}

// this is exported
var Greeter greeting
var Toolkit toolkit

// func main() {
func (g greeting) Greet() {
	log.Println("Hello Universe")
	Init()
	// ToolkitMain()
}

func (g greeting) JcarrButton() {
	log.Println("Hello GreetButton meet Universe")
	addButton("Greet foo")
	addButton("Greet foo 2")
}

func addGroup(name string) {
	log.Println("addGroup()", name)
	currentY = 2
	currentX += groupSize + 6
}

func (g greeting) AddButton(name string) {
// func (g greeting) AddButton() {
	log.Println("gui.gocui.AddButton()", name)
	addButton(name)
}

/*
func addButton(name string) error {
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
*/
