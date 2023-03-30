package main

import (
	"fmt"
	"errors"
	"strconv"

	"github.com/awesome-gocui/gocui"
	// "git.wit.org/wit/gui/toolkit"
)

func click(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	log(logNow, "click() START", v.Name())
	i, err := strconv.Atoi(v.Name())
	if (err != nil) {
		log(logNow, "click() Can't find widget. error =", err)
	} else {
		log(logNow, "click() Found widget id =", i)
		if (me.widgets[i] != nil) {
			w := me.widgets[i]
			log(logNow, "click() Found widget =", w)
		}
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2, 0); err == nil || errors.Is(err, gocui.ErrUnknownView) {
		v.Clear()
		v.SelBgColor = gocui.ColorCyan
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, l)
	}

	// this seems to delete the button(?)
	// g.SetViewOnBottom(v.Name())
	log(logNow, "click() END")
	return nil
}
