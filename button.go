package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

// THIS IS CLEAN

// This is the default mouse click handler
// Every mouse click that hasn't been assigned to
// something specific will fall into this routine
// By default, all it runs is the call back to
// the main program that is using this library
//
// This routine MUST be here as this is how the andlabs/ui works
// This is the raw routine passed to every button in andlabs libui / ui
//
// There is a []GuiButton which has all the buttons. We search
// for the button and then call the function below
//
func defaultButtonClick(button *ui.Button) {
	log.Println("gui.defaultButtonClick() LOOK FOR BUTTON button =", button)
	for key, foo := range Data.AllButtons {
		if (Data.Debug) {
			log.Println("gui.defaultButtonClick() Data.AllButtons =", key, foo)
			// spew.Dump(foo)
		}
		if Data.AllButtons[key].B == button {
			log.Println("\tgui.defaultButtonClick() BUTTON MATCHED")
			log.Println("\tgui.defaultButtonClick() Data.AllButtons[key].Action =", Data.AllButtons[key].Action)
			if Data.AllButtons[key].custom != nil {
				log.Println("\tgui.defaultButtonClick() DOING CUSTOM FUNCTION")
				Data.AllButtons[key].custom(Data.AllButtons[key])
				return
			}
			if (Data.MouseClick != nil) {
				Data.MouseClick(Data.AllButtons[key])
			} else {
				log.Println("\tgui.defaultButtonClick() IGNORING BUTTON. MouseClick() is nil")
			}
			return
		}
	}
	log.Println("\tgui.defaultButtonClick() BUTTON NOT FOUND")
	if (Data.Debug) {
		panic("gui.defaultButtonClick() SHOULD NOT HAVE UNMAPPED BUTTONS")
	}
}

func CreateButton(box *GuiBox, a *pb.Account, vm *pb.Event_VM, name string, action string, custom func(*GuiButton)) *GuiButton {
	newUiB := ui.NewButton(name)
	newUiB.OnClicked(defaultButtonClick)

	var newB *GuiButton
	newB		= new(GuiButton)
	newB.B		= newUiB
	if (box.Window == nil) {
		log.Println("CreateButton() box.Window == nil")
		panic("crap")
	}
	newB.GW		= box.Window
	newB.Account	= a
	newB.VM		= vm
	newB.Box	= box
	newB.Action	= action
	newB.custom	= custom
	Data.AllButtons	= append(Data.AllButtons, newB)

	box.UiBox.Append(newB.B, false)
	return newB
}

func CreateFontButton(box *GuiBox, action string) *GuiButton {

        // create a 'fake' button entry for the mouse clicks
	var newGB	GuiButton
	newGB.Action	= action
	newGB.FB	= ui.NewFontButton()
	newGB.Box	= box
	Data.AllButtons	= append(Data.AllButtons, &newGB)

	newGB.FB.OnChanged(func (*ui.FontButton) {
		log.Println("FontButton.OnChanged() START mouseClick(&newBM)", newGB)
		if (Data.MouseClick != nil) {
			Data.MouseClick(&newGB)
		}
	})
	return &newGB
}
