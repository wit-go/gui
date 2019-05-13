// based off andlabs/ui/examples/table.go

package gui

import "fmt"
import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

var img [2]*ui.Image

/*
        img[0] = ui.NewImage(16, 16)
        img[1] = ui.NewImage(16, 16)
*/

type CellData struct {
	Index		int
	HumanID		int
	Value		ui.TableValue
	Raw		string			// shove stuff in here and figure out how to make a ui.TableValue later
	Name		string			// what type of cell is this?
	Event		func()			// what function to call if there is an event on this
}

// hmm. will this stand the test of time?
type RowData struct {
	Name		string			// what kind of row is this?
	Status		string			// status of the row?
/*
	// These may or may not be implementable
	click		func()			// what function to call if the user clicks on it
	doubleclick	func()			// what function to call if the user double clicks on it
*/
	Cells		[20]CellData
	Human		[20]HumanCellData
}

// hmm. will this stand the test of time?
type HumanCellData struct {
	Name		string			// what kind of row is this?
	Text		ui.TableString
	TextID		int
	Color		ui.TableColor
	ColorID		int
	Event		func()			// what function to call if there is an event on this
}

type TableData struct {
	RowCount		int			// This is the number of 'rows' which really means data elements not what the human sees
	RowWidth		int			// This is how wide each row is
	Rows			[]RowData		// This is all the table data by row
	generatedColumnTypes	[]ui.TableValue		// generate this dynamically
	libUIevent 	 	func(*TableData, *ui.TableModel, int, int, ui.TableValue)
	cellChangeEvent  	func(int, int, ui.TableValue)
}

func initRowBTcolor(mh *TableData, row int, intBG int, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "BG"
	mh.Rows[row].Human[humanInt].ColorID = intBG
	mh.Rows[row].Human[humanInt].TextID  = -1

	mh.Rows[row].Cells[intBG].Name       = "BG"
	mh.Rows[row].Cells[intBG].HumanID    = humanInt

	// alternate background of each row light and dark
	if (row % 2) == 1 {
		mh.Rows[row].Cells[intBG].Value   = ui.TableColor{0.5, 0.5, 0.5, .7}
	} else {
		mh.Rows[row].Cells[intBG].Value   = ui.TableColor{0.1, 0.1, 0.1, .1}
	}
}

func initRowButtonColumn(mh *TableData, row int, buttonID int, junk string, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "BUTTON"
	mh.Rows[row].Human[humanInt].ColorID = -1
	mh.Rows[row].Human[humanInt].TextID  = buttonID

	mh.Rows[row].Cells[buttonID].Value   = ui.TableString(fmt.Sprintf("%s %d", junk, row))
	mh.Rows[row].Cells[buttonID].Name    = "BUTTON"
	mh.Rows[row].Cells[buttonID].HumanID = humanInt
}

func initRowTextColorColumn(mh *TableData, row int, stringID int, colorID int, junk string, color ui.TableColor, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "EDIT"
	mh.Rows[row].Human[humanInt].ColorID = colorID
	mh.Rows[row].Human[humanInt].TextID  = stringID

	// text for Column humanInt
	mh.Rows[row].Cells[stringID].Value   = ui.TableString(fmt.Sprintf("%s %d", junk, row))
	mh.Rows[row].Cells[stringID].Name    = "EDIT"
	mh.Rows[row].Cells[stringID].HumanID = humanInt

	mh.Rows[row].Cells[colorID].Value    = color
	mh.Rows[row].Cells[colorID].Name     = "COLOR"
	mh.Rows[row].Cells[colorID].HumanID  = humanInt
}

func initRowTextColumn(mh *TableData, row int, stringID int, junk string, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "EDIT"
	mh.Rows[row].Human[humanInt].ColorID = -1
	mh.Rows[row].Human[humanInt].TextID  = stringID

	mh.Rows[row].Cells[stringID].Name    = "EDIT"
	mh.Rows[row].Cells[stringID].Value   = ui.TableString(fmt.Sprintf("%s %d", junk, row))
	mh.Rows[row].Cells[stringID].HumanID = humanInt
}

func appendTextColorColumn(mh *TableData, table *ui.Table, stringID int, colorID int, columnName string) {
	table.AppendTextColumn(columnName, stringID, ui.TableModelColumnAlwaysEditable,
		&ui.TableTextColumnOptionalParams{
			ColorModelColumn:               colorID,
	});
}

func appendTextColumn(mh *TableData, table *ui.Table, stringID int, columnName string) {
	table.AppendTextColumn(columnName, stringID, ui.TableModelColumnAlwaysEditable, nil)
}

func defaultSetCellValue(mh *TableData, m *ui.TableModel, row, column int, value ui.TableValue) {
	if (mh.Rows[row].Cells[column].Name == "EDIT") {
		mh.Rows[row].Cells[column].Value = value
	}
	if (mh.Rows[row].Cells[column].Name == "BUTTON") {
		log.Println("FOUND THE BUTTON!!!!!!!   Button was pressed START", row, column)
	}
	return
}

func simpleSetCellValue(mh *TableData, row, column int, value string) {
	if (mh.Rows[row].Cells[column].Name == "EDIT") {
		mh.Rows[row].Cells[column].Value = ui.TableString(value)
	}
	if (mh.Rows[row].Cells[column].Name == "BUTTON") {
		log.Println("simpleSetCellValue() FOUND THE BUTTON!!!!!!!  Button was pressed:", row, column)
	}
	return
}
