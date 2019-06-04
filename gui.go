package gui

import "log"
import "time"
import "regexp"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

const Xaxis = 0 // box that is horizontal
const Yaxis = 1 // box that is vertical

func GuiInit() {
	Data.buttonMap	= make(map[*ui.Button]*GuiButton)
	Data.WindowMap	= make(map[string]*GuiWindow)

	ui.OnShouldQuit(func() bool {
                ui.Quit()
		return true
	})
}

<<<<<<< Updated upstream
func InitGuiWindow(name string, gw *GuiWindow) *GuiWindow {
	log.Println("InitGuiWindow() START")
	var newGuiWindow GuiWindow
	newGuiWindow.Width	= Config.Width
	newGuiWindow.Height	= Config.Height
	newGuiWindow.Name	= name
	newGuiWindow.MakeWindow	= gw.MakeWindow
	newGuiWindow.UiWindow	= gw.UiWindow
	newGuiWindow.UiTab	= gw.UiTab
	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	Data.Windows		= append(Data.Windows, &newGuiWindow)

	if (Data.WindowMap == nil) {
		log.Println("gui.InitGuiWindow() making the Data.WindowMap here")
		Data.WindowMap  = make(map[string]*GuiWindow)
	}
	Data.WindowMap[name]	= &newGuiWindow

	// make a blank entry for testing
	// newGuiWindow.EntryMap["test"] = nil

	if (Data.buttonMap == nil) {
		GuiInit()
	}
	log.Println("InitGuiWindow() END *GuiWindow =", &newGuiWindow)
	return &newGuiWindow
}


func StartNewWindow(bg bool, name string, callback func(*GuiWindow) *GuiBox) {
	log.Println("StartNewWindow() Create a new window")
	var junk GuiWindow
	junk.MakeWindow = callback
	window := InitGuiWindow(name, &junk)
	if (bg) {
		log.Println("StartNewWindow() START NEW GOROUTINE for ui.Main()")
		go ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")
			go InitTabWindow(window)
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
	log.Println("InitTabWindow() START. THIS WINDOW IS NOT YET SHOWN")

	gw.UiWindow = ui.NewWindow(gw.Name, int(gw.Width), int(gw.Height), true)
	gw.UiWindow.SetBorderless(false)

	gw.UiWindow.OnClosing(func(*ui.Window) bool {
		log.Println("InitTabWindow() OnClosing() THIS WINDOW IS CLOSING gw=", gw)
                ui.Quit()
		return true
	})

	gw.UiTab = ui.NewTab()
	gw.UiWindow.SetChild(gw.UiTab)
	gw.UiWindow.SetMargined(true)


	box := gw.MakeWindow(gw)
	log.Println("InitTabWindow() END box =", box)
	log.Println("InitTabWindow() END gw =", gw)
	gw.UiWindow.Show()
}

=======
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

func MessageWindow(gw *GuiWindow, msg1 string, msg2 string) {
	log.Println("gui.MessageWindow() msg1 =", msg1)
	log.Println("gui.MessageWindow() msg2 =", msg2)
	ui.MsgBox(gw.UiWindow, msg1, msg2)
}

func ErrorWindow(gw *GuiWindow, msg1 string, msg2 string) {
	log.Println("gui.ErrorWindow() msg1 =", msg1)
	log.Println("gui.ErrorWindow() msg2 =", msg2)
	ui.MsgBoxError(gw.UiWindow, msg1, msg2)
}
=======
>>>>>>> Stashed changes
