package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

func GoShowVM() {
	ui.Main(ShowVM)
}

func ShowVM() {
	name := Data.CurrentVM.Name
	log.Println("ShowVM() START Data.CurrentVM=", Data.CurrentVM)
	VMwin := ui.NewWindow("VM " + name, 500, 300, false)

        // create a 'fake' button entry for the mouse clicks
	var newButtonMap ButtonMap
	newButtonMap.Action   = "WINDOW CLOSE"
	newButtonMap.W        = VMwin
	Data.AllButtons = append(Data.AllButtons, newButtonMap)

	VMwin.OnClosing(func(*ui.Window) bool {
		mouseClick(&newButtonMap)
		return true
	})
	ui.OnShouldQuit(func() bool {
		mouseClick(&newButtonMap)
		return true
	})

	VMtab := ui.NewTab()
	VMwin.SetChild(VMtab)
	VMwin.SetMargined(true)

	CreateVmBox(VMtab, Data.CurrentVM)
	VMwin.Show()
}

func AddVmConfigureTab(name string, pbVM *pb.Event_VM) {
	CreateVmBox(window1.T, Data.CurrentVM)
}

func CreateVmBox(tab *ui.Tab, vm *pb.Event_VM) {
	log.Println("CreateVmBox() START")
	log.Println("CreateVmBox() vm.Name", vm.Name)
	spew.Dump(vm)
	if (Data.Debug) {
		spew.Dump(vm)
	}
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

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

	hboxButtons.Append(CreateButton(nil, vm, "Power On",  "POWERON",  nil), false)
	hboxButtons.Append(CreateButton(nil, vm, "Power Off", "POWEROFF", nil), false)
	hboxButtons.Append(CreateButton(nil, vm, "Destroy",   "DESTROY",  nil), false)
	hboxButtons.Append(CreateButton(nil, vm, "ping",      "PING",     runPingClick), false)
	hboxButtons.Append(CreateButton(nil, vm, "Console",   "XTERM",    runTestExecClick), false)
	hboxButtons.Append(CreateButton(nil, vm, "Save",      "SAVE",     nil), false)
	hboxButtons.Append(CreateButton(nil, vm, "Done",      "DONE",     nil), false)

	AddBoxToTab(Data.CurrentVM.Name, tab, vbox)
}

func createAddVmBox(tab *ui.Tab, name string, b *ButtonMap) {
	log.Println("createAddVmBox() START")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

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

	var newb ButtonMap
	newb.Action	= "CREATE"
	newb.VM		= b.VM
	newb.Account	= b.Account
	newb.T		= tab
	hostname.B	= &newb
	memory.B	= &newb
	disk.B		= &newb
	hboxButtons.Append(AddButton(&newb, "Add Virtual Machine"), false)

	// hboxButtons.Append(CreateButton(nil, nil, "Add Virtual Machine","CREATE",nil), false)
	hboxButtons.Append(CreateButton(nil, nil, "Cancel",		"CLOSE", nil), false)

	name += " (" + b.Account.Nick + ")"
	AddBoxToTab(name, tab, vbox)
}
