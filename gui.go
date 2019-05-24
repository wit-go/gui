package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/gookit/config"
import "github.com/davecgh/go-spew/spew"

//
// All GUI Data Structures and functions that are external
// If you need cross platform support, these might only
// be the safe way to interact with the GUI
//
var Data	GuiDataStructure

type GuiTabStructure struct {
	me		*ui.Tab
	parentWindow	*ui.Window
	firstBox	*ui.Box
	tabOffset	int

	// this means only one table per tab
	mh		*TableData

	// stuff for the 'area'
	// this means only one area per tab
	fontButton	*ui.FontButton
	attrstr		*ui.AttributedString
	splashArea	*ui.Area
}

type GuiDataStructure struct {
	State		string
	Width		int
	Height		int

	// a fallback default function to handle mouse events 
	// if nothing else is defined to handle them
	MouseClick	func(*ButtonMap)

	// general information
	Version		string
	GitCommit	string
	GoVersion	string
	Buildtime	string
	HomeDir		string
	Debug		bool
	DebugTable	bool

	// official hostname and IPv6 address for this box
	Hostname	string
	IPv6		string

	// account entry textboxes
	AccNick	string
	AccUser	string
	AccPass	string

	// A map of all buttons everywhere on all
	// windows, all tabs, across all goroutines
	// This is "GLOBAL"
	AllButtons	[]ButtonMap

	// a tab (maybe the one the user is playing with?)
	CurrentTab	*GuiTabStructure
	// a VM (maybe the one the user is playing with?)
	CurrentVM	string

	// All the tabs
	Tabs		[]GuiTabStructure

	// stuff for the splash screen / setup tabs
	// MainWindow	*ui.Window
	cloudWindow	*ui.Window
	cloudTab	*ui.Tab
	cloudBox	*ui.Box
	smallBox	*ui.Box

	mainwin		*ui.Window
	maintab		*ui.Tab
	tabcount	int

	// stuff for the 'area'
	MyArea		*ui.Area
	fontButton	*ui.FontButton
	attrstr		*ui.AttributedString
	splashArea	*ui.Area
}

type TableColumnData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
}

type ButtonMap struct {
	B		*ui.Button
	FB		*ui.FontButton
	onClick		func (int, string)
	onChanged	func (int, string)
	custom		func (*ButtonMap, string)
	Name		string	// the text on the button
	Note		string	// what type of button
	AccNick		string  // what account this button is for
}

/*
func AddNewTab(mytab *ui.Tab, newbox ui.Control, tabOffset int) {
	mytab.Append("Cloud Info", newbox)
	mytab.SetMargined(tabOffset, true)
}
*/

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

func AddTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []TableColumnData) *TableData {
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

	myAddVM := addVmButton("Add Virtual Machine")
	hbox.Append(myAddVM, false)

	myClose := closeButton("Close", mytab)
	hbox.Append(myClose, false)

	vbox.Append(hbox, false)

	return mh
}

// This is the default mouse click handler
// Every mouse click that hasn't been assigned to
// something specific will fall into this routine
// By default, all it runs is the call back to
// the main program that is using this library
func mouseClick(b *ButtonMap, s string) {
	log.Println("gui.mouseClick() START b, s =", b, s)

	if (Data.MouseClick != nil) {
		log.Println("\tData.MouseClick() START")
		Data.MouseClick(b)
	}
}

func defaultButtonClick(button *ui.Button) {
	log.Println("defaultButtonClick() button =", button)
	for key, foo := range Data.AllButtons {
		log.Println("Data.AllButtons =", key, foo)
		if Data.AllButtons[key].B == button {
			log.Println("\tBUTTON MATCHED")
			log.Println("\tData.AllButtons[key].Name", Data.AllButtons[key].Name)
			log.Println("\tData.AllButtons[key].Note", Data.AllButtons[key].Note)
			if Data.AllButtons[key].custom != nil {
				log.Println("\tDOING CUSTOM FUNCTION")
				Data.AllButtons[key].custom(&Data.AllButtons[key], "SOMETHING CUSTOM")
				return
			}
			if (Data.MouseClick != nil) {
				Data.MouseClick( &Data.AllButtons[key])
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

func defaultFontButtonClick(button *ui.FontButton) {
	log.Println("defaultButtonClick() button =", button)
	for key, foo := range Data.AllButtons {
		log.Println("Data.AllButtons =", key, foo)
	}
}

func CreateButton(name string, note string, custom func(*ButtonMap, string)) *ui.Button {
	newB := ui.NewButton(name)

	newB.OnClicked(defaultButtonClick)

	var newmap ButtonMap
	newmap.B = newB
	newmap.Note = note
	newmap.Name = name
	newmap.custom = custom
	Data.AllButtons = append(Data.AllButtons, newmap)

	return newB
}

func CreateAccountButton(account string, custom func(*ButtonMap, string)) *ui.Button {
	name := "Show " + account
	newB := ui.NewButton(name)

	newB.OnClicked(defaultButtonClick)

	var newmap ButtonMap
	newmap.B = newB
	newmap.Note = "SHOW"
	newmap.Name = name
	newmap.AccNick = account
	newmap.custom = custom
	Data.AllButtons = append(Data.AllButtons, newmap)

	return newB
}

func CreateLoginButton(account string, custom func(*ButtonMap, string)) *ui.Button {
	name := "Login " + account
	newB := ui.NewButton(name)

	newB.OnClicked(defaultButtonClick)

	var newmap ButtonMap
	newmap.B = newB
	newmap.Note = "LOGIN"
	newmap.Name = name
	newmap.AccNick = account
	newmap.custom = custom
	Data.AllButtons = append(Data.AllButtons, newmap)

	return newB
}

func CreateFontButton(name string, note string, custom func(*ButtonMap, string)) *ui.FontButton {
	newB := ui.NewFontButton()

	newB.OnChanged(defaultFontButtonClick)

	var newmap ButtonMap
	newmap.FB = newB
	newmap.Note = note
	newmap.custom = custom
	Data.AllButtons = append(Data.AllButtons, newmap)

	return newB
}

func closeButtonClick(button *ui.Button) {
	log.Println("closeButtonClick() hostname =", config.String("hostname"), button)
	spew.Dump(button)
}

func closeButton(name string, mytab *ui.Tab) ui.Control {
	tmpButton := ui.NewButton(name)
	tmpButton.OnClicked(defaultButtonClick)

	return tmpButton
}

func addVmButtonClick(button *ui.Button) {
	log.Println("addVMButtonClick()")
	spew.Dump(button)
}

func addVmButton(name string) ui.Control {
	tmpButton := ui.NewButton(name)
	tmpButton.OnClicked(addVmButtonClick)

	return tmpButton
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
