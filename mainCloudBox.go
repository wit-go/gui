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

// THIS IS NOT CLEAN

func makeCloudInfoBox(gw *GuiWindow) *GuiBox {
	var gb *GuiBox
	gb = new(GuiBox)
	gb.W = gw

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	// gw.Box1 = hbox
	gb.UiBox = hbox

	if (Data.Debug) {
		log.Println("makeCloudInfoBox() add debugging buttons")
		/*
		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		hbox.Append(vbox, false)
		*/

		addDebuggingButtons(gb)

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

	anew := CreateButton(gb, nil, nil, "Edit", "EDIT", nil)
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
		l := CreateButton(gb, Data.Config.Accounts[key], nil, name, "LOGIN", nil)
		agrid.Append(l.B, 3, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name  = "Show " + Data.Config.Accounts[key].Nick
		b := CreateButton(gb, Data.Config.Accounts[key], nil, name, "SHOW", nil)
		agrid.Append(b.B, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		row += 1
	}

	row += 1
	agrid.Append(ui.NewLabel(""),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	row += 1
	a := CreateButton(gb, nil, nil, "Add Account", "ADD TAB", nil)
	agrid.Append(a.B, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	q := CreateButton(gb, nil, nil, "Quit", "QUIT", nil)
	agrid.Append(q.B, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(agrid, false)
	return gb
}

func ShowAccountQuestionTab(gw *GuiWindow) {
	log.Println("ShowAccountQuestionTab() gw =", gw)
	if (gw.UiTab == nil) {
		log.Println("ShowAccountQuestionTab() gw.UiTab = nil THIS IS BAD")
		os.Exit(-1)
	}
	gw.UiTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	// abox := AddAccountQuestionBox(gw)
	abox := gw.MakeTab(gw)
	gw.BoxMap["Box2"] = abox
	gw.UiTab.InsertAt("New Account?", 0, abox.UiBox)
	gw.UiTab.SetMargined(0, true)
}

func ShowMainTab(gw *GuiWindow) {
	log.Println("ShowMainTab() gw =", gw)
	log.Println("ShowMainTab() gw.UiTab =", gw.UiTab)
	gw.UiTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	abox := makeCloudInfoBox(gw)
	gw.BoxMap["Box3"] = abox
	gw.UiTab.InsertAt("Main", 0, abox.UiBox)
	gw.UiTab.SetMargined(0, true)
}

func GuiInit() {
	ui.OnShouldQuit(func() bool {
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})
}

func StartNewWindow(c *pb.Config, bg bool, action string, maketab func(*GuiWindow) *GuiBox) {
	log.Println("InitNewWindow() Create a new window")
	var newGuiWindow GuiWindow
	newGuiWindow.Width   = int(c.Width)
	newGuiWindow.Height  = int(c.Height)
	newGuiWindow.Action  = action
	newGuiWindow.MakeTab = maketab
	Data.Windows = append(Data.Windows, &newGuiWindow)

	// make(newGuiWindow.BoxMap)
	newGuiWindow.BoxMap = make(map[string]*GuiBox)

	if (bg) {
		log.Println("ShowWindow() IN NEW GOROUTINE")
		go ui.Main(func() {
			InitTabWindow(&newGuiWindow)
		})
		time.Sleep(2000 * time.Millisecond)
	} else {
		log.Println("ShowWindow() WAITING for ui.Main()")
		ui.Main(func() {
			InitTabWindow(&newGuiWindow)
		})
	}
}

func getSplashText(a string) *ui.AttributedString {
	var aText  *ui.AttributedString
	aText = ui.NewAttributedString(a)
	return aText
}

func InitTabWindow(gw *GuiWindow) {
	log.Println("InitTabWindow() THIS WINDOW IS NOT YET SHOWN")

	gw.UiWindow = ui.NewWindow("", int(gw.Width), int(gw.Height), true)
	gw.UiWindow.SetBorderless(false)

        // create a 'fake' button entry for the mouse clicks
	var newBM GuiButton
	newBM.Action	= "QUIT"
//	newBM.W		= gw.UiWindow
	newBM.GW	= gw
	Data.AllButtons = append(Data.AllButtons, &newBM)

	gw.UiWindow.OnClosing(func(*ui.Window) bool {
		log.Println("InitTabWindow() OnClosing() THIS WINDOW IS CLOSING gw=", gw)
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})

	gw.UiTab = ui.NewTab()
	gw.UiWindow.SetChild(gw.UiTab)
	gw.UiWindow.SetMargined(true)

	log.Println("InitTabWindow() gw =", gw)

	abox := gw.MakeTab(gw)

	gw.UiTab.Append("WIT Splash", abox.UiBox)
	gw.UiTab.SetMargined(0, true)

	Data.State = "splash"
	gw.UiWindow.Show()
}

func AddBoxToTab(name string, tab *ui.Tab, box *ui.Box) {
	tab.Append(name, box)
	tab.SetMargined(0, true)
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
		if Data.AllEntries[key].UiEntry == e {
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
	newEntry.UiEntry  = e
	newEntry.Edit     = edit
	newEntry.Action   = action
	if (action == "Memory") {
		newEntry.Normalize = normalizeInt
	}
	Data.AllEntries = append(Data.AllEntries, &newEntry)

	return &newEntry
}
