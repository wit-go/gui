package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

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

func AddTableTab(wm *GuiWindow, mytab *ui.Tab, junk int, name string, rowcount int, parts []TableColumnData, account *pb.Account) *TableData {
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

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(table, true)
	mytab.Append(name, vbox)
	// mytab.SetMargined(mytabcount, true)

	vbox.Append(ui.NewVerticalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	hbox.Append(CreateButton(wm, account, nil, "Add Virtual Machine", "createAddVmBox", nil), false)
	hbox.Append(CreateButton(wm, account, nil, "Close", "CLOSE", nil), false)

	vbox.Append(hbox, false)

	return mh
}

func SocketError(wm *GuiWindow) {
	ui.MsgBoxError(wm.W,
		"There was a socket error",
		"More detailed information can be shown here.")
}

func MessageWindow(wm *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBox(wm.W, msg1, msg2)
}

func ErrorWindow(wm *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBoxError(wm.W, msg1, msg2)
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
			createAddVmBox(b.WM, b.T, "Create New Virtual Machine", b)
			return
		}
		/*
		if (b.Action == "SHOW VM") {
			Data.CurrentVM = b.VM
			if (Data.Debug) {
				go ui.Main(ShowVM)
			} else {
				CreateVmBox(Data.Window1.T, b.VM)
			}
		}
		*/
		if (b.Action == "WINDOW CLOSE") {
			b.W.Hide()
			// TODO: fix this (seems to crash? maybe because we are in the button here?)
			// b.W.Destroy()
			return
		}
		if (b.Action == "ADD") {
			log.Println("\tgui.mouseClick() SHOULD ADD VM HERE?")
		}
	}

	if (Data.MouseClick == nil) {
		log.Println("\tgui.mouseClick() Data.MouseClick() IS nil. NOT DOING ANYTHING")
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
				var tmp GuiButton
				tmp = Data.AllButtons[key]
				// spew.Dump(tmp)
				Data.AllButtons[key].custom(&tmp)
				return
			}
			mouseClick(&Data.AllButtons[key])
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
	Data.AllButtons	= append(Data.AllButtons, *b)
	return newB
}

func CreateButton(wm *GuiWindow, a *pb.Account, vm *pb.Event_VM,
		name string, action string, custom func(*GuiButton)) *ui.Button {
	newUiB := ui.NewButton(name)
	newUiB.OnClicked(defaultButtonClick)

	var newB GuiButton
	newB.B	= newUiB
	newB.T	= wm.T
	newB.Account	= a
	newB.VM	= vm
	newB.WM	= wm
	newB.Action	= action
	newB.custom	= custom
	Data.AllButtons	= append(Data.AllButtons, newB)

	return newUiB
}

func CreateFontButton(wm *GuiWindow, action string) *GuiButton {
	newB := ui.NewFontButton()

        // create a 'fake' button entry for the mouse clicks
	var newBM	GuiButton
	newBM.Action	= action
	newBM.FB	= newB
	newBM.AH	= wm.AH
	Data.AllButtons	= append(Data.AllButtons, newBM)

	newB.OnChanged(func (*ui.FontButton) {
		log.Println("FontButton.OnChanged() START mouseClick(&newBM)", &newBM)
		mouseClick(&newBM)
	})
	return &newBM
}
