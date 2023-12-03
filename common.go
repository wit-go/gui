package gui

// Common actions for widgets like 'Enable' or 'Hide'

import (
	"regexp"
	"git.wit.org/wit/gui/toolkit"
)

// functions for handling text related GUI elements

func (n *Node) Show() *Node {
	a := newAction(n, toolkit.Show)
	sendAction(a)
	return n
}

func (n *Node) Hide() *Node {
	a := newAction(n, toolkit.Hide)
	sendAction(a)
	return n
}

func (n *Node) Enable() *Node {
	a := newAction(n, toolkit.Enable)
	sendAction(a)
	return n
}

func (n *Node) Disable() *Node {
	a := newAction(n, toolkit.Disable)
	sendAction(a)
	return n
}

func (n *Node) Add(str string) {
	log(debugGui, "gui.Add() value =", str)

	n.S = str

	a := newAction(n, toolkit.Add)
	sendAction(a)
}

func (n *Node) AddText(str string) {
	log(debugChange, "AddText() value =", str)

	n.Text = str
	n.S = str

	a := newAction(n, toolkit.AddText)
	sendAction(a)
}

func (n *Node) SetText(text string) *Node {
	log(debugChange, "SetText() value =", text)

	n.Text = text
	n.S = text

	a := newAction(n, toolkit.SetText)
	sendAction(a)
	return n
}

func (n *Node) SetNext(w int, h int) {
	n.NextW = w
	n.NextH = h
	log(debugNow, "SetNext() w,h =", n.NextW, n.NextH)
}

func (n *Node) Set(val any) {
	log(debugChange, "Set() value =", val)

	switch v := val.(type) {
	case bool:
		n.B = val.(bool)
	case string:
		n.Text = val.(string)
		n.S = val.(string)
	case int:
		n.I = val.(int)
	default:
		log(debugError, "Set() unknown type =", v)
	}

	a := newAction(n, toolkit.Set)
	sendAction(a)
}

func (n *Node) AppendText(str string) {
	tmp := n.S + str
	n.Text = tmp
	n.S = tmp

	a := newAction(n, toolkit.SetText)
	sendAction(a)
}

func (n *Node) GetText() string {
	return n.S
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
	// This might be common everywhere finally (2023/03/01)
	if (n.Custom == nil) {
		log(debugChange, "Not Running n.Custom(n) == nil")
	} else {
		log(debugChange, "Running n.Custom(n)")
		n.Custom()
	}
}

func (n *Node) Margin() *Node {
	a := newAction(n, toolkit.Margin)
	sendAction(a)
	return n
}

func (n *Node) Unmargin() *Node {
	a := newAction(n, toolkit.Unmargin)
	sendAction(a)
	return n
}

func (n *Node) Pad() *Node {
	a := newAction(n, toolkit.Pad)
	sendAction(a)
	return n
}

func (n *Node) Unpad() *Node {
	a := newAction(n, toolkit.Unpad)
	sendAction(a)
	return n
}

// is this better?
// yes, this is better. it allows Internationalization very easily
//  me.window = myGui.New2().Window("DNS and IPv6 Control Panel").Standard()
//  myFunnyWindow = myGui.NewWindow("Hello").Standard().SetText("Hola")

func (n *Node) Window(title string) *Node {
	log(debugError, "Window()", n)
	return n.NewWindow(title)
}

// This should not really do anything. as per the docs, the "Standard()" way
// should be the default way
func (n *Node) Standard() *Node {
	log(debugError, "Standard() not implemented yet")
	return n
}

func (n *Node) SetMargin() *Node {
	log(debugError, "DoMargin() not implemented yet")
	return n
}
