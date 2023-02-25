package gui

// import "errors"
import "regexp"

// functions for handling text related GUI elements

func (n *Node) SetText(str string) bool {
	log(debugChange, "gui.SetText() value = FIXME. NOT DOING ANYTHING", str)
	return true
}

func (n *Node) GetText() string {
	return "TODO: GetText() = {}"
}

/*
// string handling examples that might be helpful for normalizeInt()
isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

for _, username := range []string{"userone", "user2", "user-three"} {
    if !isAlpha(username) {
        log(debugGui, "%q is not valid\n", username)
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
		log(debugGui, "normalizeInt() regexp.Compile() ERROR =", err)
		return s
	}
	clean := reg.ReplaceAllString(s, "")
	log(debugGui, "normalizeInt() s =", clean)
	return clean
}

func commonCallback(n *Node) {
	// TODO: make all of this common code to all the widgets
	if (n.OnChanged == nil) {
		log(debugChange, "Not Running n.OnChanged(n) == nil")
	} else {
		log(debugChange, "Running n.OnChanged(n)")
		n.OnChanged(n)
	}

	if (n.custom == nil) {
		log(debugChange, "Not Running n.custom(n) == nil")
	} else {
		log(debugChange, "Running n.custom()")
		n.custom()
	}
}
