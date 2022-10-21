package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// make new Group here
func (t Toolkit) NewButton(name string) *Toolkit {
	var newt Toolkit
	var b *ui.Button

	if (t.uiBox == nil) {
		log.Println("gui.ToolboxNode.NewButton() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log.Println("probably could just make a box here?")
		os.Exit(0)
		return nil
	}

	log.Println("gui.Toolbox.NewGroup() create", name)
	b = ui.NewButton(name)
	newt.uiButton = b

	b.OnClicked(func(*ui.Button) {
		log.Println("TODO: IN TOOLKIT GOROUTINE. SHOULD LEAVE HERE VIA channels. button name =", name)
		t.Dump()
		newt.Dump()
		log.Println("wit/gui/toolkit NewButton() Should do something here")
		if (newt.Custom == nil) {
			log.Println("wit/gui/toolkit NewButton() toolkit.Custom == nil")
		} else {
			log.Println("wit/gui/toolkit NewButton() toolkit.Custom() START")
			newt.Custom()
			log.Println("wit/gui/toolkit NewButton() toolkit.Custom() END")
		}
		if (t.Custom == nil) {
			log.Println("wit/gui/toolkit NewButton() parent toolkit.Custom == nil")
		} else {
			log.Println("wit/gui/toolkit NewButton() running parent toolkit.Custom() START (IS THIS A BAD IDEA?)")
			t.Custom()
			log.Println("wit/gui/toolkit NewButton() running parent toolkit.Custom() END   (IS THIS A BAD IDEA?)")
		}
		log.Println("TODO: LEFT TOOLKIT GOROUTINE button name =", name)
	})

	t.uiBox.Append(b, stretchy)

	return &newt
}
