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

func InitColumns(mh *TableData, parts []InputData) {
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

func AddTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []InputData) *TableData {
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
