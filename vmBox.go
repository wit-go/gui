package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

func CreateVmBox(gw *GuiWindow, vm *pb.Event_VM) {
	log.Println("CreateVmBox() START")
	log.Println("CreateVmBox() vm.Name =", vm.Name)
	log.Println("CreateVmBox() gw =", gw)

	var box *GuiBox
	box = new(GuiBox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	log.Println("CreateVmBox() vbox =", vbox)
	log.Println("CreateVmBox() box.UiBox =", box.UiBox)
	box.UiBox = vbox
	log.Println("CreateVmBox() box.W =", box.W)
	box.W = gw
	log.Println("CreateVmBox() gw.BoxMap =", gw.BoxMap)
	gw.BoxMap[vm.Name] = box

	if (Data.Debug) {
		spew.Dump(vm)
	}

	hboxAccount := ui.NewHorizontalBox()
	hboxAccount.SetPadded(true)
	vbox.Append(hboxAccount, false)

	// Add hostname entry box
	makeEntryVbox(hboxAccount, "hostname:",	vm.Hostname,			true, "Hostname")
	makeEntryVbox(hboxAccount, "IPv6:",	vm.IPv6,			true, "IPv6")
	makeEntryVbox(hboxAccount, "RAM:",	fmt.Sprintf("%d",vm.Memory),	true, "Memory")
	makeEntryVbox(hboxAccount, "CPU:",	fmt.Sprintf("%d",vm.Cpus),	true, "Cpus")
	makeEntryVbox(hboxAccount, "Disk (GB):",fmt.Sprintf("%d",vm.Disk),	true, "Disk")
	makeEntryVbox(hboxAccount, "OS Image:",	vm.BaseImage,			true, "BaseImage")

	vbox.Append(ui.NewHorizontalSeparator(), false)

	hboxButtons := ui.NewHorizontalBox()
	hboxButtons.SetPadded(true)
	vbox.Append(hboxButtons, false)

	a := CreateButton(box, nil, vm, "Power On",  "POWERON",  nil)
	hboxButtons.Append(a.B, false)
	a = CreateButton(box, nil, vm, "Power Off", "POWEROFF", nil)
	hboxButtons.Append(a.B, false)
	a = CreateButton(box, nil, vm, "Destroy",   "DESTROY",  nil)
	hboxButtons.Append(a.B, false)
	a = CreateButton(box, nil, vm, "ping",      "PING",     runPingClick)
	hboxButtons.Append(a.B, false)
	a = CreateButton(box, nil, vm, "Console",   "XTERM",    runTestExecClick)
	hboxButtons.Append(a.B, false)
	a = CreateButton(box, nil, vm, "Save",      "SAVE",     nil)
	hboxButtons.Append(a.B, false)
	a = CreateButton(box, nil, vm, "Done",      "DONE",     nil)
	hboxButtons.Append(a.B, false)

	AddBoxToTab(vm.Name, gw.UiTab, vbox)
}

func createAddVmBox(gw *GuiWindow, b *GuiButton) {
	log.Println("createAddVmBox() START")
	name := "(" + b.Account.Nick + ")"

	var box *GuiBox
	box = new(GuiBox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	box.UiBox = vbox
	box.W = gw
	gw.BoxMap["ADD VM" + name] = box

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	// Add hostname entry box
	hostname := makeEntryHbox(vbox, "Hostname:", "testhost", true, "Hostname")
	memory   := makeEntryHbox(vbox, "Memory:",   "512", true, "Memory")
	disk     := makeEntryHbox(vbox, "Disk:",     "20", true, "Disk")

	log.Println("createAddVmBox() hostname, memory, disk =", hostname, memory, disk)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	hboxButtons := ui.NewHorizontalBox()
	hboxButtons.SetPadded(true)
	vbox.Append(hboxButtons, false)

	var newb GuiButton
	newb.Action	= "CREATE"
	newb.VM		= b.VM
	newb.Account	= b.Account
	hostname.B	= &newb
	memory.B	= &newb
	disk.B		= &newb
	hboxButtons.Append(AddButton(&newb, "Add Virtual Machine"), false)

	a := CreateButton(box, nil, nil, "Cancel",		"CLOSE", nil)
	hboxButtons.Append(a.B, false)

	AddBoxToTab(name, gw.UiTab, vbox)
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

	mh := AddTableTab(gw, name, count, parts, a)
//	mh := 
	return mh
}
