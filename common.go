package gui

// Common actions for widgets like 'Enable' or 'Hide'

import (
	"regexp"
	"errors"
	"go.wit.com/log"
	"go.wit.com/gui/toolkits"
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
	log.Log(GUI, "gui.Add() value =", str)

	n.S = str

	a := newAction(n, toolkit.Add)
	sendAction(a)
}

func (n *Node) AddText(str string) {
	log.Log(CHANGE, "AddText() value =", str)

	n.Text = str
	n.S = str

	a := newAction(n, toolkit.AddText)
	sendAction(a)
}

func (n *Node) SetText(text string) *Node {
	log.Log(CHANGE, "SetText() value =", text)

	n.Text = text
	n.S = text

	a := newAction(n, toolkit.SetText)
	sendAction(a)
	return n
}

func (n *Node) SetNext(w int, h int) {
	n.NextW = w
	n.NextH = h
	log.Info("SetNext() w,h =", n.NextW, n.NextH)
}

func (n *Node) Set(val any) {
	log.Log(CHANGE, "Set() value =", val)

	switch v := val.(type) {
	case bool:
		n.B = val.(bool)
	case string:
		n.Text = val.(string)
		n.S = val.(string)
	case int:
		n.I = val.(int)
	default:
		log.Error(errors.New("Set() unknown type"), "v =", v)
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

// THESE TWO FUNCTIONS ARE TERRIBLY NAMED AND NEED TO BE FIXED
// 5 seconds worth of ideas:
// Value() ?
// Progname() Reference() ?

// should get the value of the node
func (n *Node) GetText() string {
	if (n.S != n.Text) {
		log.Warn("GetText() is screwed up. TODO: fix this dumb crap")
		stuff := log.ListFlags()
		log.Warn("ListFlags() =", stuff)
	}
	if (n.S != "") {
		return n.S
	}
	return n.Text
}

// should get the value of the node
// myButton = myGroup.NewButton("hit ball", nil).SetName("HIT")
// myButton.GetName() should return "HIT"
// n = Find("HIT") should return myButton
func (n *Node) GetName() string {
	return n.Name
}

/*
// string handling examples that might be helpful for normalizeInt()
isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

for _, username := range []string{"userone", "user2", "user-three"} {
    if !isAlpha(username) {
        log.Log(GUI, "%q is not valid\n", username)
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
		log.Log(GUI, "normalizeInt() regexp.Compile() ERROR =", err)
		return s
	}
	clean := reg.ReplaceAllString(s, "")
	log.Log(GUI, "normalizeInt() s =", clean)
	return clean
}

func commonCallback(n *Node) {
	// TODO: make all of this common code to all the widgets
	// This might be common everywhere finally (2023/03/01)
	if (n.Custom == nil) {
		log.Log(CHANGE, "Not Running n.Custom(n) == nil")
	} else {
		log.Log(CHANGE, "Running n.Custom(n)")
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

func (n *Node) Expand() *Node {
	a := newAction(n, toolkit.Pad)
	a.Expand = true
	sendAction(a)
	return n
}

// is this better?
// yes, this is better. it allows Internationalization very easily
//  me.window = myGui.New2().Window("DNS and IPv6 Control Panel").Standard()
//  myFunnyWindow = myGui.NewWindow("Hello").Standard().SetText("Hola")

func (n *Node) Window(title string) *Node {
	log.Warn("Window()", n)
	return n.NewWindow(title)
}

// This should not really do anything. as per the docs, the "Standard()" way
// should be the default way
/*
func (n *Node) Standard() *Node {
	log.Warn("Standard() not implemented yet")
	return n
}

func (n *Node) SetMargin() *Node {
	log.Warn("DoMargin() not implemented yet")
	return n
}
*/
