package main

import (
	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

func setupWidgetT(a *toolkit.Action) *cuiWidget {
	var w *cuiWidget
	w = new(cuiWidget)

	w.name = a.Name
	w.text = a.Text

	w.widgetType = a.WidgetType
	w.id = a.WidgetId
	if (w.id > me.highest) {
		me.highest = w.id
	}
	w.parentId = a.ParentId
	me.widgets[w.id] = w

	// w.showWidgetPlacement(logNow)
	return w
}

// ColorBlack ColorRed ColorGreen ColorYellow ColorBlue ColorMagenta ColorCyan ColorWhite
// gocui.GetColor("#FFAA55")  // Dark Purple
func (w *cuiWidget) SetDefaultWidgetColor() {
	log(logInfo, "SetDefaultWidgetColor() on", w.widgetType, w.name)
	if (w.v == nil) {
		log(logError, "SetDefaultWidgetColor() failed on view == nil")
		return
	}
	w.SetDefaultHighlight()
	switch w.widgetType {
	case toolkit.Button:
		w.v.BgColor = gocui.ColorGreen
		w.v.FrameColor = gocui.ColorGreen
	case toolkit.Checkbox:
		w.v.BgColor = gocui.GetColor("#FFAA55")  // Dark Purple
		w.v.FrameColor = gocui.GetColor("#FFEE11")
	case toolkit.Dropdown:
		w.v.BgColor = gocui.ColorCyan
		w.v.FrameColor = gocui.ColorGreen
	case toolkit.Textbox:
		w.v.BgColor = gocui.ColorYellow
		w.v.FrameColor = gocui.ColorGreen
	case toolkit.Slider:
		w.v.BgColor = gocui.GetColor("#FFAA55")  // Dark Purple
		w.v.FrameColor = gocui.ColorRed
	case toolkit.Label:
		w.v.FrameColor = gocui.ColorRed
	default:
		w.v.BgColor = gocui.ColorYellow
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

func (w *cuiWidget) SetDefaultHighlight() {
	if (w.v == nil) {
		log(logError, "SetColor() failed on view == nil")
		return
	}
	w.v.SelBgColor = gocui.ColorGreen
	w.v.SelFgColor = gocui.ColorBlack
}
