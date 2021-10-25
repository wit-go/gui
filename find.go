package gui

import (
	"log"

//	"github.com/andlabs/ui"
//	_ "github.com/andlabs/ui/winmanifest"
//	"github.com/davecgh/go-spew/spew"
)


func FindWindow(s string) *GuiWindow {
	for name, window := range Data.WindowMap {
		if name == s {
			return window
		}
	}
	log.Printf("COULD NOT FIND WINDOW", s)
	return nil
}

func FindBox(s string) *GuiBox {
	for name, window := range Data.WindowMap {
		if name != s {
			continue
		}
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			return abox
		}
		log.Println("gui.FindBox() NEED TO INIT WINDOW name =", name)
	}
	log.Println("gui.FindBox() COULD NOT FIND BOX", s)
	return nil
}
