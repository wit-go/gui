package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

// THIS IS NOT CLEAN

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

	box.UiBox = hboxAccount

	// Add hostname entry box
	MakeEntryVbox(box, "hostname:",	vm.Hostname,			true, "Hostname")
	MakeEntryVbox(box, "IPv6:",	vm.IPv6,			true, "IPv6")
	MakeEntryVbox(box, "RAM:",	fmt.Sprintf("%d",vm.Memory),	true, "Memory")
	MakeEntryVbox(box, "CPU:",	fmt.Sprintf("%d",vm.Cpus),	true, "Cpus")
	MakeEntryVbox(box, "Disk (GB):",fmt.Sprintf("%d",vm.Disk),	true, "Disk")
	MakeEntryVbox(box, "OS Image:",	vm.BaseImage,			true, "BaseImage")

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

func CreateAddVmBox(gw *GuiWindow, b *GuiButton, name string) *GuiBox{
	log.Println("CreateAddVmBox() START name =", name)

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

	// abox := gw.MakeTab(gw)

	AddBoxToTab(name, gw.UiTab, vbox)

	return box

	/*
	// Add hostname entry box
	hostname := MakeEntryHbox(box, "Hostname:", "testhost", true, "Hostname")
	memory   := MakeEntryHbox(box, "Memory:",   "512", true, "Memory")
	disk     := MakeEntryHbox(box, "Disk:",     "20", true, "Disk")

	log.Println("CreateAddVmBox() hostname, memory, disk =", hostname, memory, disk)

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
	*/
}
