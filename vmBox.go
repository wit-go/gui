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
