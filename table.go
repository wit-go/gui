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
	HumanData	[20]HumanCellData
}

// hmm. will this stand the test of time?
type HumanCellData struct {
	Name		string			// what kind of row is this?
	Text		ui.TableString
	TextID		int
	Color		ui.TableColor
	ColorID		int
}

type HumanMap struct {
	Name		string			// what kind of row is this?
	TextID		int
	ColorID		int
}

type TableData struct {
	RowCount		int			// This is the number of 'rows' which really means data elements not what the human sees
	RowWidth		int			// This is how wide each row is
	Rows			[]RowData		// This is all the table data by row
	generatedColumnTypes	[]ui.TableValue		// generate this dynamically

	Cells			[20]CellData
	Human			[20]HumanMap
}

func initRowBTcolor(mh *TableData, intBG int, cell TableColumnData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Human[humanInt].Name    = "BG"
	mh.Human[humanInt].ColorID = intBG
	mh.Human[humanInt].TextID  = -1

	mh.Cells[intBG].Name       = "BG"
	mh.Cells[intBG].HumanID    = humanInt

	log.Println("intBG, humanInt", intBG, humanInt)
}

func initRowButtonColumn(mh *TableData, buttonID int, junk string, cell TableColumnData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Human[humanInt].Name    = "BUTTON"
	mh.Human[humanInt].ColorID = -1
	mh.Human[humanInt].TextID  = buttonID

	mh.Cells[buttonID].Name    = "BUTTON"
	mh.Cells[buttonID].HumanID = humanInt

	log.Println("buttonID, humanInt", buttonID, humanInt)
}

func initRowTextColorColumn(mh *TableData, stringID int, colorID int, junk string, color ui.TableColor, cell TableColumnData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Human[humanInt].Name    = "EDIT"
	mh.Human[humanInt].ColorID = colorID
	mh.Human[humanInt].TextID  = stringID

	// text for Column humanInt
	mh.Cells[stringID].Name    = "EDIT"
	mh.Cells[stringID].HumanID = humanInt

	mh.Cells[colorID].Name     = "COLOR"
	mh.Cells[colorID].HumanID  = humanInt
}

func initRowTextColumn(mh *TableData, stringID int, junk string, cell TableColumnData) {
	humanInt := cell.Index

	// setup mapping from human readable indexes to internal libUI indexes
	mh.Human[humanInt].Name    = "EDIT"
	mh.Human[humanInt].ColorID = -1
	mh.Human[humanInt].TextID  = stringID

	mh.Cells[stringID].Name    = "EDIT"
	mh.Cells[stringID].HumanID = humanInt
}
