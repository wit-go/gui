package gui

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func HorizontalBreak(box *GuiBox) {
	log.Println("VerticalSeparator added to box =", box.Name)
	tmp := ui.NewHorizontalSeparator()
	if (box == nil) {
		return
	}
	if (box.UiBox == nil) {
		return
	}
	box.UiBox.Append(tmp, false)
}

func VerticalBreak(box *GuiBox) {
	log.Println("VerticalSeparator  added to box =", box.Name)
	tmp := ui.NewVerticalSeparator()
	box.UiBox.Append(tmp, false)
}
