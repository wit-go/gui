package gui

import "log"
// import "errors"
import "regexp"

// functions for handling text related GUI elements

func (n *Node) NewLabel(text string) *Node {
	// make new node here
	newNode := n.New(text)
	newNode.Dump()

	t := n.toolkit.NewLabel(text)
	newNode.toolkit = t

	return newNode
}

func (n *Node) SetText(str string) bool {
	if (Config.Options.DebugChange) {
		log.Println("gui.SetText() value =", str)
	}
	if (n.toolkit == nil) {
		return false
	}

	return n.toolkit.SetText(str)
}

func (n *Node) GetText() string {
	return n.toolkit.GetText()
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
