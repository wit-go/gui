package gui

//
// These functions are the hooks to the andlabs libui
// They eventually feed information to the OS native GUI toolkits
// and feed back user interaction with the GUI
//

import "os"
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
	return mh.Rows[row].Cells[column].Value
}

func (mh *TableData) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	log.Println("SetCellValue() START row=", row, "column=", column, "value=", value)

	if (mh.libUIevent == nil) {
		log.Println("CellValue NOT DEFINED. This table wasn't setup correctly! mh.scanCellValue == nil")
		os.Exit(-1)
	}

	mh.libUIevent(mh, m, row, column, value)
	if (mh.cellChangeEvent != nil) {
		mh.cellChangeEvent(row, column, value)
	}

	log.Println("mh.Rows[0].Cells[column].HumanID =", mh.Rows[0].Cells[column].HumanID)
	log.Println("mh.Rows[row].Cells[column].HumanID =", mh.Rows[row].Cells[column].HumanID)

	humanID := mh.Rows[row].Cells[column].HumanID
	log.Println("mh.Rows[row].Human[humanID].ColorID =", mh.Rows[row].Human[humanID].ColorID)
	log.Println("mh.Rows[row].Human[humanID].TextID =",  mh.Rows[row].Human[humanID].TextID)

	if (column == mh.Rows[row].Human[humanID].TextID) {
		log.Println("THIS COLUMN IS A TEXT COLUMN")
		mh.Rows[row].Cells[column].Value = mh.Rows[row].Human[humanID].Text
	}

	log.Println("SetCellValue() END")
}
