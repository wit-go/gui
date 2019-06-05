package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"
// import "github.com/davecgh/go-spew/spew"

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
		if (Config.Debug) {
			log.Println("gui.defaultButtonClick() Data.AllButtons =", key, foo)
			// spew.Dump(foo)
		}
		if Data.AllButtons[key].B == button {
			log.Println("\tgui.defaultButtonClick() BUTTON MATCHED")
			guiButtonClick(Data.AllButtons[key])
			return
		}
	}
	log.Println("\tgui.defaultButtonClick() ERROR: BUTTON NOT FOUND")
	if (Config.Debug) {
		panic("gui.defaultButtonClick() ERROR: UNMAPPED ui.Button")
	}
}

func guiButtonClick(button *GuiButton) {
	log.Println("\tgui.guiButtonClick() button.Name =", button.Name)
	if button.Custom != nil {
		log.Println("\tgui.guiButtonClick() DOING CUSTOM FUNCTION")
		button.Custom(button)
		return
	}
	if (Data.MouseClick != nil) {
		Data.MouseClick(button)
	} else {
		log.Println("\tgui.guiButtonClick() IGNORING BUTTON. MouseClick() is nil")
	}
}

func CreateButton(box *GuiBox, custom func(*GuiButton), name string, values interface {}) *GuiButton {
	newUiB := ui.NewButton(name)
	newUiB.OnClicked(defaultButtonClick)

	var newB *GuiButton
	newB		= new(GuiButton)
	newB.B		= newUiB
	if (box.Window == nil) {
		log.Println("CreateButton() box.Window == nil")
		// ErrorWindow(box.Window, "Login Failed", msg) // can't even do this
		panic("maybe print an error and return nil? or make a fake button?")
	}
	newB.Box	= box
	newB.Custom	= custom
	newB.Values	= values

	Data.AllButtons	= append(Data.AllButtons, newB)

	box.UiBox.Append(newB.B, false)
	return newB
}

func CreateFontButton(box *GuiBox, action string) *GuiButton {
        // create a 'fake' button entry for the mouse clicks
	var newGB	GuiButton
	newGB.Name	= "FONT"
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

func CreateColorButton(box *GuiBox, custom func(*GuiButton), name string, values interface {}) *GuiButton {
        // create a 'fake' button entry for the mouse clicks
	var newCB	GuiButton
	newCB.Name	= name
	newCB.CB	= ui.NewColorButton()
	newCB.Box	= box
	newCB.Custom	= custom
	newCB.Values	= values

	Data.AllButtons	= append(Data.AllButtons, &newCB)

	newCB.CB.OnChanged(func (*ui.ColorButton) {
		log.Println("ColorButton.OnChanged() START Color Button Click")
		r, g, b, a := newCB.CB.Color()
		log.Println("ColorButton.OnChanged() Color() =", r, g, b, a)
		if (newCB.Custom != nil) {
			newCB.Custom(&newCB)
		} else if (Data.MouseClick != nil) {
			Data.MouseClick(&newCB)
		}
	})
	box.UiBox.Append(newCB.CB, false)
	return &newCB
}
