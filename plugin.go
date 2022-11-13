package gui

// This is based off of the excellent example and documentation here:
// https://github.com/vladimirvivien/go-plugin-example
// There truly are great people in this world.
// It's a pleasure to be here with all of you

import (
	"log"
	"os"
	"plugin"

	"git.wit.org/wit/gui/toolkit"
)

var err error
type Symbol any

type aplug struct {
	// Ok bool
	name string
	filename string
	plug *plugin.Plugin
	sym *plugin.Symbol
	LoadOk bool
	InitOk bool
	MainOk bool

	Init func()
	Main func(func ())
	Queue func(func ())
	Quit func()
	NewWindow func(*toolkit.Widget)
	NewButton func(*toolkit.Widget, *toolkit.Widget)
	NewGroup func(*toolkit.Widget, *toolkit.Widget)
	NewCheckbox func(*toolkit.Widget, *toolkit.Widget)
	NewTab func(*toolkit.Widget, *toolkit.Widget)
	NewLabel func(*toolkit.Widget, *toolkit.Widget)
	NewTextbox func(*toolkit.Widget, *toolkit.Widget)
	NewSlider func(*toolkit.Widget, *toolkit.Widget)
	NewSpinner func(*toolkit.Widget, *toolkit.Widget)

	NewDropdown func(*toolkit.Widget, *toolkit.Widget)
	AddDropdownName func(*toolkit.Widget, string)
}

var allPlugins []*aplug

// loads and initializes a toolkit (andlabs/ui, gocui, etc)
func LoadToolkit(name string) bool {
	var newPlug aplug

	log.Println("gui.LoadToolkit() START")
	newPlug.LoadOk = false

	for _, aplug := range allPlugins {
		log.Println("gui.LoadToolkit() already loaded toolkit plugin =", aplug.name)
		if (aplug.name == name) {
			log.Println("gui.LoadToolkit() SKIPPING")
			return true
		}
	}

	// locate the shared library file
	filename := name + ".so"
	loadPlugin(&newPlug, filename)
	if (newPlug.plug == nil) {
		return false
	}
	// newPlug.Ok = true
	newPlug.name = name

	// map all the functions
	newPlug.Init = loadFuncE(&newPlug, "Init")
	newPlug.Quit = loadFuncE(&newPlug, "Quit")

	// this should be laodFuncE()
	newPlug.Main  = loadFuncF(&newPlug, "Main")
	newPlug.Queue = loadFuncF(&newPlug, "Queue")

	newPlug.NewWindow = loadFunc1(&newPlug, "NewWindow")

	newPlug.NewButton = loadFunc2(&newPlug, "NewButton")
	newPlug.NewGroup = loadFunc2(&newPlug, "NewGroup")
	newPlug.NewCheckbox = loadFunc2(&newPlug, "NewCheckbox")
	newPlug.NewTab = loadFunc2(&newPlug, "NewTab")
	newPlug.NewLabel = loadFunc2(&newPlug, "NewLabel")
	newPlug.NewTextbox = loadFunc2(&newPlug, "NewTextbox")
	newPlug.NewSlider = loadFunc2(&newPlug, "NewSlider")
	newPlug.NewSpinner = loadFunc2(&newPlug, "NewSpinner")

	newPlug.NewDropdown = loadFunc2(&newPlug, "NewDropdown")
	newPlug.AddDropdownName = loadFuncS(&newPlug, "AddDropdownName")

	allPlugins = append(allPlugins, &newPlug)

	log.Println("gui.LoadToolkit() END", newPlug.name, filename)
	newPlug.LoadOk = true
	return true
}

func loadFuncE(p *aplug, funcName string) func() {
	var newfunc func()
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log.Println("DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func())
	if !ok {
		log.Println("function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

func loadFunc1(p *aplug, funcName string) func(*toolkit.Widget) {
	var newfunc func(*toolkit.Widget)
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log.Println("DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(*toolkit.Widget))
	if !ok {
		log.Println("function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

func loadFuncS(p *aplug, funcName string) func(*toolkit.Widget, string) {
	var newfunc func(*toolkit.Widget, string)
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log.Println("DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(*toolkit.Widget, string))
	if !ok {
		log.Println("function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

func loadFunc2(p *aplug, funcName string) func(*toolkit.Widget, *toolkit.Widget) {
	var newfunc func(*toolkit.Widget, *toolkit.Widget)
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log.Println("DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(*toolkit.Widget, *toolkit.Widget))
	if !ok {
		log.Println("function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

// This is probably dangerous and should never be done
// executing arbitrary functions will cause them to run inside the goroutine that
// the GUI toolkit itself is running in. TODO: move to channels here
func loadFuncF(p *aplug, funcName string) func(func ()) {
	var newfunc func(func ())
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log.Println("DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(func ()))
	if !ok {
		log.Println("function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

func loadPlugin(p *aplug, name string) {
	var filename string

	// attempt to write out the file from the internal resource
	internalName := "toolkit/" + name 
	soFile, err := res.ReadFile(internalName)
	if (err != nil) {
		log.Println(err)
	} else {
		err = os.WriteFile("/tmp/wit/" + name, soFile, 0644)
		if (err != nil) {
			log.Println(err)
		}
	}

	filename = "/tmp/wit/" + name
	p.plug = loadfile(filename)
	if (p.plug != nil) {
		p.filename = filename
		return
	}

	filename = "/usr/share/wit/gui/" + name
	p.plug = loadfile(filename)
	if (p.plug != nil) {
		p.filename = filename
		return
	}
	return
}

// load module
// 1. open the shared object file to load the symbols
func loadfile(filename string) *plugin.Plugin {
	plug, err := plugin.Open(filename)
	log.Println("plug =", plug)
	if err != nil {
		log.Println(err)
		return nil
	}
	return plug
}
