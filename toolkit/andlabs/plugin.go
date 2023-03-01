package main

// if you include more than just this import
// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
import "git.wit.org/wit/gui/toolkit"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"

//
// This should be called ?
// Pass() ?
// This handles all interaction between the wit/gui package (what golang knows about)
// and this plugin that talks to the OS and does scary and crazy things to make
// a GUI on whatever OS or whatever GUI toolkit you might have (GTK, QT, WASM, libcurses)
//
// Once you are here, you should be in a protected goroutine created by the golang wit/gui package
//
// TODO: make sure you can't escape this goroutine
//
func Send(p *toolkit.Widget, c *toolkit.Widget) {
	if (p == nil) {
		log(debugPlugin, "Send() parent = nil")
	} else {
		log(debugPlugin, "Send() parent =", p.Name, ",", p.Type)
	}
	log(debugPlugin, "Send() child  =", c.Name, ",", c.Action, ",", c.Type)

	switch c.Type {
	case toolkit.Window:
		newWindow(c)
	case toolkit.Tab:
		newTab(p, c)
	case toolkit.Group:
		newGroup(p, c)
	case toolkit.Button:
		doButton(p, c)
	case toolkit.Checkbox:
		doCheckbox(p, c)
	case toolkit.Label:
		newLabel(p, c)
	case toolkit.Textbox:
		doTextbox(p, c)
	case toolkit.Slider:
		newSlider(p, c)
	case toolkit.Spinner:
		newSpinner(p, c)
	default:
		log(true, "unknown parent =", p.Name, p.Type)
		log(true, "unknown child  =", c.Name, c.Type)
		log(true, "Don't know how to do", c.Type, "yet")
	}
}

// delete the child widget from the parent
// p = parent, c = child
func destroy(p *toolkit.Widget, c *toolkit.Widget) {
	log(true, "delete()", c.Name, c.Type)

	pt := mapToolkits[p]
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "delete FAILED (ct = mapToolkit[c] == nil) for c", c.Name, c.Type)
		// this pukes out a whole universe of shit
		// listMap()
		return
	}

	switch ct.Type {
	case toolkit.Button:
		log(true, "Should delete Button here:", c.Name)
		log(true, "Parent:")
		pt.Dump(true)
		log(true, "Child:")
		ct.Dump(true)
		if (pt.uiBox == nil) {
			log(true, "Don't know how to destroy this")
		} else {
			log(true, "Fuck it, destroy the whole box", pt.Name)
			// pt.uiBox.Destroy() // You have a bug: You cannot destroy a uiControl while it still has a parent.
			pt.uiBox.SetPadded(false)
			pt.uiBox.Delete(4)
			ct.uiButton.Disable()
			// ct.uiButton.Hide()
			ct.uiButton.Destroy()
		}

	case toolkit.Window:
		log(true, "Should delete Window here:", c.Name)
	default:
		log(true, "Don't know how to delete c =", c.Type, c.Name)
		log(true, "Don't know how to delete pt =", pt.Type, pt.Name, pt.uiButton)
		log(true, "Don't know how to delete ct =", ct.Type, ct.Name, ct.uiButton)
		log(true, "Parent:")
		pt.Dump(true)
		log(true, "Child:")
		ct.Dump(true)
		log(true, "Fuckit, let's destroy a button", c.Name, c.Type)
		if (ct.uiButton != nil) {
			pt.uiBox.Delete(4)
			ct.uiButton.Destroy()
		}
	}
}
