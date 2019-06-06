package gui

import "log"
// import "time"
import "regexp"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

const Xaxis = 0 // box that is horizontal
const Yaxis = 1 // box that is vertical

func init() {
	log.Println("gui.init() has been run")
}

func GuiInit() {
	Data.buttonMap	= make(map[*ui.Button]*GuiButton)
	Data.WindowMap	= make(map[string]*GuiWindow)

	ui.OnShouldQuit(func() bool {
                ui.Quit()
		return true
	})
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
