package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
)

func (w *Widget) AddButton() {
// func (g greeting) AddButton() {
	log.Println("gui.gocui.AddButton()", w.Name)
	addButton2(w.Name, w.Event)
}

func addButton2(name string, e func(*Widget) *Widget) {
	addButton(name)
}

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
