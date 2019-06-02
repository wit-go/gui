package gui

import "log"
import "time"
import "fmt"
import "reflect"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

// import pb "git.wit.com/wit/witProtobuf"

// THIS IS NOT CLEAN (but probably doesn't need to be. it's debugging)

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

/*
func addDebuggingButtons(box *GuiBox) {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	box.UiBox.Append(vbox, false)

	vbox.Append(ui.NewLabel("Debugging:"), false)

	vbox.Append(ui.NewColorButton(), false)
	a := CreateButton(box, nil, nil, "Add Account", "ADD", nil)
	vbox.Append(a.B, false)
	a = CreateButton(box, nil, nil, "Quit", "QUIT", nil)
	vbox.Append(a.B, false)

	// ATTEMPT TO ADD THE TABLE HERE
	add2button := ui.NewButton("Add a Test Table")
	add2button.OnClicked(func(*ui.Button) {
		log.Println("send over socket")
		addTableTab()
	})
	vbox.Append(add2button, false)
	// ATTEMPT TO ADD THE TABLE HERE END

	a = CreateButton(box, nil, nil, "Hide & Show Box1&2", "HIDE", runTestHide)
	vbox.Append(a.B, false)

	a = CreateButton(box, nil, nil, "Close GUI", "QUIT", nil)
	vbox.Append(a.B, false)
	a = CreateButton(box, nil, nil, "DEBUG goroutines", "DEBUG", nil)
	vbox.Append(a.B, false)
	a = CreateButton(box, nil, nil, "xterm", "XTERM", runTestExecClick)
	vbox.Append(a.B, false)
	a = CreateButton(box, nil, nil, "Load test.json config file", "CONFIG", nil)
	vbox.Append(a.B, false)
}
*/

func runTestHide(b *GuiButton) {
	/*
	log.Println("runTestHide START")
	Data.Window1.Box1.Hide()
	Data.Window1.Box2.Hide()
	// time.Sleep(2000 * time.Millisecond)
	Data.State = "HIDE"
	log.Println("runTestHide END")
	*/
}

/*
func runPingClick(b *GuiButton) {
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

func runTestExecClick(b *GuiButton) {
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
*/

//
// this watches the GUI primarily to process protobuf's
// this is pointless or wrong but I use it for debugging
//
func WatchGUI() {
	count := 0

	for {
		if (count > 20) {
			log.Println("Sleep() in watchGUI() Data.State =", Data.State)
			for i, window := range Data.Windows {
				log.Println("watchGUI() Data.Windows i =", i, "Action =", window.Action)
				for name, abox := range window.BoxMap {
					log.Println("\twatchGUI() BOX name =", name)
					if (name == "SplashArea3") {
						log.Println("\t\twatchGUI() BOX abox =", reflect.TypeOf(abox))
						win := abox.Window
						log.Println("\t\twatchGUI() BOX win =", reflect.TypeOf(win))
						area := win.Area
						log.Println("\t\twatchGUI() BOX area =", reflect.TypeOf(area), area.UiArea)
						// spew.Dump(area.UiArea)
						// area.UiArea.Show()
						// time.Sleep(2000 * time.Millisecond)
						// os.Exit(0)
					}
				}
			}
			count = 0
		}
		count += 1
		time.Sleep(200 * time.Millisecond)
	}
}
