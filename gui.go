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

func AddTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []TableColumnData, account *pb.Account) *TableData {
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
	mytab.SetMargined(mytabcount, true)

	vbox.Append(ui.NewVerticalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	hbox.Append(CreateButton(account, nil, "Add Virtual Machine", "createAddVmBox", mouseClick), false)
	hbox.Append(CreateButton(account, nil, "Close", "CLOSE", mouseClick), false)

	vbox.Append(hbox, false)

	return mh
}

func SocketError() {
	ui.MsgBoxError(Data.cloudWindow,
		"There was a socket error",
		"More detailed information can be shown here.")
}

func MessageWindow(msg1 string, msg2 string) {
	ui.MsgBox(Data.cloudWindow, msg1, msg2)
}

func ErrorWindow(msg1 string, msg2 string) {
	ui.MsgBoxError(Data.cloudWindow, msg1, msg2)
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
func mouseClick(b *ButtonMap) {
	log.Println("gui.mouseClick() START b =", b)

	if (b != nil) {
		if (b.Action == "createAddVmBox") {
			log.Println("gui.mouseClick() createAddVmBox for b =", b)
			createAddVmBox(Data.cloudTab, "Create New Virtual Machine", mouseClick)
			return
		}
	}

	if (Data.MouseClick == nil) {
		log.Println("Data.MouseClick() IS nil. NOT DOING ANYTHING")
	} else {
		log.Println("\tData.MouseClick() START")
		Data.MouseClick(b)
	}
}

/*
func buttonMapClick(b *ButtonMap) {
	log.Println("gui.buttonVmClick() START")
	if (Data.MouseClick == nil) {
		log.Println("Data.MouseClick() IS nil. NOT DOING ANYTHING")
	} else {
		log.Println("Data.MouseClick() START")
		Data.MouseClick(b)
	}
}
*/

// This is the raw routine passed to every button in andlabs libui / ui
func defaultButtonClick(button *ui.Button) {
	log.Println("defaultButtonClick() button =", button)
	for key, foo := range Data.AllButtons {
		log.Println("Data.AllButtons =", key, foo)
		if Data.AllButtons[key].B == button {
			log.Println("\tBUTTON MATCHED")
			// log.Println("\tData.AllButtons[key].Name =", Data.AllButtons[key].Name)
			log.Println("\tData.AllButtons[key].Action =", Data.AllButtons[key].Action)
			if Data.AllButtons[key].custom != nil {
				log.Println("\tDOING CUSTOM FUNCTION")
				var tmp ButtonMap
				tmp = Data.AllButtons[key]
				spew.Dump(tmp)
				Data.AllButtons[key].custom(&tmp)
				return
			}
			if (Data.MouseClick != nil) {
				Data.MouseClick(&Data.AllButtons[key])
			}
			return
		}
	}
	log.Println("\tBUTTON NOT FOUND")
	// still run the mouse click handler
	if (Data.MouseClick != nil) {
		Data.MouseClick(nil)
	}
}

// This is the raw routine passed to every button in andlabs libui / ui
// (this has to be different for FontButtons)
// TODO; merge the logic with the function above
func defaultFontButtonClick(button *ui.FontButton) {
	log.Println("defaultFontButtonClick() button =", button)
	for key, foo := range Data.AllButtons {
		log.Println("Data.AllButtons =", key, foo)
	}
}

func CreateButton(a *pb.Account, vm *pb.Event_VM,
		name string, note string, custom func(*ButtonMap)) *ui.Button {
	newB := ui.NewButton(name)
	newB.OnClicked(defaultButtonClick)

	var newmap ButtonMap
	newmap.B = newB
	newmap.Account = a
	newmap.VM      = vm
	newmap.Action  = note
	newmap.custom  = custom
	newmap.aTab    = Data.CurrentTab
	Data.AllButtons = append(Data.AllButtons, newmap)

	return newB
}

func CreateFontButton(name string, note string, custom func(*ButtonMap)) *ui.FontButton {
	newB := ui.NewFontButton()

	newB.OnChanged(defaultFontButtonClick)

	var newmap ButtonMap
	newmap.FB = newB
	newmap.Action = note
	newmap.custom = custom
	Data.AllButtons = append(Data.AllButtons, newmap)

	return newB
}
