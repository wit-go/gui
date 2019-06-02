package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

// THIS IS CLEAN (all that is left is the 'ADD VM')

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

//	gb.EntryMap = make(map[string]*GuiEntry)
//	gb.EntryMap["test"] = nil

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	gb.UiBox = vbox
	gb.Window = gw
	gw.BoxMap[name] = gb
	mh.Box = gb

	vbox.Append(table, true)
	gw.UiTab.Append(name, vbox)

	vbox.Append(ui.NewVerticalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

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

func GetText(box *GuiBox, name string) string {
	if (box == nil) {
		log.Println("gui.GetText() ERROR box == nil")
		return ""
	}
	if (box.Window.EntryMap == nil) {
		log.Println("gui.GetText() ERROR b.Box.Window.EntryMap == nil")
		return ""
	}
	spew.Dump(box.Window.EntryMap)
	if (box.Window.EntryMap[name] == nil) {
		log.Println("gui.GetText() ERROR box.Window.EntryMap[", name, "] == nil ")
		return ""
	}
	e := box.Window.EntryMap[name]
	log.Println("gui.GetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
	log.Println("gui.GetText() END")
	return e.UiEntry.Text()
}

func SetText(box *GuiBox, name string, value string) error {
	if (box == nil) {
		return fmt.Errorf("gui.SetText() ERROR box == nil")
	}
	if (box.Window.EntryMap == nil) {
		return fmt.Errorf("gui.SetText() ERROR b.Box.Window.EntryMap == nil")
	}
	spew.Dump(box.Window.EntryMap)
	if (box.Window.EntryMap[name] == nil) {
		return fmt.Errorf("gui.SetText() ERROR box.Window.EntryMap[", name, "] == nil ")
	}
	e := box.Window.EntryMap[name]
	log.Println("gui.SetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
	e.UiEntry.SetText(value)
	log.Println("gui.SetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
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
	box.Window.EntryMap[name] = ge

	return ge
}

func HardHorizontalBreak(box *GuiBox) {
	log.Println("HardHorizontalBreak START")
	gw := box.Window
	mainbox := gw.BoxMap["MAIN"]
	if (mainbox == nil) {
		log.Println("HardHorizontalBreak ERROR MAIN box == nil")
		return
	}
	uibox := mainbox.UiBox

	tmp := ui.NewHorizontalSeparator()
	uibox.Append(tmp, false)

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(true)
	box.UiBox = hbox
	uibox.Append(hbox, true)
	log.Println("HardHorizontalBreak END")
}

func HardVerticalBreak(box *GuiBox) {
	log.Println("HardVerticalBreak START")
	gw := box.Window
	mainbox := gw.BoxMap["MAIN"]
	if (mainbox == nil) {
		log.Println("HardHorizontalBreak ERROR MAIN box == nil")
		return
	}

	tmp := ui.NewVerticalSeparator()
	mainbox.UiBox.Append(tmp, false)

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(true)
	box.UiBox = hbox
	mainbox.UiBox.Append(hbox, false)
	log.Println("HardVerticalBreak END")
}

func HorizontalBreak(box *GuiBox) {
	tmp := ui.NewHorizontalSeparator()
	box.UiBox.Append(tmp, false)
}

func VerticalBreak(box *GuiBox) {
	tmp := ui.NewVerticalSeparator()
	box.UiBox.Append(tmp, false)
}

func AddGenericBox(gw *GuiWindow) *GuiBox {
	var gb *GuiBox
	gb = new(GuiBox)

//	gb.EntryMap = make(map[string]*GuiEntry)
//	gb.EntryMap["test"] = nil

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	// gw.Box1 = vbox
	gb.UiBox = vbox
	gb.Window = gw

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	return gb
}

func CreateGenericBox(gw *GuiWindow, b *GuiButton, name string) *GuiBox{
	log.Println("CreateAddVmBox() START name =", name)

	var box *GuiBox
	box = new(GuiBox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	box.UiBox = vbox
	box.Window = gw
	gw.BoxMap["ADD VM" + name] = box

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	AddBoxToTab(name, gw.UiTab, vbox)

	return box
}

func CreateBox(gw *GuiWindow, name string) *GuiBox {
	log.Println("CreateVmBox() START")
	log.Println("CreateVmBox() vm.Name =", name)
	log.Println("CreateVmBox() gw =", gw)

	var box *GuiBox
	box = new(GuiBox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	log.Println("CreateVmBox() vbox =", vbox)
	log.Println("CreateVmBox() box.UiBox =", box.UiBox)
	box.UiBox = vbox
	log.Println("CreateVmBox() box.Window =", box.Window)
	box.Window = gw
	log.Println("CreateVmBox() gw.BoxMap =", gw.BoxMap)
	gw.BoxMap[name] = box

	hboxAccount := ui.NewHorizontalBox()
	hboxAccount.SetPadded(true)
	vbox.Append(hboxAccount, false)

	box.UiBox = hboxAccount

	AddBoxToTab(name, gw.UiTab, vbox)

	return box
}
