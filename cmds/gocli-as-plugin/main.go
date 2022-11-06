// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"plugin"
//	"github.com/awesome-gocui/gocui"
)

type Greeter interface {
	Greet()
}

var plugGocli *plugin.Plugin
var plugHello *plugin.Plugin

func main() {
	log.Println("attempt plugin")

	go loadPlugin(plugHello, "../../toolkit/hello.so")
	loadPlugin(plugGocli, "../../toolkit/gocli.so")
}

func loadPlugin(plug *plugin.Plugin, name string) {
	// load module
	// 1. open the so file to load the symbols
	plug, err = plugin.Open(name)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("symGreater", symGreeter)

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	// var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		log.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	greeter.Greet()
}
