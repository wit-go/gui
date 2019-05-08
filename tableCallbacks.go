package main

//
// These functions are the hooks to the andlabs libui
// They eventually feed information to the OS native GUI toolkits
// and feed back user interaction with the GUI
//

import "os"
import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (mh *tableData) NumRows(m *ui.TableModel) int {
	return mh.rowcount
}

// FYI: this routine seems to be called around 10 to 100 times a second for each table
func (mh *tableData) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return mh.generatedColumnTypes
}

// TODO: Figure out why this is being called 1000 times a second (10 times for each row & column)
// Nevermind this TODO. Who gives a shit. This is a really smart way to treat the OS toolkits
func (mh *tableData) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	return mh.rows[row].cells[column].value
}

func (mh *tableData) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	log.Println("SetCallValue() START row=", row, "column=", column, "value=", value)
	// spew.Dump(m)
	// spew.Dump(mh)
	if (mh.libUIevent == nil) {
		log.Println("CellValue NOT DEFINED. This table wasn't setup correctly! mh.scanCellValue == nil")
		os.Exit(-1)
	}
	// spew.Dump(m)
	mh.libUIevent(mh, m, row, column, value)
	log.Println("SetCallValue() END")
}
