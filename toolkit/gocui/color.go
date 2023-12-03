package main

import (
	"math/rand"
	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// ColorBlack ColorRed ColorGreen ColorYellow ColorBlue ColorMagenta ColorCyan ColorWhite
// gocui.GetColor("#FFAA55")  // Dark Purple
func (n *node) setDefaultWidgetColor() {
	w := n.tk
	log(logInfo, "setDefaultWidgetColor() on", n.WidgetType, n.Name)
	v, _ := me.baseGui.View(w.cuiName)
	if (v == nil) {
		log(logError, "setDefaultWidgetColor() failed on view == nil")
		return
	}
	sleep(.05)
	// v.BgColor = gocui.GetColor("#FFAA55")  // Dark Purple
	// v.BgColor = gocui.GetColor("#88AA55") // heavy purple
	// v.BgColor = gocui.GetColor("#111111") // crazy red
	// v.BgColor = gocui.GetColor("#FF9911") // heavy red
	// v.SelBgColor = gocui.GetColor("#FFEE11") // blood red

	// v.BgColor = gocui.GetColor("#55AAFF") // super light grey
	// v.BgColor = gocui.GetColor("#FFC0CB") // 'w3c pink' yellow
	switch n.WidgetType {
	case toolkit.Root:
		v.FrameColor = gocui.ColorRed
		v.BgColor = gocui.GetColor("#B0E0E6") // w3c 'powerder blue'
	case toolkit.Flag:
		v.FrameColor = gocui.ColorRed
		v.BgColor = gocui.GetColor("#B0E0E6") // w3c 'powerder blue'
	case toolkit.Window:
		v.FgColor = gocui.ColorCyan
		v.SelBgColor = gocui.ColorBlue
		v.FrameColor = gocui.ColorBlue
	case toolkit.Tab:
		v.SelBgColor = gocui.ColorBlue
		v.FrameColor = gocui.ColorBlue
	case toolkit.Button:
		v.BgColor = gocui.ColorWhite
		v.FrameColor = gocui.ColorGreen
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
	case toolkit.Label:
		v.BgColor = gocui.GetColor("#55AAFF") // super light grey
		v.SelBgColor = gocui.GetColor("#55AAFF") // super light grey
	case toolkit.Box:
		v.FrameColor = gocui.ColorRed
		// v.BgColor = gocui.GetColor("#FFC0CB") // 'w3c pink' yellow
		v.BgColor = gocui.GetColor("#DDDDDD") // light purple
	case toolkit.Grid:
		// v.FgColor = gocui.ColorCyan
		// v.SelBgColor = gocui.ColorBlue
		// v.FrameColor = gocui.ColorBlue
	case toolkit.Group:
		v.BgColor = gocui.GetColor("#55AAFF") // super light grey
	default:
	}
}

// SetColor("#FFAA55") // purple
func (w *cuiWidget) SetColor(c string) {
	if (w.v == nil) {
		log(logError, "SetColor() failed on view == nil")
		return
	}
	w.v.SelBgColor = gocui.ColorCyan
	w.v.SelFgColor = gocui.ColorBlack
	switch c {
	case "Green":
		w.v.BgColor = gocui.ColorGreen
	case "Purple":
		w.v.BgColor = gocui.GetColor("#FFAA55")
	case "Yellow":
		w.v.BgColor = gocui.ColorYellow
	case "Blue":
		w.v.BgColor = gocui.ColorBlue
	case "Red":
		w.v.BgColor = gocui.ColorRed
	default:
		w.v.BgColor = gocui.GetColor(c)
	}
}

func (n *node) setDefaultHighlight() {
	w := n.tk
	if (w.v == nil) {
		log(logError, "SetColor() failed on view == nil")
		return
	}
	w.v.SelBgColor = gocui.ColorGreen
	w.v.SelFgColor = gocui.ColorBlack
}

func randColor() gocui.Attribute {
	colors := []string{"Green", "#FFAA55", "Yellow", "Blue", "Red", "Black", "White"}
	i := rand.Intn(len(colors))
	log("randColor() i =", i)
	return gocui.GetColor(colors[i])
}

func (n *node) redoColor(draw bool) {
	w := n.tk
	if (w == nil) {
		return
	}

	sleep(.05)
	n.setDefaultHighlight()
	n.setDefaultWidgetColor()

	for _, child := range n.children {
		child.redoColor(draw)
	}
}
