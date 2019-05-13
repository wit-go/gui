package gui

import "log"
import "time"
// import "fmt"

import "github.com/gookit/config"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

var mainwin *ui.Window
var maintab *ui.Tab
var tabcount int

var jcarrButton *ui.Button
var jcarrEntry  *ui.MultilineEntry

type InputData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
}

func setupUI() {
	mainwin = ui.NewWindow("Cloud Control Panel", config.Int("width"), config.Int("height"), false)
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

	// maintab.Append("List examples", makeNumbersPage())
	tabcount = 0
	// maintab.SetMargined(tabcount, true)

/*
	maintab.Append("Choosers examples", makeDataChoosersPage())
	tabcount += 1
	maintab.SetMargined(tabcount, true)

	maintab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	maintab.SetMargined(tabcount, true)
*/

	mainwin.Show()
}

func AddChoosersDemo() {
	maintab.Append("Choosers examples", makeDataChoosersPage())
	maintab.SetMargined(tabcount, true)
	tabcount += 1
}

func AddNewTab(mytab *ui.Tab, newbox ui.Control, tabOffset int) {
	mytab.Append("Cloud Info", newbox)
	mytab.SetMargined(tabOffset, true)
}

// This hangs on GTK
func AddEntriesDemo() {
	maintab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	maintab.SetMargined(tabcount, true)
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
	for key, foo := range parts {
		log.Println(key, foo)
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

func AddSampleTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []InputData) {
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

	time.Sleep(1 * 1000 * 1000 * 1000)

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
			table.AppendButtonColumn("button3", tmpBTindex, ui.TableModelColumnAlwaysEditable)
		} else if (foo.CellType == "TEXTCOLOR") {
			tmpBTindex += 1
			appendTextColorColumn   (mh, table, tmpBTindex, tmpBTindex + 1, "testcolor")
			tmpBTindex += 1
		} else if (foo.CellType == "TEXT") {
			tmpBTindex += 1
			appendTextColumn        (mh, table, tmpBTindex, "jwc1col")
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}

	mytab.Append(name, table)
	mytab.SetMargined(mytabcount, true)
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
	ui.Main(setupUI)

	log.Println("GUI exited. Not sure what to do here. os.Exit() ?")

	// not sure how to pass this back to the main program
	// onExit()
}
