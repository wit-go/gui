package gui

import (
	"regexp"
	// "errors"
	// "git.wit.org/wit/gui/toolkit"
)

// functions for handling text related GUI elements

func (n *Node) Add(str string) {
	log(debugGui, "gui.Add() value =", str)
	n.widget.Action = "Add"
	n.widget.S = str
	send(n.parent, n)
}

func (n *Node) SetText(str string) bool {
	log(debugChange, "gui.SetText() value =", str)
	n.widget.Action = "SetText"
	n.widget.S = str
	send(n.parent, n)
	return true
}

func (n *Node) Set(a any) bool {
	log(debugChange, "gui.Set() value =", a)
	n.widget.Action = "Set"
	switch v := a.(type) {
	case bool:
		n.widget.B = a.(bool)
	case string:
		n.widget.S = a.(string)
	case int:
		n.widget.I = a.(int)
	default:
		log(debugError, "gui.Set() unknown type =", v, "a =", a)
	}
	send(n.parent, n)
	return true
}

func (n *Node) AppendText(str string) bool {
	n.widget.Action = "Set"
	tmp := n.widget.S + str
	log(debugChange, "gui.AppendText() value =", tmp)
	n.widget.S = tmp
	send(n.parent, n)
	return true
}

func (n *Node) GetText() string {
	return n.widget.S
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

func Delete(c *Node) {
	c.widget.Action = "Delete"
	send(c.parent, c)
}

func commonCallback(n *Node) {
	// TODO: make all of this common code to all the widgets
	// This might be common everywhere finally (2023/03/01)
	if (n.Custom == nil) {
		log(debugChange, "Not Running n.Custom(n) == nil")
	} else {
		log(debugChange, "Running n.Custom(n)")
		n.Custom()
	}
}
