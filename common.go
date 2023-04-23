package gui

// Common actions for widgets like 'Enable' or 'Hide'

import (
	"regexp"
	"git.wit.org/wit/gui/toolkit"
)

// functions for handling text related GUI elements

func (n *Node) Show() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Show
	newaction(&a, n, nil)
	return n
}

func (n *Node) Hide() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Hide
	newaction(&a, n, nil)
	return n
}

func (n *Node) Enable() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Enable
	newaction(&a, n, nil)
	return n
}

func (n *Node) Disable() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Disable
	newaction(&a, n, nil)
	return n
}

func (n *Node) Add(str string) {
	log(debugGui, "gui.Add() value =", str)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.S = str
	newaction(&a, n, nil)
}

func (n *Node) AddText(str string) {
	log(debugChange, "AddText() value =", str)

	n.Text = str
	var a toolkit.Action
	a.ActionType = toolkit.AddText
	a.S = str
	newaction(&a, n, nil)
}

func (n *Node) SetText(text string) *Node{
	log(debugChange, "SetText() value =", text)

	n.Text = text
	var a toolkit.Action
	a.ActionType = toolkit.SetText
	a.S = text
	newaction(&a, n, nil)
	return n
}

func (n *Node) SetNext(x int, y int) {
	n.NextX = x
	n.NextY = y
	log(debugError, "SetNext() x,y =", n.NextX, n.NextY)
	log(debugError, "SetNext() x,y =", n.NextX, n.NextY)
	log(debugError, "SetNext() x,y =", n.NextX, n.NextY)
	log(debugError, "SetNext() x,y =", n.NextX, n.NextY)
	log(debugError, "SetNext() x,y =", n.NextX, n.NextY)
	log(debugError, "SetNext() x,y =", n.NextX, n.NextY)
}

func (n *Node) Set(val any) {
	log(debugChange, "Set() value =", val)
	var a toolkit.Action
	a.ActionType = toolkit.Set

	switch v := val.(type) {
	case bool:
		n.B = val.(bool)
		a.B = val.(bool)
	case string:
		n.Text = val.(string)
		a.S = val.(string)
	case int:
		n.I = val.(int)
		a.I = val.(int)
	default:
		log(debugError, "Set() unknown type =", v, "a =", a)
	}

	newaction(&a, n, nil)
}

func (n *Node) AppendText(str string) {
	var a toolkit.Action
	a.ActionType = toolkit.SetText
	tmp := n.S + str
	log(debugChange, "AppendText() value =", tmp)
	a.S = tmp
	n.Text = tmp
	newaction(&a, n, nil)
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
	var a toolkit.Action
	a.ActionType = toolkit.Margin
	newaction(&a, n, nil)
	return n
}

func (n *Node) Unmargin() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Unmargin
	newaction(&a, n, nil)
	return n
}

func (n *Node) Pad() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Pad
	newaction(&a, n, nil)
	return n
}

func (n *Node) Unpad() *Node {
	var a toolkit.Action
	a.ActionType = toolkit.Unpad
	newaction(&a, n, nil)
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

func (n *Node) DoMargin() *Node {
	log(debugError, "DoMargin() not implemented yet")
	return n
}
