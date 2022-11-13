package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

func NewButton(parentW *toolkit.Widget, w *toolkit.Widget) {
	log.Println("gui.gocui.AddButton()", w.Name)
	addButton(w.Name)
	// viewWidget[v] = w
	stringWidget[w.Name] = w
	listMap()
}

func addButton(name string) *gocui.View {
	t := len(name)
	if (baseGui == nil) {
		panic("WTF")
	}
	v, err := baseGui.SetView(name, currentX, currentY, currentX+t+3, currentY+2, 0)
	if err == nil {
		log.Println("wit/gui internal plugin error", err)
		return nil
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		log.Println("wit/gui internal plugin error", err)
		return nil
	}

	v.Wrap = true
	fmt.Fprintln(v, " " + name)
	fmt.Fprintln(v, strings.Repeat("foo\n", 2))

	currentView, err := baseGui.SetCurrentView(name)
	if err != nil {
		log.Println("wit/gui internal plugin error", err)
		return nil
	}
	log.Println("wit/gui addbutton() current view name =", currentView.Name())

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	currentY += 3
	if (groupSize < len(name)) {
		groupSize = len(name)
	}
	return currentView
}
