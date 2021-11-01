package gui

import "log"
// import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"
// import "github.com/davecgh/go-spew/spew"

// functions for handling text entry boxes

func (n *Node) NewLabel(text string) *Node {
	// make new node here
	// n.Append(ui.NewLabel(text), false)
	newNode := makeNode(n, text, 333, 334)
	newNode.Dump()
	// panic("node.NewLabel()")

	n.Append(newNode)
	return newNode
}

func (n *Node) SetText(value string) error {
	log.Println("gui.SetText() value =", value)
	if (n.uiText != nil) {
		n.uiText.SetText(value)
		return nil
	}
	if (n.uiButton != nil) {
		n.uiButton.SetText(value)
		return nil
	}
	return nil
}

func defaultEntryChange(e *ui.Entry) {
	for key, em := range Data.AllEntries {
		if (Config.Debug) {
			log.Println("\tdefaultEntryChange() Data.AllEntries =", key, em)
		}
		if Data.AllEntries[key].UiEntry == e {
			log.Println("defaultEntryChange() FOUND", 
				"Name =", Data.AllEntries[key].Name,
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
	newEntry.Name     = action
	if (action == "INT") {
		newEntry.Normalize = normalizeInt
	}
	Data.AllEntries = append(Data.AllEntries, &newEntry)

	return &newEntry
}
