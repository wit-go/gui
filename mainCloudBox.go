package gui

import "log"
import "time"
import "regexp"
import "os"
// import "reflect"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

func makeCloudInfoBox(gw *GuiWindow) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	if (Data.Debug) {
		log.Println("makeCloudInfoBox() add debugging buttons")
		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		hbox.Append(vbox, false)

		addDebuggingButtons(gw, vbox)

		hbox.Append(ui.NewVerticalSeparator(), false)
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	hostnamebox := ui.NewHorizontalBox()
	hostnamebox.SetPadded(true)
	vbox.Append(hostnamebox, false)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	hostnamebox.Append(entryForm, true)

	hostnameEntry :=  ui.NewEntry()
	entryForm.Append("hostname:", hostnameEntry, false)
	tmp := Data.Hostname + " (" + Data.IPv6 + ")"
	hostnameEntry.SetText(tmp)
	hostnameEntry.SetReadOnly(true)

	anew := CreateButton(gw, nil, nil, "Edit", "EDIT", nil)
	hostnamebox.Append(anew.B, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	agrid := ui.NewGrid()
	agrid.SetPadded(true)

	agrid.Append(ui.NewLabel("Accounts:"),   0, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Domain Name"), 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Email"),       2, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	row := 1

	for key, a := range Data.Config.Accounts {
		log.Println("account =          ", key, a)
		log.Println("Accounts[key] =    ", Data.Config.Accounts[key])
		log.Println("account.Nick =     ", Data.Config.Accounts[key].Nick)
		log.Println("account.Username = ", Data.Config.Accounts[key].Username)
		log.Println("account.Token =    ", Data.Config.Accounts[key].Token)

		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Domain),	1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Email),	2, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name := "Login " + Data.Config.Accounts[key].Nick
		l := CreateButton(gw, Data.Config.Accounts[key], nil, name, "LOGIN", nil)
		agrid.Append(l.B, 3, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name  = "Show " + Data.Config.Accounts[key].Nick
		b := CreateButton(gw, Data.Config.Accounts[key], nil, name, "SHOW", nil)
		agrid.Append(b.B, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		row += 1
	}

	row += 1
	agrid.Append(ui.NewLabel(""),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	row += 1
	a := CreateButton(gw, nil, nil, "Add Account", "ADD TAB", nil)
	agrid.Append(a.B, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	q := CreateButton(gw, nil, nil, "Quit", "QUIT", nil)
	agrid.Append(q.B, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(agrid, false)
	return hbox
}

//
// THIS IS THE STANDARD VM DISPLAY TABLE
// This maps the 'human' indexed cells in the table
// to the machine's andlabs/libui values. That way
// if you want to work against column 4, then you
// can just reference 4 instead of the internal number
// which could be anything since TEXTCOLOR, TEXT, BG, etc
// fields use between 1 and 3 values internally
//
func AddVmsTab(gw *GuiWindow, name string, count int, a *pb.Account) *TableData {
	var parts []TableColumnData

	human := 0

	tmp := TableColumnData{}
	tmp.CellType = "BG"
	tmp.Heading  = "background"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "name"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "hostname"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "IPv6"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "cpus"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "memory"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "BUTTON"
	tmp.Heading  = "Details"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	mh := AddTableTab(gw, gw.T, 1, name, count, parts, a)
	return mh
}

func ShowAccountQuestionTab(gw *GuiWindow) {
	log.Println("ShowAccountQuestionTab() gw =", gw)
	if (gw.T == nil) {
		log.Println("ShowAccountQuestionTab() gw.T = nil THIS IS BAD")
		os.Exit(-1)
	}
	gw.T.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	gw.Box2 = AddAccountQuestionBox(gw)
	gw.T.InsertAt("New Account?", 0, gw.Box2)
	gw.T.SetMargined(0, true)
}

func ShowAccountTab(gw *GuiWindow, i int) {
	log.Println("ShowAccountTab() START")

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	// Create the things for the Account Tab
	abox := AddAccountBox(gw)

	// Set the parents and data structure links
	// aTab.me = gw.T
	// aTab.parentWindow = Data.Window1.W
	// aTab.tabOffset = 0

	if (i >= 0) {
		log.Println("ShowAccountTab() InsertAt i=", i)
		gw.T.Delete(0)
		gw.T.InsertAt("Add Account", i, abox)
		gw.T.SetMargined(0, true)
	} else {
		// TODO: After append try to discover the tab index #
		log.Println("ShowAccountTab() Append")
		AddBoxToTab("Create New Account", gw.T, abox)
	}
}

func ShowMainTab(gw *GuiWindow) {
	log.Println("ShowMainTab() gw =", gw)
	log.Println("ShowMainTab() gw.T =", gw.T)
	gw.T.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	gw.Box2 = makeCloudInfoBox(gw)
	gw.T.InsertAt("Main", 0, gw.Box2)
	gw.T.SetMargined(0, true)
}

func GuiInit() {
	ui.OnShouldQuit(func() bool {
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})
}

func StartNewWindow(c *pb.Config, bg bool, action string) {
	log.Println("InitNewWindow() Create a new window")
	var newGuiWindow GuiWindow
	newGuiWindow.C = c
	newGuiWindow.Action = action
	Data.Windows = append(Data.Windows, &newGuiWindow)

	if (bg) {
		log.Println("ShowWindow() IN NEW GOROUTINE")
		go ui.Main(func() {
			InitWindow(&newGuiWindow)
		})
		time.Sleep(2000 * time.Millisecond)
	} else {
		log.Println("ShowWindow() WAITING for ui.Main()")
		ui.Main(func() {
			InitWindow(&newGuiWindow)
		})
	}
}

func getSplashText(a string) *ui.AttributedString {
	var aText  *ui.AttributedString
	aText = ui.NewAttributedString(a)
	return aText
}


func InitWindow(gw *GuiWindow) {
	log.Println("InitWindow() THIS WINDOW IS NOT YET SHOWN")

	c := gw.C
	gw.W = ui.NewWindow("", int(c.Width), int(c.Height), true)
	gw.W.SetBorderless(false)

        // create a 'fake' button entry for the mouse clicks
	var newBM GuiButton
	newBM.Action	= "QUIT"
	newBM.W		= gw.W
	newBM.WM	= gw
	Data.AllButtons = append(Data.AllButtons, &newBM)

	gw.W.OnClosing(func(*ui.Window) bool {
		log.Println("InitWindow() OnClosing() THIS WINDOW IS CLOSING gw=", gw)
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})

	gw.T = ui.NewTab()
	gw.W.SetChild(gw.T)
	gw.W.SetMargined(true)

	log.Println("InitWindow() gw =", gw)
	log.Println("InitWindow() gw.Action =", gw.Action)

	if (gw.Action == "SPLASH") {
		log.Println("InitWindow() TRYING SPLASH")
		damnit := "click" + string(c.Hostname)
		tmp := getSplashText(damnit)
		log.Println("InitWindow() TRYING SPLASH tmp =", tmp)
		gw.Box1 = ShowSplashBox(gw, tmp)

		gw.T.Append("WIT Splash", gw.Box1)
		gw.T.SetMargined(0, true)
	}

	Data.State = "splash"
	gw.W.Show()
}

// makeEntryBox(box, "hostname:", "blah.foo.org") {
func makeEntryVbox(hbox *ui.Box, a string, startValue string, edit bool, action string) *GuiEntry {
	// Start 'Nickname' vertical box
	vboxN := ui.NewVerticalBox()
	vboxN.SetPadded(true)
	vboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)

	vboxN.Append(e.E, false)
	hbox.Append(vboxN, false)
	// End 'Nickname' vertical box

	return e
}

/*
// string handling examples that might be helpful for normalizeInt()
isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

for _, username := range []string{"userone", "user2", "user-three"} {
    if !isAlpha(username) {
        fmt.Printf("%q is not valid\n", username)
    }
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

func alphaOnly(s string) bool {
   for _, char := range s {
      if !strings.Contains(alpha, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}
*/

func normalizeInt(s string) string {
	// reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Println("normalizeInt() regexp.Compile() ERROR =", err)
		return s
	}
	clean := reg.ReplaceAllString(s, "")
	log.Println("normalizeInt() s =", clean)
	return clean
}

func defaultEntryChange(e *ui.Entry) {
	for key, em := range Data.AllEntries {
		if (Data.Debug) {
			log.Println("\tdefaultEntryChange() Data.AllEntries =", key, em)
		}
		if Data.AllEntries[key].E == e {
			log.Println("defaultEntryChange() FOUND", 
				"action =", Data.AllEntries[key].Action,
				"Last =", Data.AllEntries[key].Last,
				"e.Text() =", e.Text())
			Data.AllEntries[key].Last = e.Text()
			if Data.AllEntries[key].Normalize != nil {
				fixed := Data.AllEntries[key].Normalize(e.Text())
				e.SetText(fixed)
			}
			return
		}
	}
	log.Println("defaultEntryChange() ERROR. MISSING ENTRY MAP. e.Text() =", e.Text())
}

func defaultMakeEntry(startValue string, edit bool, action string) *GuiEntry {
	e := ui.NewEntry()
	e.SetText(startValue)
	if (edit == false) {
		e.SetReadOnly(true)
	}
	e.OnChanged(defaultEntryChange)

	// add the entry field to the global map
	var newEntry GuiEntry
	newEntry.E      = e
	newEntry.Edit   = edit
	newEntry.Action = action
	if (action == "Memory") {
		newEntry.Normalize = normalizeInt
	}
	Data.AllEntries = append(Data.AllEntries, &newEntry)

	return &newEntry
}

func makeEntryHbox(hbox *ui.Box, a string, startValue string, edit bool, action string) *GuiEntry {
	// Start 'Nickname' vertical box
	hboxN := ui.NewHorizontalBox()
	hboxN.SetPadded(true)
	hboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)
	hboxN.Append(e.E, false)

	hbox.Append(hboxN, false)
	// End 'Nickname' vertical box

	return e
}

func AddBoxToTab(name string, tab *ui.Tab, box *ui.Box) {
	tab.Append(name, box)
	tab.SetMargined(0, true)
}
