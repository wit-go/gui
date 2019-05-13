package gui

//
// These functions are the hooks to the andlabs libui
// They eventually feed information to the OS native GUI toolkits
// and feed back user interaction with the GUI
//

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (mh *TableData) NumRows(m *ui.TableModel) int {
	return mh.RowCount
}

// FYI: this routine seems to be called around 10 to 100 times a second for each table
func (mh *TableData) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return mh.generatedColumnTypes
}

// TODO: Figure out why this is being called 1000 times a second (10 times for each row & column)
// Nevermind this TODO. Who gives a shit. This is a really smart way to treat the OS toolkits
func (mh *TableData) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	humanID := mh.Cells[column].HumanID
	if (column == mh.Human[humanID].TextID) {
		return mh.Rows[row].HumanData[humanID].Text
	}
	if (column == mh.Human[humanID].ColorID) {
		return mh.Rows[row].HumanData[humanID].Color
	}
	panic("not sure what sort of ui.TableValue to return in CellValue()")
	return ui.TableString("")
}

func (mh *TableData) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	log.Println("SetCellValue() START row=", row, "column=", column, "value=", value)

	defaultSetCellValue(mh, row, column)

	if (mh.cellChangeEvent != nil) {
		mh.cellChangeEvent(row, column, value)
	}

	log.Println("mh.Cells[column].HumanID =", mh.Cells[column].HumanID)
	// log.Println("mh.Rows[row].Cells[column].HumanID =", mh.Rows[row].Cells[column].HumanID)

	humanID := mh.Cells[column].HumanID
	log.Println("mh.Human[humanID].ColorID =", mh.Human[humanID].ColorID)
	log.Println("mh.Human[humanID].TextID =",  mh.Human[humanID].TextID)

	log.Println("SetCellValue() END")
}

func defaultSetCellValue(mh *TableData, row int, column int) {
	if (mh.Cells[column].Name == "BUTTON") {
		log.Println("defaultSetCellValue() FOUND THE BUTTON!!!!!!!   Button was pressed START", row, column)
	}
}
