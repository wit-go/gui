package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"
import "github.com/davecgh/go-spew/spew"
// import pb "git.wit.com/wit/witProtobuf"

// THIS IS CLEAN (except the 'Memory' normalization example)

// functions for handling text entry boxes

func NewLabel(box *GuiBox, text string) {
	box.UiBox.Append(ui.NewLabel(text), false)
}

func GetText(box *GuiBox, name string) string {
	if (box == nil) {
		log.Println("gui.GetText() ERROR box == nil")
		return ""
	}
	if (box.Window.EntryMap == nil) {
		log.Println("gui.GetText() ERROR b.Box.Window.EntryMap == nil")
		return ""
	}
	spew.Dump(box.Window.EntryMap)
	if (box.Window.EntryMap[name] == nil) {
		log.Println("gui.GetText() ERROR box.Window.EntryMap[", name, "] == nil ")
		return ""
	}
	e := box.Window.EntryMap[name]
	log.Println("gui.GetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
	log.Println("gui.GetText() END")
	return e.UiEntry.Text()
}

func SetText(box *GuiBox, name string, value string) error {
	if (box == nil) {
		return fmt.Errorf("gui.SetText() ERROR box == nil")
	}
	if (box.Window.EntryMap == nil) {
		return fmt.Errorf("gui.SetText() ERROR b.Box.Window.EntryMap == nil")
	}
	spew.Dump(box.Window.EntryMap)
	if (box.Window.EntryMap[name] == nil) {
		return fmt.Errorf("gui.SetText() ERROR box.Window.EntryMap[", name, "] == nil ")
	}
	e := box.Window.EntryMap[name]
	log.Println("gui.SetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
	e.UiEntry.SetText(value)
	log.Println("gui.SetText() box.Window.EntryMap[", name, "] = ", e.UiEntry.Text())
	log.Println("gui.SetText() END")
	return nil
}

// makeEntryBox(box, "hostname:", "blah.foo.org") {
func MakeEntryVbox(box *GuiBox, a string, startValue string, edit bool, action string) *GuiEntry {
	// Start 'Nickname' vertical box
	vboxN := ui.NewVerticalBox()
	vboxN.SetPadded(true)
	vboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)

	vboxN.Append(e.UiEntry, false)
	box.UiBox.Append(vboxN, false)
	// End 'Nickname' vertical box

	return e
}

func MakeEntryHbox(box *GuiBox, a string, startValue string, edit bool, action string) *GuiEntry {
	// Start 'Nickname' vertical box
	hboxN := ui.NewHorizontalBox()
	hboxN.SetPadded(true)
	hboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)
	hboxN.Append(e.UiEntry, false)

	box.UiBox.Append(hboxN, false)
	// End 'Nickname' vertical box

	return e
}

func AddEntry(box *GuiBox, name string) *GuiEntry {
	var ge *GuiEntry
	ge = new(GuiEntry)

	ue := ui.NewEntry()
	ue.SetReadOnly(false)
	ue.OnChanged(func(*ui.Entry) {
		log.Println("gui.AddEntry() OK. ue.Text() =", ue.Text())
	})
	box.UiBox.Append(ue, false)

	ge.UiEntry = ue
	box.Window.EntryMap[name] = ge

	return ge
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
