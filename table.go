// based off andlabs/ui/examples/table.go

package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

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

func InitColumns(mh *TableData, parts []TableColumnData) {
	tmpBTindex := 0
	humanID := 0
	for key, foo := range parts {
		log.Println("key, foo =", key, foo)

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

func AddTableTab(gw *GuiWindow, name string, rowcount int, parts []TableColumnData) *TableData {
	node := InitWindow(nil, gw, name, Yaxis)

	return AddTableBox(node.box, name, rowcount, parts)
}

func AddTableBox(box *GuiBox, name string, rowcount int, parts []TableColumnData) *TableData {
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

	// is this needed?
	// gw.BoxMap[name] = box
	mh.Box = box

	box.UiBox.Append(table, true)

	return mh
}
