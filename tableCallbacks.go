package gui

//
// These functions are the hooks to the andlabs libui
// They eventually feed information to the OS native GUI toolkits
// and feed back user interaction with the GUI
//

import "log"
import "fmt"
import "image/color"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/gookit/config"

var CurrentVM string

func (mh *TableData) NumRows(m *ui.TableModel) int {
	if (config.String("debugging") == "true") {
		log.Println("NumRows = mh.RowCount = ", mh.RowCount)
	}
	return mh.RowCount
}

// FYI: this routine seems to be called around 10 to 100 times a second for each table
func (mh *TableData) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	if (config.String("debugging") == "true") {
		log.Println("ColumnTypes = ", mh.generatedColumnTypes)
	}
	return mh.generatedColumnTypes
}

func libuiColorToGOlangColor(rgba color.RGBA) ui.TableColor {
	return ui.TableColor{float64(rgba.R) / 256, float64(rgba.G) / 256, float64(rgba.B) / 256, float64(rgba.A) / 256}
}

// TODO: Figure out why this is being called 1000 times a second (10 times for each row & column)
// Nevermind this TODO. Who gives a shit. This is a really smart way to treat the OS toolkits
func (mh *TableData) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if (config.String("debugging") == "true") {
		log.Println("CellValue() row, column =", row, column)
	}
	humanID := mh.Cells[column].HumanID
	if (column == mh.Human[humanID].TextID) {
		return ui.TableString(mh.Rows[row].HumanData[humanID].Text)
	}
	if (column == mh.Human[humanID].ColorID) {
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
		vmname := mh.Rows[row].HumanData[humanID].Text
		log.Println("vmname =",  vmname)
		log.Println("defaultSetCellValue() FOUND THE BUTTON!!!!!!!   Button was pressed START", row, column)
		CurrentVM = fmt.Sprintf("%s",vmname)
		log.Println("CurrentVM =", CurrentVM)
		go ui.Main(ShowVM)
	}
}
