package gui

import "log"
import "time"
import "fmt"
import "strings"
import "os/exec"
import "runtime"

import "github.com/gookit/config"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

// import pb "git.wit.com/wit/witProtobuf"

// can not pass any args to this (?)
func setupCloudUI() {
	Data.cloudWindow = ui.NewWindow("Cloud Control Panel", Data.Width, config.Int("height"), false)
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

func addDebuggingButtons(vbox *ui.Box, custom func(*ButtonMap)) {
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

	vbox.Append(CreateButton("Close GUI", "QUIT", custom), false)
	vbox.Append(CreateButton("DEBUG goroutines", "DEBUG", custom), false)
	// vbox.Append(CreateButton("ping", "PING", runPingClick), false)
	vbox.Append(CreateButton("xterm", "XTERM", runTestExecClick), false)
	vbox.Append(CreateButton("Load test.json config file", "CONFIG", custom), false)
}

func runPingClick(b *ButtonMap) {
	log.Println("runPingClick START")
	spew.Dump(b)
	var tmp []string
	tmp = append(tmp, "xterm", "-e", "ping localhost")
	runCommand(tmp)
	log.Println("runPingClick END")
}

func runTestExecClick(b *ButtonMap) {
	log.Println("runTestExecClick START")
	if runtime.GOOS == "linux" {
		go runSimpleCommand("xterm -report-fonts")
	} else if runtime.GOOS == "windows" {
		go runSimpleCommand("cmd.exe")
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
	log.Println("runXterm START")
	log.Println("runXterm START")
	log.Println("runXterm START")
	log.Println("runXterm START", cmdArgs)
	process := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	// process := exec.Command("xterm", "-e", "ping localhost")
	log.Println("runXterm process.Start()")
	process.Start()
	log.Println("runXterm process.Wait()")
	process.Wait()
	log.Println("runXterm END")
	log.Println("runXterm END")
	log.Println("runXterm END")
}
