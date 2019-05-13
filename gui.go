package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var mainwin	*ui.Window
var maintab	*ui.Tab
var tabcount	int

var Width	int
var Height	int

type InputData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
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

func initColumnNames(mh *TableData, cellJWC string, junk string) {
	if (cellJWC == "BG") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableColor{})
	} else if (cellJWC == "BUTTON") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
	} else if (cellJWC == "TEXTCOLOR") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableColor{})
	} else if (cellJWC == "TEXT") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
	} else {
		panic("I don't know what this is in initColumnNames")
	}
}

func initRow(mh *TableData, row int, parts []InputData) {
	tmpBTindex := 0
	humanID := 0
	for key, foo := range parts {
		log.Println(key, foo)

		parts[key].Index = humanID
		humanID += 1

		if (foo.CellType == "BG") {
			initRowBTcolor        (mh, row, tmpBTindex, parts[key])
			tmpBTindex += 1
		} else if (foo.CellType == "BUTTON") {
			initRowButtonColumn   (mh, row, tmpBTindex,    parts[key].Heading, parts[key])
			tmpBTindex += 1
		} else if (foo.CellType == "TEXTCOLOR") {
			initRowTextColorColumn(mh, row, tmpBTindex, tmpBTindex + 1, parts[key].Heading, ui.TableColor{0.0, 0, 0.9, 1}, parts[key])
			tmpBTindex += 2
		} else if (foo.CellType == "TEXT") {
			initRowTextColumn     (mh, row, tmpBTindex,    parts[key].Heading, parts[key])
			tmpBTindex += 1
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}
}

func AddTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []InputData) *TableData {
	mh := new(TableData)

	mh.RowCount    = rowcount
	mh.Rows        = make([]RowData, mh.RowCount)

	// This is the standard callback function from libUI when the user does something
	mh.libUIevent      = defaultSetCellValue

	tmpBTindex := 0

	for key, foo := range parts {
		log.Println(key, foo)
		initColumnNames(mh, foo.CellType, foo.Heading)
	}

	for row := 0; row < mh.RowCount; row++ {
		initRow(mh, row, parts)
	}
	log.Println(mh)

	model := ui.NewTableModel(mh)
	table := ui.NewTable(
		&ui.TableParams{
			Model:	model,
			RowBackgroundColorModelColumn:	tmpBTindex,
	})

	for key, foo := range parts {
		log.Println(key, foo)
		initColumnNames(mh, foo.CellType, foo.Heading)
		if (foo.CellType == "BG") {
		} else if (foo.CellType == "BUTTON") {
			tmpBTindex += 1
			table.AppendButtonColumn(foo.Heading, tmpBTindex, ui.TableModelColumnAlwaysEditable)
		} else if (foo.CellType == "TEXTCOLOR") {
			tmpBTindex += 1
			appendTextColorColumn   (mh, table, tmpBTindex, tmpBTindex + 1, foo.Heading)
			tmpBTindex += 1
		} else if (foo.CellType == "TEXT") {
			tmpBTindex += 1
			appendTextColumn        (mh, table, tmpBTindex, foo.Heading)
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}

	mytab.Append(name, table)
	mytab.SetMargined(mytabcount, true)

	return mh
}

func DoGUI() {
	for {
		ui.Main(setupUI)
		log.Println("GUI exited. Not sure what to do here. os.Exit() ?")
	}
}
