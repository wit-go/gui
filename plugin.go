package gui

// This is based off of the excellent example and documentation here:
// https://github.com/vladimirvivien/go-plugin-example
// There truly are great people in this world.
// It's a pleasure to be here with all of you

import (
	"log"
	"os"

	"plugin"
	"github.com/davecgh/go-spew/spew"
)

// TODO: could a protobuf work here?
type Greeter interface {
	Greet()
	JcarrButton()
	AddButton(string)
}

var PlugGocli *plugin.Plugin
var PlugGocliOk bool
var PlugHello *plugin.Plugin

// var gBut plugin.Symbol
var jcarrBut plugin.Symbol
var symGreeter plugin.Symbol
var greeter Greeter
var ok bool

var typeToolkit plugin.Symbol
var typeToolkitCast Greeter

func LoadPlugin(name string) *plugin.Plugin {
	scs := spew.ConfigState{MaxDepth: 1}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(name)
	log.Println("plug =")
	log.Println(scs.Sdump(plug))
	if err != nil {
		log.Println(err)
		return nil
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	typeToolkit, err = plug.Lookup("Toolkit")
	log.Println("plugin.Toolkit", typeToolkit)
	log.Println(scs.Sdump(typeToolkit))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	symGreeter, err = plug.Lookup("Greeter")
	log.Println("symGreater", symGreeter)
	log.Println(scs.Sdump(symGreeter))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	// var greeter Greeter
	greeter, ok = symGreeter.(Greeter)
	log.Println("greeter", symGreeter)
	log.Println(scs.Sdump(greeter))
	if !ok {
		log.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	/*
	typeToolkitCast, ok = typeToolkit.(Greeter)
	if !ok {
		log.Println("unexpected cast of Toolkit to Greeter")
		os.Exit(1)
	}
	*/
	return plug
}

func RunGreet() {
	log.Println("gui.RunGreet() START")
	if (greeter == nil) {
		log.Println("wit/gui gocui plugin didn't load")
		return
	}
	PlugGocliOk = true
	greeter.Greet()
}

func LookupJcarrButton() {
	log.Println("lookupJcarrButton() START")

	if (greeter == nil) {
		log.Println("wit/gui gocui plugin didn't load")
		return
	}
	greeter.JcarrButton()
}

func GocuiAddButton(name string) {
	log.Println("GocuiAddButton() START", name)

	if (greeter == nil) {
		log.Println("wit/gui gocui plugin didn't load")
		return
	}
	greeter.AddButton(name)
}
