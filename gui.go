package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

// THIS IS CLEAN (all that is left is the 'createAddVmBox')

func InitColumns(mh *TableData, parts []TableColumnData) {
	tmpBTindex := 0
	humanID := 0
	for key, foo := range parts {
		log.Println("key, foo =", key, foo)

		parts[key].Index = humanID
		humanID += 1

		if (foo.CellType == "BG") {
			mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableColor{})
			initRowBTcolor        (mh, tmpBTindex, parts[key])
			tmpBTindex += 1
		} else if (foo.CellType == "BUTTON") {
			mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
			initRowButtonColumn   (mh, tmpBTindex,    parts[key].Heading, parts[key])
			tmpBTindex += 1
		} else if (foo.CellType == "TEXTCOLOR") {
			mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
			mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableColor{})
			initRowTextColorColumn(mh, tmpBTindex, tmpBTindex + 1, parts[key].Heading, ui.TableColor{0.0, 0, 0.9, 1}, parts[key])
			tmpBTindex += 2
		} else if (foo.CellType == "TEXT") {
			mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
			initRowTextColumn     (mh, tmpBTindex,    parts[key].Heading, parts[key])
			tmpBTindex += 1
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}
}

func AddTableTab(gw *GuiWindow, name string, rowcount int, parts []TableColumnData, account *pb.Account) *TableData {
	mh := new(TableData)

	mh.RowCount    = rowcount
	mh.Rows        = make([]RowData, mh.RowCount)

	InitColumns(mh, parts)

	model := ui.NewTableModel(mh)
	table := ui.NewTable(
		&ui.TableParams{
			Model:	model,
			RowBackgroundColorModelColumn:	0,	// Row Background color is always index zero
	})

	tmpBTindex := 0
	for key, foo := range parts {
		log.Println(key, foo)
		if (foo.CellType == "BG") {
		} else if (foo.CellType == "BUTTON") {
			tmpBTindex += 1
			table.AppendButtonColumn(foo.Heading, tmpBTindex, ui.TableModelColumnAlwaysEditable)
		} else if (foo.CellType == "TEXTCOLOR") {
			tmpBTindex += 1
			table.AppendTextColumn(foo.Heading, tmpBTindex, ui.TableModelColumnAlwaysEditable,
					&ui.TableTextColumnOptionalParams{
						ColorModelColumn:   tmpBTindex + 1,
			});
			tmpBTindex += 1
		} else if (foo.CellType == "TEXT") {
			tmpBTindex += 1
			table.AppendTextColumn(foo.Heading, tmpBTindex, ui.TableModelColumnAlwaysEditable, nil)
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}

	var gb *GuiBox
	gb = new(GuiBox)

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	gb.UiBox = vbox
	gb.W = gw
	gw.BoxMap[name] = gb
	mh.Box = gb

	vbox.Append(table, true)
	gw.UiTab.Append(name, vbox)
	// mytab.SetMargined(mytabcount, true)

	vbox.Append(ui.NewVerticalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	a := CreateButton(gb, account, nil, "Add Virtual Machine", "createAddVmBox", nil)
	hbox.Append(a.B, false)

	vbox.Append(hbox, false)

	return mh
}

func MessageWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBox(gw.UiWindow, msg1, msg2)
}

func ErrorWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBoxError(gw.UiWindow, msg1, msg2)
}

// This is the default mouse click handler
// Every mouse click that hasn't been assigned to
// something specific will fall into this routine
// By default, all it runs is the call back to
// the main program that is using this library

// This is one of the routines that is called from the
// defaultButtonClick() below when the button is found
// in the AllButtons %map
// TODO: clean up the text above
// TODO: remove this all together going only to main()
func mouseClick(b *GuiButton) {
	log.Println("gui.mouseClick() START")
	if (b == nil) {
		log.Println("\tgui.mouseClick() START b = nil")
	} else {
		log.Println("\tgui.mouseClick() START b.Action =", b.Action)
		if (b.Action == "createAddVmBox") {
			log.Println("\tgui.mouseClick() createAddVmBox for b =", b)
			createAddVmBox(b.GW, b)
			return
		}
		/*
		if (b.Action == "WINDOW CLOSE") {
			b.W.Hide()
			// TODO: fix this (seems to crash? maybe because we are in the button here?)
			// b.W.Destroy()
			return
		}
		if (b.Action == "ADD") {
			log.Println("\tgui.mouseClick() SHOULD ADD VM HERE?")
		}
		*/
	}

	if (Data.MouseClick == nil) {
		log.Println("\tgui.mouseClick() Data.MouseClick() IS nil. NOT DOING ANYTHING")
		log.Println("\tgui.mouseClick() Your application did not set a MouseClick() callback function")
	} else {
		log.Println("\tgui.mouseClick() Data.MouseClick() START")
		Data.MouseClick(b)
	}
}

//
// This routine MUST be here as this is how the andlabs/ui works
// This is the raw routine passed to every button in andlabs libui / ui
//
// There is a []GuiButton which has all the buttons. We search
// for the button and then call the function below
//
func defaultButtonClick(button *ui.Button) {
	log.Println("defaultButtonClick() LOOK FOR BUTTON button =", button)
	for key, foo := range Data.AllButtons {
		if (Data.Debug) {
			log.Println("defaultButtonClick() Data.AllButtons =", key, foo)
			spew.Dump(foo)
		}
		if Data.AllButtons[key].B == button {
			log.Println("\tdefaultButtonClick() BUTTON MATCHED")
			// log.Println("\tData.AllButtons[key].Name =", Data.AllButtons[key].Name)
			log.Println("\tdefaultButtonClick() Data.AllButtons[key].Action =", Data.AllButtons[key].Action)
			if Data.AllButtons[key].custom != nil {
				log.Println("\tdefaultButtonClick() DOING CUSTOM FUNCTION")
				var tmp *GuiButton
				tmp = Data.AllButtons[key]
				// spew.Dump(tmp)
				Data.AllButtons[key].custom(tmp)
				return
			}
			mouseClick(Data.AllButtons[key])
			return
		}
	}
	log.Println("\tdefaultButtonClick() BUTTON NOT FOUND")
	if (Data.Debug) {
		panic("defaultButtonClick() SHOULD NOT HAVE UNMAPPED BUTTONS")
	}
	mouseClick(nil)
}

func AddButton(b *GuiButton, name string) *ui.Button {
	newB := ui.NewButton(name)
	newB.OnClicked(defaultButtonClick)

	b.B	= newB
	Data.AllButtons	= append(Data.AllButtons, b)
	return newB
}

func AddButtonToBox(box *GuiBox, button *GuiButton) {
	box.UiBox.Append(button.B, false)
}

func CreateButton(box *GuiBox, a *pb.Account, vm *pb.Event_VM, name string, action string, custom func(*GuiButton)) *GuiButton {
	newUiB := ui.NewButton(name)
	newUiB.OnClicked(defaultButtonClick)

	var newB *GuiButton
	newB		= new(GuiButton)
	newB.B		= newUiB
	if (box.W == nil) {
		log.Println("CreateButton() box.W == nil")
		panic("crap")
	}
	newB.GW		= box.W
	newB.Account	= a
	newB.VM		= vm
	newB.Box	= box
	newB.Action	= action
	newB.custom	= custom
	Data.AllButtons	= append(Data.AllButtons, newB)

	return newB
}

func CreateFontButton(box *GuiBox, action string) *GuiButton {

        // create a 'fake' button entry for the mouse clicks
	var newGB	GuiButton
	newGB.Action	= action
	newGB.FB	= ui.NewFontButton()
	newGB.Box	= box
	newGB.Area	= box.Area
	Data.AllButtons	= append(Data.AllButtons, &newGB)

	newGB.FB.OnChanged(func (*ui.FontButton) {
		log.Println("FontButton.OnChanged() START mouseClick(&newBM)", newGB)
		mouseClick(&newGB)
	})
	return &newGB
}

func GetText(box *GuiBox, name string) string {
	if (box == nil) {
		log.Println("gui.GetText() ERROR box == nil")
		return ""
	}
	if (box.EntryMap == nil) {
		log.Println("gui.GetText() ERROR b.Box.EntryMap == nil")
		return ""
	}
	spew.Dump(box.EntryMap)
	if (box.EntryMap[name] == nil) {
		log.Println("gui.GetText() ERROR box.EntryMap[", name, "] == nil ")
		return ""
	}
	e := box.EntryMap[name]
	log.Println("gui.GetText() box.EntryMap[", name, "] = ", e.UiEntry.Text())
	log.Println("gui.GetText() END")
	return e.UiEntry.Text()
}

func SetText(box *GuiBox, name string, value string) error {
	if (box == nil) {
		return fmt.Errorf("gui.SetText() ERROR box == nil")
	}
	if (box.EntryMap == nil) {
		return fmt.Errorf("gui.SetText() ERROR b.Box.EntryMap == nil")
	}
	spew.Dump(box.EntryMap)
	if (box.EntryMap[name] == nil) {
		return fmt.Errorf("gui.SetText() ERROR box.EntryMap[", name, "] == nil ")
	}
	e := box.EntryMap[name]
	log.Println("gui.SetText() box.EntryMap[", name, "] = ", e.UiEntry.Text())
	e.UiEntry.SetText(value)
	log.Println("gui.SetText() box.EntryMap[", name, "] = ", e.UiEntry.Text())
	log.Println("gui.SetText() END")
	return nil
}

// makeEntryBox(box, "hostname:", "blah.foo.org") {
func MakeEntryVbox(box *GuiBox, a string, startValue string, edit bool, action string) *GuiEntry {
	// Start 'Nickname' vertical box
	vboxN := ui.NewVerticalBox()
	vboxN.SetPadded(true)
	vboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)

	vboxN.Append(e.UiEntry, false)
	box.UiBox.Append(vboxN, false)
	// End 'Nickname' vertical box

	return e
}

func MakeEntryHbox(box *GuiBox, a string, startValue string, edit bool, action string) *GuiEntry {
	// Start 'Nickname' vertical box
	hboxN := ui.NewHorizontalBox()
	hboxN.SetPadded(true)
	hboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)
	hboxN.Append(e.UiEntry, false)

	box.UiBox.Append(hboxN, false)
	// End 'Nickname' vertical box

	return e
}

func NewLabel(box *GuiBox, text string) {
	box.UiBox.Append(ui.NewLabel(text), false)
}

func AddEntry(box *GuiBox, name string) *GuiEntry {
	var ge *GuiEntry
	ge = new(GuiEntry)

	ue := ui.NewEntry()
	ue.SetReadOnly(false)
	ue.OnChanged(func(*ui.Entry) {
		log.Println("gui.AddEntry() OK. ue.Text() =", ue.Text())
	})
	box.UiBox.Append(ue, false)

	ge.UiEntry = ue
	box.EntryMap[name] = ge

	return ge
}

func AddGenericBox(gw *GuiWindow) *GuiBox {
	var gb *GuiBox
	gb = new(GuiBox)

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	// gw.Box1 = vbox
	gb.UiBox = vbox
	gb.W = gw

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	return gb
}

func HorizontalBreak(box *GuiBox) {
	tmp := ui.NewHorizontalSeparator()
	box.UiBox.Append(tmp, false)
}
