package gui

//
// These functions are the hooks to the andlabs libui
// They eventually feed information to the OS native GUI toolkits
// and feed back user interaction with the GUI
//

import "log"
// import "fmt"
import "image/color"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (mh *TableData) NumRows(m *ui.TableModel) int {
	if (Config.Debug) {
		log.Println("NumRows = mh.RowCount = ", mh.RowCount, "(last Row & Column =", mh.lastRow, mh.lastColumn, ")")
	}
	return mh.RowCount
}

// FYI: this routine seems to be called around 10 to 100 times a second for each table
func (mh *TableData) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	if (Config.DebugTable) {
		log.Println("ColumnTypes = ", mh.generatedColumnTypes)
	}
	return mh.generatedColumnTypes
}

// TODO: Figure out why this is being called 1000 times a second (10 times for each row & column)
//
// Nevermind that TODO. Who gives a shit. This is a really smart way to treat the OS toolkits
func (mh *TableData) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if (Config.DebugTable) {
		log.Println("CellValue() row, column =", row, column)
	}
	mh.lastRow = row
	mh.lastColumn = column
	humanID := mh.Cells[column].HumanID
	if (column == mh.Human[humanID].TextID) {
		return ui.TableString(mh.Rows[row].HumanData[humanID].Text)
	}
	if (column == mh.Human[humanID].ColorID) {
		if (column == 0) {
			// ignore BG color on windows for now
			if (runtime.GOOS == "windows") {
				// TODO: fix colors on windows
				// log.Println("CellValue() WINDOWS is BG COLOR row, column =", row, column)
				// return nil
			}

			if (mh.Rows[row].HumanData[humanID].Color == color.RGBA{255, 255, 255, 255}) {
				log.Println("CellValue() color.RGBA{255, 255, 255, 255} so return nil. row, column =", row, column)
				return nil
			}

			bgcolor := libuiColorToGOlangColor(mh.Rows[row].HumanData[humanID].Color)
			if (Config.Debug) {
				log.Println("CellValue() BGCOLOR =", bgcolor)
			}
			return bgcolor
		}
		return libuiColorToGOlangColor(mh.Rows[row].HumanData[humanID].Color)
	}
	log.Println("CellValue() FAILURE")
	log.Println("CellValue() FAILURE")
	log.Println("CellValue() row, column = ", row, column)
	log.Println("CellValue() FAILURE")
	log.Println("CellValue() FAILURE")
	log.Println("CellValue() mh.Cells", mh.Cells)
	log.Println("CellValue() mh.Human", mh.Human)
	panic("CellValue() not sure what sort of ui.TableValue to return in CellValue()")
	return ui.TableString("")
}

func (mh *TableData) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	log.Println("SetCellValue() START row=", row, "column=", column, "value=", value)

	defaultSetCellValue(mh, row, column)

	log.Println("mh.Cells[column].HumanID =", mh.Cells[column].HumanID)
	// log.Println("mh.Rows[row].Cells[column].HumanID =", mh.Rows[row].Cells[column].HumanID)

	humanID := mh.Cells[column].HumanID
	log.Println("mh.Human[humanID].ColorID =", mh.Human[humanID].ColorID)
	log.Println("mh.Human[humanID].TextID =",  mh.Human[humanID].TextID)

	log.Println("SetCellValue() END")
}

func defaultSetCellValue(mh *TableData, row int, column int) {
	if (mh.Cells[column].Name == "BUTTON") {
		humanID := mh.Cells[column].HumanID
		log.Println("defaultSetCellValue() FOUND THE TABLE BUTTON ", row, humanID)

		n := mh.Rows[row].HumanData[humanID].N
		if (n != nil) {
			// TODO: fixme. removed on Oct 31 2021
			if (n.OnChanged != nil) {
				n.OnChanged()
			}
			return
		}
		log.Println("defaultSetCellValue() ERROR: UNKNOWN BUTTON IN TABLE")
		if (Config.Debug) {
			panic("defaultSetCellValue() GOT AN UNKNOWN BUTTON CLICK IN TABLE")
		}
	}
}
