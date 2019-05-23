package gui

import "log"
import "time"
import "fmt"
import "runtime/debug"
import "runtime"

import "github.com/gookit/config"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

func makeCloudInfoBox(custom func(int, string)) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewColorButton(), false)

	addXbutton := CreateButton("Show bmath's Account", "BMATH", custom)
	vbox.Append(addXbutton, false)

	vbox.Append(CreateButton("Add Account", "ADD", custom), false)

	vbox.Append(CreateButton("Quit", "QUIT", custom), false)

	if (config.String("debugging") == "true") {
		addDebuggingButtons(vbox, custom)
	}

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	/*
	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)
	*/

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	vbox.Append(entryForm, false)

	hostnameEntry :=  ui.NewEntry()
	entryForm.Append("hostname:", hostnameEntry, false)
	hostnameEntry.SetText("librem15.lab.wit.com")

	IPv6entry :=  ui.NewEntry()
	entryForm.Append("IPv6:", IPv6entry, false)
	IPv6entry.SetText("2604:bbc0:3:3:0:10:0:1003")

	vbox.Append(ui.NewHorizontalSeparator(), false)

	vbox.Append(ui.NewLabel("Accounts:"), false)

	agrid := ui.NewGrid()
	agrid.SetPadded(true)

	for account, _ := range config.StringMap("accounts") {
		hostname	:= config.String("accounts." + account + ".hostname")
		domainname	:= config.String("accounts." + account + ".domainname")
		port		:= config.String("accounts." + account + ".port")
		username	:= config.String("accounts." + account + ".username")
		a := account + " " + hostname + " " + domainname + " " + port + " " + username
		log.Println("ACCOUNT: ", a)
		vbox.Append(ui.NewLabel(a), false)
	}
	return hbox
}

// can not pass any args to this (?)
func setupCloudUI() {
	Data.cloudWindow = ui.NewWindow("Cloud Control Panel", config.Int("width"), config.Int("height"), false)
	Data.cloudWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		Data.cloudWindow.Destroy()
		return true
	})

	Data.cloudTab = ui.NewTab()
	Data.cloudWindow.SetChild(Data.cloudTab)
	Data.cloudWindow.SetMargined(true)

	Data.tabcount = 0
	Data.cloudTab.Append("Cloud Info", makeCloudInfoBox(nil))
	Data.cloudTab.SetMargined(Data.tabcount, true)

	Data.cloudWindow.Show()
}

func addTableTab() {
	var parts []TableColumnData

	for key, foo := range []string{"BG", "TEXTCOLOR", "BUTTON", "TEXTCOLOR", "TEXTCOLOR", "TEXT", "BUTTON", "TEXT", "BUTTON"} {
		log.Println(key, foo)

		var b TableColumnData
		b.CellType = foo
		b.Heading  = fmt.Sprintf("heading%d", key)
		parts = append(parts, b)
	}

	log.Println("Sleep for 2 seconds, then try to add new tabs")
	time.Sleep(1 * 1000 * 1000 * 1000)
	AddTableTab(Data.cloudTab, 1, "test seven", 7, parts)
}

func AddVmsTab(count int) *TableData {
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
	tmp.Heading     = "Details"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	mh := AddTableTab(Data.cloudTab, 1, "Virtual Machines", count, parts)
	return mh
}

func addDebuggingButtons(vbox *ui.Box, custom func(int, string)) {
	vbox.Append(ui.NewLabel("Debugging:"), false)

	// ATTEMPT TO ADD THE TABLE HERE
	add2button := ui.NewButton("Add a Test Table")
	add2button.OnClicked(func(*ui.Button) {
		log.Println("send over socket")
		addTableTab()
	})
	vbox.Append(add2button, false)
	// ATTEMPT TO ADD THE TABLE HERE END

	// hbox.Append(ui.NewVerticalSeparator(), false)

	// Send a test protobuf Event to localhost
	add3button := CreateButton("Add buf to chan", "ADD CHAN BUF", custom)
/*
	add3button := ui.NewButton("Add buf to chann")
	add3button.OnClicked(func(*ui.Button) {
		log.Println("add protobuf event to the channel")
		addSampleEvent()
	})
*/
	vbox.Append(add3button, false)

	add4button := CreateButton("Add Demo Event", "ADD DEMO EVENT", custom)
/*
	add4button := ui.NewButton("Add Demo Event")
	add4button.OnClicked(func(*ui.Button) {
		log.Println("add demo protobuf event to the channel")
		msg := pb.CreateSampleEvent()
		msg.Name = "generated in addSampleEvent()"
		msg.Type = pb.Event_DEMO
		addEvent(msg)
	})
*/
	vbox.Append(add4button, false)

	add4abutton := ui.NewButton("Close Demo GUI")
	add4abutton.OnClicked(func(*ui.Button) {
		CloseDemoUI()
	})
	vbox.Append(add4abutton, false)

	// Send a protobuf Event over the WIT socket
	add5button := CreateButton("Send protobuf to localhost", "SEND PROTOBUF TO LOCALHOST", custom)
/*
	add5button := ui.NewButton("Send protobuf to localhost")
	add5button.OnClicked(func(*ui.Button) {
		log.Println("sent a Marshal'd protobuf to a localhost socket")
		sendDataToDest()
	})
*/
	vbox.Append(add5button, false)

	// Send a protobuf over a gorilla websocket
	add6button := CreateButton("gorillaSendProtobuf()", "SEND PROTOBUF TO GORILLA SOCKET", custom)
/*
	add6button := ui.NewButton("gorillaSendProtobuf()")
	add6button.OnClicked(func(*ui.Button) {
		log.Println("gorillaSendProtobuf()")
		gorillaSendProtobuf()
	})
*/
	vbox.Append(add6button, false)

	// debug all the golang goroutines
	add7button := ui.NewButton("debug.PrintStack()")
	add7button.OnClicked(func(*ui.Button) {
		log.Println("debug.PrintStack() (SHOULD BE JUST THIS goroutine)")
		debug.PrintStack()

		log.Println("ATTEMPT FULL STACK DUMP")
		log.Println("ATTEMPT FULL STACK DUMP")
		log.Println("ATTEMPT FULL STACK DUMP")
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		fmt.Printf("%s", buf)
	})
	vbox.Append(add7button, false)
}
