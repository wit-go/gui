package gui

import "log"
import "time"
import "fmt"
import "strings"
import "os/exec"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

// import pb "git.wit.com/wit/witProtobuf"

// can not pass any args to this (?)
/*
func setupCloudUI() {
	Data.Window1.W = ui.NewWindow("Cloud Control Panel", Data.Width, Data.Height, false)
	Data.Window1.W.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		Data.Window1.W.Destroy()
		return true
	})

	Data.Window1.T = ui.NewTab()
	Data.Window1.W.SetChild(Data.Window1.T)
	Data.Window1.W.SetMargined(true)

	// Data.tabcount = 0
	Data.Window1.T.Append("Cloud Info", makeCloudInfoBox())
	// Data.Window1.T.SetMargined(Data.tabcount, true)

	Data.Window1.W.Show()
}
*/

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
	// AddTableTab(Data.Window1.T, 1, "test seven", 7, parts, nil)
}

func addDebuggingButtons(wm *WindowMap, vbox *ui.Box) {
	vbox.Append(ui.NewLabel("Debugging:"), false)

	vbox.Append(ui.NewColorButton(), false)
	vbox.Append(CreateButton(wm, nil, nil, "Add Account", "ADD", nil), false)
	vbox.Append(CreateButton(wm, nil, nil, "Quit", "QUIT", nil), false)

	// ATTEMPT TO ADD THE TABLE HERE
	add2button := ui.NewButton("Add a Test Table")
	add2button.OnClicked(func(*ui.Button) {
		log.Println("send over socket")
		addTableTab()
	})
	vbox.Append(add2button, false)
	// ATTEMPT TO ADD THE TABLE HERE END

	vbox.Append(CreateButton(wm, nil, nil, "Hide & Show Box1&2", "HIDE", runTestHide), false)

	vbox.Append(CreateButton(wm, nil, nil, "Close GUI", "QUIT", nil), false)
	vbox.Append(CreateButton(wm, nil, nil, "DEBUG goroutines", "DEBUG", nil), false)
	vbox.Append(CreateButton(wm, nil, nil, "xterm", "XTERM", runTestExecClick), false)
	vbox.Append(CreateButton(wm, nil, nil, "Load test.json config file", "CONFIG", nil), false)
}

func runTestHide(b *ButtonMap) {
	/*
	log.Println("runTestHide START")
	Data.Window1.Box1.Hide()
	Data.Window1.Box2.Hide()
	// time.Sleep(2000 * time.Millisecond)
	Data.State = "HIDE"
	log.Println("runTestHide END")
	*/
}

func runPingClick(b *ButtonMap) {
	log.Println("runPingClick START")
	log.Println("runTestExecClick b.VM", b.VM)
	hostname := "localhost"
	if (b.VM != nil) {
		hostname = b.VM.Hostname
	}
	spew.Dump(b)
	var tmp []string
	tmp = append(tmp, "xterm", "-geometry", "120x30", "-e", "ping " + hostname + ";sleep 3")
	go runCommand(tmp)
	log.Println("runPingClick END")
}

func runTestExecClick(b *ButtonMap) {
	log.Println("runTestExecClick START")
	if runtime.GOOS == "linux" {
		go runSimpleCommand("xterm -report-fonts")
	} else if runtime.GOOS == "windows" {
		go runSimpleCommand("mintty.exe")
	} else {
		go runSimpleCommand("xterm")
	}
	log.Println("runTestExecClick END")
}

func runSimpleCommand(s string) {
	cmd := strings.TrimSpace(s) // this is like 'chomp' in perl
        cmdArgs := strings.Fields(cmd)
	runCommand(cmdArgs)
}

func runCommand(cmdArgs []string) {
	log.Println("runCommand() START", cmdArgs)
	process := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	// process := exec.Command("xterm", "-e", "ping localhost")
	log.Println("runCommand() process.Start()")
	process.Start()
	log.Println("runCommand() process.Wait()")
	process.Wait()
	log.Println("runCommand() NEED TO CHECK THE TIME HERE TO SEE IF THIS WORKED")
	log.Println("runCommand() OTHERWISE INFORM THE USER")
	log.Println("runCommand() END")
}
