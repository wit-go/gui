package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/gookit/config"
import "github.com/davecgh/go-spew/spew"

var mainwin	*ui.Window
var maintab	*ui.Tab
var tabcount	int

var Width	int
var Height	int

type TableColumnData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
}

type ButtonMap struct {
	B		*ui.Button
}

func setupUI() {
	mainwin = ui.NewWindow("Cloud Control Panel", Width, Height, false)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	maintab = ui.NewTab()
	mainwin.SetChild(maintab)
	mainwin.SetMargined(true)

	tabcount = 0
	mainwin.Show()
}

func AddNewTab(mytab *ui.Tab, newbox ui.Control, tabOffset int) {
	mytab.Append("Cloud Info", newbox)
	mytab.SetMargined(tabOffset, true)
}

func InitColumns(mh *TableData, parts []TableColumnData) {
	tmpBTindex := 0
	humanID := 0
	for key, foo := range parts {
		log.Println("key, foo =", key, foo)
		// log.Println("mh.Cells =", mh.Cells)
		// log.Println("mh.Human =", mh.Human)

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

func DoGUI() {
	for {
		ui.Main(setupUI)
		log.Println("GUI exited. Not sure what to do here. os.Exit() ?")
	}
}

func closeButtonClick(button *ui.Button) {
	log.Println("closeButtonClick() hostname =", config.String("hostname"), button)
	spew.Dump(button)
}

func closeButton(name string, mytab *ui.Tab) ui.Control {
	tmpButton := ui.NewButton(name)
	tmpButton.OnClicked(buttonClick)

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
