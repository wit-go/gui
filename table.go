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

type cellData struct {
	index		int
	value		ui.TableValue
	name		string			// what type of cell is this?
	event		func()			// what function to call if there is an event on this
}

// hmm. will this stand the test of time?
type rowData struct {
	name		string			// what kind of row is this?
	status		string			// status of the row?
/*
	// These may or may not be implementable
	click		func()			// what function to call if the user clicks on it
	doubleclick	func()			// what function to call if the user double clicks on it
*/
	cells		[20]cellData
}

type tableData struct {
	rowcount		int			// This is the number of 'rows' which really means data elements not what the human sees
	rowWidth		int			// This is how wide each row is
	rows			[]rowData		// This is all the table data by row
	generatedColumnTypes	[]ui.TableValue		// generate this dynamically
	libUIevent 	 	func(*tableData, *ui.TableModel, int, int, ui.TableValue)
}

func initRowBTcolor(mh *tableData, row int, intBG int) {
	// alternate background of each row light and dark
	if (row % 2) == 1 {
		mh.rows[row].cells[intBG].value = ui.TableColor{0.5, 0.5, 0.5, .7}
		mh.rows[row].cells[intBG].name = "BG"
	} else {
		mh.rows[row].cells[intBG].value = ui.TableColor{0.1, 0.1, 0.1, .1}
		mh.rows[row].cells[intBG].name = "BG"
	}
}

func initRowButtonColumn(mh *tableData, row int, buttonID int, junk string) {
	// set the button text for Column ?
	mh.rows[row].cells[buttonID].value = ui.TableString(fmt.Sprintf("%s %d", junk, row))
	mh.rows[row].cells[buttonID].name = "BUTTON"
}

func initRowTextColorColumn(mh *tableData, row int, stringID int, colorID int, junk string, color ui.TableColor) {
	// text for Column ?
	mh.rows[row].cells[stringID].value = ui.TableString(fmt.Sprintf("%s %d", junk, row))
	mh.rows[row].cells[stringID].name = "EDIT"

	// text color for Column ?
	mh.rows[row].cells[colorID].value = color
	mh.rows[row].cells[colorID].name = "COLOR"
}

func initRowTextColumn(mh *tableData, row int, stringID int, junk string) {
	mh.rows[row].cells[stringID].value = ui.TableString(fmt.Sprintf("%s %d", junk, row))
	mh.rows[row].cells[stringID].name = "EDIT"
}

func appendTextColorColumn(mh *tableData, table *ui.Table, stringID int, colorID int, columnName string) {
	table.AppendTextColumn(columnName, stringID, ui.TableModelColumnAlwaysEditable,
		&ui.TableTextColumnOptionalParams{
			ColorModelColumn:               colorID,
	});
}

func appendTextColumn(mh *tableData, table *ui.Table, stringID int, columnName string) {
	table.AppendTextColumn(columnName, stringID, ui.TableModelColumnAlwaysEditable, nil)
}

func defaultSetCellValue(mh *tableData, m *ui.TableModel, row, column int, value ui.TableValue) {
	if (mh.rows[row].cells[column].name == "EDIT") {
		mh.rows[row].cells[column].value = value
	}
	if (mh.rows[row].cells[column].name == "BUTTON") {
		log.Println("FOUND THE BUTTON!!!!!!!   Button was pressed START", row, column)
	}
	return
}
