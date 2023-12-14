package main

import (
	"math/rand"
	"github.com/awesome-gocui/gocui"
)

//w.v.SelBgColor = gocui.ColorCyan
//color.go:       w.v.SelFgColor = gocui.ColorBlack
//color.go:               w.v.BgColor = gocui.ColorGreen

type colorT struct {
	frame gocui.Attribute
	fg gocui.Attribute
	bg gocui.Attribute
	selFg gocui.Attribute
	selBg gocui.Attribute
	name string
}

var none gocui.Attribute = gocui.AttrNone
var lightPurple gocui.Attribute = gocui.GetColor("#DDDDDD") // light purple
var darkPurple gocui.Attribute = gocui.GetColor("#FFAA55")  // Dark Purple
var heavyPurple gocui.Attribute = gocui.GetColor("#88AA55") // heavy purple
var powdererBlue gocui.Attribute = gocui.GetColor("#B0E0E6") // w3c 'powerder blue'
var superLightGrey gocui.Attribute = gocui.GetColor("#55AAFF") // super light grey

// Standard defined colors from gocui:
// ColorBlack ColorRed ColorGreen ColorYellow ColorBlue ColorMagenta ColorCyan ColorWhite

// v.BgColor = gocui.GetColor("#111111") // crazy red
// v.BgColor = gocui.GetColor("#FF9911") // heavy red
// v.SelBgColor = gocui.GetColor("#FFEE11") // blood red

// v.BgColor = gocui.GetColor("#55AAFF") // super light grey
// v.BgColor = gocui.GetColor("#FFC0CB") // 'w3c pink' yellow

//                                                                 Normal Text                             On mouseover
//                                     Widget Frame         Text              background              Text           background
var colorWindow  colorT       = colorT{ none         ,    gocui.ColorBlue,   none            ,   none           ,   powdererBlue   ,  "normal window"}
var colorActiveW colorT       = colorT{ none         ,    none           ,   powdererBlue    ,   none           ,   powdererBlue   ,  "active window"}

var colorTab     colorT       = colorT{gocui.ColorBlue,   gocui.ColorBlue,   none            ,   none          ,    powdererBlue   ,  "normal tab"}
var colorActiveT colorT       = colorT{gocui.ColorBlue,   none           ,   powdererBlue    ,   none          ,    powdererBlue   ,  "active tab"}

var colorButton  colorT       = colorT{gocui.ColorGreen,  none          ,    gocui.ColorWhite,   gocui.ColorGreen,  gocui.ColorBlack, "normal button"}
var colorLabel   colorT       = colorT{ none           ,  none          ,    superLightGrey  ,   none            ,  superLightGrey  , "normal label"}
var colorGroup   colorT       = colorT{ none           ,  none          ,    superLightGrey  ,   none            ,  superLightGrey  , "normal group"}

// widget debugging colors. these widgets aren't displayed unless you are debugging
var colorRoot    colorT       = colorT{gocui.ColorRed ,   none          ,    powdererBlue  ,     none          ,    gocui.ColorBlue,  "debug root"}
var colorFlag    colorT       = colorT{gocui.ColorRed ,   none          ,    powdererBlue  ,     none          ,    gocui.ColorGreen, "debug flag"}
var colorBox     colorT       = colorT{gocui.ColorRed ,   none          ,    lightPurple   ,     none          ,    gocui.ColorCyan,  "debug box"}
var colorGrid    colorT       = colorT{gocui.ColorRed ,   none          ,    lightPurple   ,     none          ,    gocui.ColorRed,   "debug grid"}
var colorNone    colorT       = colorT{ none          ,   none          ,    none          ,     none          ,    none          ,   "debug none"}

// actually sets the colors for the gocui element 
// the user will see the colors change when this runs
// TODO: add black/white only flag for ttyS0 
// TODO: or fix kvm/qemu serial console & SIGWINCH.
// TODO: and minicom and uboot and 5 million other things.
// TODO: maybe enough of us could actually do that if we made it a goal.
// TODO: start with riscv boards and fix it universally there
// TODO: so just a small little 'todo' item here
func (n *node) setColor(newColor *colorT) {
	tk := n.tk
	if (tk.color == newColor) {
		// nothing to do since the colors have nto changed
		return
	}
	tk.color = newColor
	if (tk.v == nil) {
		return
	}
	if (tk.color == nil) {
		log(true, "Set the node to color = nil")
		tk.color = &colorNone
	}
	log(true, "Set the node to color =", tk.color.name)
	n.recreateView()
}

func (n *node) setDefaultWidgetColor() {
	n.showView()
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
