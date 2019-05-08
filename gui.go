package gui

import "log"
// import "fmt"

import "github.com/gookit/config"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

var mainwin *ui.Window
var maintab *ui.Tab
var tabcount int

var jcarrButton *ui.Button
var jcarrEntry  *ui.MultilineEntry

func buttonClick(button *ui.Button) {
	log.Println("hostname =", config.String("hostname"), button)
	spew.Dump(button)
	if (jcarrButton == button) {
		log.Println("This is the jcarrButton")
		cur := jcarrEntry.Text()
		jcarrEntry.SetText(cur + "THIS IS A GREAT IDEA\n")
	} else {
		log.Println("This is NOT the jcarrButton")
	}
}

func hostnameButton(hostname string) ui.Control {
	tmpbox := ui.NewHorizontalBox()
	tmpbox.SetPadded(true)

	tmpButton := ui.NewButton(hostname)
	tmpbox.Append(tmpButton, false)
	tmpButton.OnClicked(buttonClick)

	jcarrButton = tmpButton

	return tmpbox
}

func makeGroupEntries() ui.Control {
	group := ui.NewGroup("Entries")
	group.SetMargined(true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	jcarrEntry = ui.NewMultilineEntry()
	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", jcarrEntry, true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	return group
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(47, 100)
	slider  := ui.NewSlider(21, 100)
	pbar    := ui.NewProgressBar()

	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)
	vbox.Append(hostnameButton("jcarrtest"), false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainwin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainwin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("Cloud Control Panel", config.Int("width"), config.Int("height"), false)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	maintab = ui.NewTab()
	mainwin.SetChild(maintab)
	mainwin.SetMargined(true)

	// maintab.Append("List examples", makeNumbersPage())
	tabcount = 0
	// maintab.SetMargined(tabcount, true)

/*
	maintab.Append("Choosers examples", makeDataChoosersPage())
	tabcount += 1
	maintab.SetMargined(tabcount, true)

	maintab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	maintab.SetMargined(tabcount, true)
*/

	mainwin.Show()
}

func AddChoosersDemo() {
	maintab.Append("Choosers examples", makeDataChoosersPage())
	maintab.SetMargined(tabcount, true)
	tabcount += 1
}

// This hangs on GTK
func AddEntriesDemo() {
	maintab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	maintab.SetMargined(tabcount, true)
}

func initRow(mh *tableData, row int) {
	initRowBTcolor        (mh, row, 0)
	initRowTextColorColumn(mh, row, 1, 2, "diff1", ui.TableColor{0.9, 0, 0, 1})
	initRowButtonColumn   (mh, row, 3,    "diff2")
	initRowTextColorColumn(mh, row, 4, 5, "diff3", ui.TableColor{0.0, 0.9, 0.4, 1})
	initRowTextColorColumn(mh, row, 6, 7, "diff4", ui.TableColor{0.3, 0.1, 0.9, 1})
	initRowTextColumn     (mh, row, 8,    "diff5")
	initRowButtonColumn   (mh, row, 9,    "diff6")
}

func AddTableTab(name string, rowcount int, row1name string) {
	mh := new(tableData)

	mh.rowcount    = rowcount
	mh.rows        = make([]rowData, mh.rowcount)

	// This is the standard callback function from libUI when the user does something
	mh.libUIevent      = defaultSetCellValue

	tmpBTindex := 0
	initBTcolor        (mh, 0)
	initTextColorColumn(mh, 1, 2, "really fun", ui.TableColor{0.9, 0, 0, 1})
	initButtonColumn   (mh, 3,    "but3ton")
	initTextColorColumn(mh, 4, 5, "jwc45blah", ui.TableColor{0.0, 0.9, 0.4, 1})
	initTextColorColumn(mh, 6, 7, "jwc67blah", ui.TableColor{0.3, 0.1, 0.9, 1})
	initTextColumn     (mh, 8,    "jwc8blah")
	initButtonColumn   (mh, 9,    "but9ton")

	// spew.Dump(mh)
	log.Println(mh)

	b := make([]rowData, 5)
	mh.rows = append(mh.rows, b...)

	initRow(mh, mh.rowcount)
	mh.rowcount    = rowcount + 1

	log.Println(mh)

/*
	mh.rowcount    = rowcount
	initRow(mh, mh.rowcount)

	spew.Dump(mh)
*/

	model := ui.NewTableModel(mh)

	table := ui.NewTable(
		&ui.TableParams{
			Model:	model,
			RowBackgroundColorModelColumn:	tmpBTindex,
	})

	appendTextColumn        (mh, table, 1, "jwc1col")
	table.AppendButtonColumn("button3", 3, ui.TableModelColumnAlwaysEditable)
	appendTextColorColumn   (mh, table, 4, 5, "testcolor")
	appendTextColorColumn   (mh, table, 6, 7, "appendTest")
	appendTextColumn        (mh, table, 8, "jwc8col")
	table.AppendButtonColumn("button9", 9, ui.TableModelColumnAlwaysEditable)

	maintab.Append(name, table)
	maintab.SetMargined(tabcount, true)
	tabcount += 1
}

func DoGUI() {
	ui.Main(setupUI)

	log.Println("GUI exited. Not sure what to do here. os.Exit() ?")

	// not sure how to pass this back to the main program
	// onExit()
}
