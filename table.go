// based off andlabs/ui/examples/table.go

package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

type CellData struct {
	Index		int
	HumanID		int
	Name		string			// what type of cell is this?
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
	// Cells		[20]CellData
	Human		[20]HumanCellData
}

// hmm. will this stand the test of time?
type HumanCellData struct {
	Name		string			// what kind of row is this?
	Text		ui.TableString
	TextID		int
	Color		ui.TableColor
	ColorID		int
}

type TableData struct {
	RowCount		int			// This is the number of 'rows' which really means data elements not what the human sees
	RowWidth		int			// This is how wide each row is
	Rows			[]RowData		// This is all the table data by row
	generatedColumnTypes	[]ui.TableValue		// generate this dynamically
	// libUIevent 	 	func(*TableData, *ui.TableModel, int, int, ui.TableValue)
	cellChangeEvent  	func(int, int, ui.TableValue)

	Cells		[20]CellData
}

func initRowBTcolor(mh *TableData, row int, intBG int, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "BG"
	mh.Rows[row].Human[humanInt].ColorID = intBG
	mh.Rows[row].Human[humanInt].TextID  = -1

	mh.Cells[intBG].Name       = "BG"
	mh.Cells[intBG].HumanID    = humanInt

	log.Println("HumanID = row, intBG, humanInt", row, intBG, humanInt)
}

func initRowButtonColumn(mh *TableData, row int, buttonID int, junk string, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "BUTTON"
	mh.Rows[row].Human[humanInt].ColorID = -1
	mh.Rows[row].Human[humanInt].TextID  = buttonID

	mh.Cells[buttonID].Name    = "BUTTON"
	mh.Cells[buttonID].HumanID = humanInt

	log.Println("HumanID = row, buttonID, humanInt", row, buttonID, humanInt)
}

func initRowTextColorColumn(mh *TableData, row int, stringID int, colorID int, junk string, color ui.TableColor, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "EDIT"
	mh.Rows[row].Human[humanInt].ColorID = colorID
	mh.Rows[row].Human[humanInt].TextID  = stringID

	// text for Column humanInt
	mh.Cells[stringID].Name    = "EDIT"
	mh.Cells[stringID].HumanID = humanInt

	mh.Cells[colorID].Name     = "COLOR"
	mh.Cells[colorID].HumanID  = humanInt
}

func initRowTextColumn(mh *TableData, row int, stringID int, junk string, cell InputData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Rows[row].Human[humanInt].Name    = "EDIT"
	mh.Rows[row].Human[humanInt].ColorID = -1
	mh.Rows[row].Human[humanInt].TextID  = stringID

	mh.Cells[stringID].Name    = "EDIT"
	mh.Cells[stringID].HumanID = humanInt
}
