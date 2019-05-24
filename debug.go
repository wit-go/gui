package gui

import "log"
import "time"
import "fmt"
import "strings"
import "os/exec"

import "github.com/gookit/config"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

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

func addDebuggingButtons(vbox *ui.Box, custom func(*ButtonMap, string)) {
	vbox.Append(ui.NewLabel("Debugging:"), false)

	vbox.Append(ui.NewColorButton(), false)
	vbox.Append(CreateButton("Add Account", "ADD", custom), false)
	vbox.Append(CreateButton("Quit", "QUIT", custom), false)

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

	vbox.Append(CreateButton("Close GUI", "QUIT", custom), false)

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

	vbox.Append(CreateButton("DEBUG goroutines", "DEBUG", custom), false)
	vbox.Append(CreateButton("xterm", "XTERM", runTestExecClick), false)
}

func runTestExecClick(b *ButtonMap, msg string) {
	log.Println("runTestExecClick START")
	go runCommand("xterm -report-fonts")
	log.Println("runTestExecClick END")
}

func runCommand(s string) {
	log.Println("runXterm START")
	log.Println("runXterm START")
	log.Println("runXterm START")
	cmd := strings.TrimSpace(s) // this is like 'chomp' in perl
        cmdArgs := strings.Fields(cmd)
	process := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	log.Println("runXterm process.Start()")
	process.Start()
	log.Println("runXterm process.Wait()")
	process.Wait()
	log.Println("runXterm END")
	log.Println("runXterm END")
	log.Println("runXterm END")
}
