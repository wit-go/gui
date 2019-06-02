package gui

import "log"
import "time"
import "regexp"
// import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

const Xaxis = 0
const Yaxis = 1

func GuiInit() {
	ui.OnShouldQuit(func() bool {
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})
}

// func InitGuiWindow(c *pb.Config, action string, maketab func(*GuiWindow) *GuiBox, uiW *ui.Window, uiT *ui.Tab) *GuiWindow {
func InitGuiWindow(c *pb.Config, action string, gw *GuiWindow) *GuiWindow {
	log.Println("InitGuiWindow() START")
	var newGuiWindow GuiWindow
	newGuiWindow.Width	= int(c.Width)
	newGuiWindow.Height	= int(c.Height)
	newGuiWindow.Action	= action
	newGuiWindow.MakeWindow	= gw.MakeWindow
	newGuiWindow.UiWindow	= gw.UiWindow
	newGuiWindow.UiTab	= gw.UiTab
	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	newGuiWindow.EntryMap["test"] = nil
	Data.Windows = append(Data.Windows, &newGuiWindow)

	log.Println("InitGuiWindow() END *GuiWindow =", &newGuiWindow)
	return &newGuiWindow
}


func StartNewWindow(c *pb.Config, bg bool, action string, callback func(*GuiWindow) *GuiBox) {
	log.Println("StartNewWindow() Create a new window")
	var junk GuiWindow
	junk.MakeWindow = callback
	window := InitGuiWindow(c, action, &junk)
	if (bg) {
		log.Println("StartNewWindow() START NEW GOROUTINE for ui.Main()")
		go ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")
			InitTabWindow(window)
		})
		time.Sleep(2000 * time.Millisecond) // this might make it more stable on windows?
	} else {
		log.Println("StartNewWindow() WAITING for ui.Main()")
		ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")
			InitTabWindow(window)
		})
	}
}

func InitTabWindow(gw *GuiWindow) {
	log.Println("InitTabWindow() THIS WINDOW IS NOT YET SHOWN")

	gw.UiWindow = ui.NewWindow("", int(gw.Width), int(gw.Height), true)
	gw.UiWindow.SetBorderless(false)

        // create a 'fake' button entry for the mouse clicks
	/*
	var newBM GuiButton
	newBM.Action	= "QUIT"
	newBM.GW	= gw
	Data.AllButtons = append(Data.AllButtons, &newBM)
	*/

	gw.UiWindow.OnClosing(func(*ui.Window) bool {
		log.Println("InitTabWindow() OnClosing() THIS WINDOW IS CLOSING gw=", gw)
                ui.Quit()
		return true
	})

	gw.UiTab = ui.NewTab()
	gw.UiWindow.SetChild(gw.UiTab)
	gw.UiWindow.SetMargined(true)

	log.Println("InitTabWindow() gw =", gw)

	gw.MakeWindow(gw)
	gw.UiWindow.Show()
}

/*
// string handling examples that might be helpful for normalizeInt()
isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

for _, username := range []string{"userone", "user2", "user-three"} {
    if !isAlpha(username) {
        fmt.Printf("%q is not valid\n", username)
    }
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

func alphaOnly(s string) bool {
   for _, char := range s {
      if !strings.Contains(alpha, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}
*/

func normalizeInt(s string) string {
	// reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Println("normalizeInt() regexp.Compile() ERROR =", err)
		return s
	}
	clean := reg.ReplaceAllString(s, "")
	log.Println("normalizeInt() s =", clean)
	return clean
}

func MessageWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBox(gw.UiWindow, msg1, msg2)
}

func ErrorWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBoxError(gw.UiWindow, msg1, msg2)
}
