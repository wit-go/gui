package gui

// Common actions for widgets like 'Enable' or 'Hide'

import (
	"regexp"
	"git.wit.org/wit/gui/toolkit"
)

// functions for handling text related GUI elements

func (n *Node) Show() {
	var a toolkit.Action
	a.Type = toolkit.Show
	newaction(&a, n, nil)
}

func (n *Node) Hide() {
	var a toolkit.Action
	a.Type = toolkit.Hide
	newaction(&a, n, nil)
}

func (n *Node) Enable() {
	var a toolkit.Action
	a.Type = toolkit.Enable
	newaction(&a, n, nil)
}

func (n *Node) Disable() {
	var a toolkit.Action
	a.Type = toolkit.Disable
	newaction(&a, n, nil)
}

func (n *Node) Add(str string) {
	log(debugGui, "gui.Add() value =", str)

	var a toolkit.Action
	a.Type = toolkit.Add
	a.S = str
	// a.Widget = &n.widget
	// action(&a)
	newaction(&a, n, nil)
}

func (n *Node) AddText(str string) {
	log(debugChange, "AddText() value =", str)

	var a toolkit.Action
	a.Type = toolkit.AddText
	a.S = str
	// a.Widget = &n.widget
	// action(&a)
	newaction(&a, n, nil)
}

func (n *Node) SetText(str string) {
	log(debugChange, "SetText() value =", str)

	var a toolkit.Action
	a.Type = toolkit.SetText
	a.S = str
	// a.Widget = &n.widget
	// action(&a)
	newaction(&a, n, nil)
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
	a.Type = toolkit.Set

	switch v := val.(type) {
	case bool:
		a.B = val.(bool)
	case string:
		a.S = val.(string)
	case int:
		a.I = val.(int)
	default:
		log(debugError, "Set() unknown type =", v, "a =", a)
	}

	// a.Widget = &n.widget
	// action(&a)
	newaction(&a, n, nil)
}

func (n *Node) AppendText(str string) {
	var a toolkit.Action
	a.Type = toolkit.SetText
	tmp := n.widget.S + str
	log(debugChange, "AppendText() value =", tmp)
	a.S = tmp
	// a.Widget = &n.widget
	// action(&a)
	newaction(&a, n, nil)
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

func (n *Node) Margin() {
	var a toolkit.Action
	a.Type = toolkit.Margin
	newaction(&a, n, nil)
}

func (n *Node) Unmargin() {
	var a toolkit.Action
	a.Type = toolkit.Unmargin
	newaction(&a, n, nil)
}

func (n *Node) Pad() {
	var a toolkit.Action
	a.Type = toolkit.Pad
	newaction(&a, n, nil)
}

func (n *Node) Unpad() {
	var a toolkit.Action
	a.Type = toolkit.Unpad
	newaction(&a, n, nil)
}

func (n *Node) New2() *Node {
	var newWin *Node
	newWin = NewWindow()
	log(debugError, "New2() END Main(f)")
	return newWin
}

func (n *Node) Window(title string) *Node {
	log(debugError, "Window()", n)
	n.SetText(title)
	return n
}

func (n *Node) Standard() *Node {
	log(debugError, "Standard()")
	return n
}

func (n *Node) DoMargin() *Node {
	log(debugError, "DoMargin()")
	return n
}
