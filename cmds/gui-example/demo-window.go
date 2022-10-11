package main

import "log"
import "reflect"

import "git.wit.org/wit/gui"

import "github.com/davecgh/go-spew/spew"

func demoClick (n *gui.Node) {
		log.Println("demoClick() Dumping node:")
		n.Dump()
}

var username = "jcarr"
var hostname = "fire"

func newClick (n *gui.Node) {
	var tmp []string
	junk := "ssh -v " + username + "@" + hostname
	log.Println("junk = " , junk)
	xterm(junk)
	log.Println("tmp = " , reflect.ValueOf(tmp).Kind())
	// spew.Dump(tmp)
}

func addDemoTab(n *gui.Node, title string) {
	newNode := n.AddTab(title, nil)
	if (gui.Config.Debug) {
		newNode.Dump()
	}
	newNode.ListChildren(false)

	groupNode1 := newNode.AddGroup("group 1")
	cbNode := groupNode1.AddComboBox("username", "root", "jcarr", "hugo")
	cbNode.OnChanged(func () {
		username = cbNode.GetText()
	})
	groupNode1.AddComboBox("demoCombo3", "foo 3", "bar", "stuff")

	groupNode1.Dump()

	butNode1 := groupNode1.AddButton("button1", demoClick)
	butNode1.Dump()

	butNode2 := groupNode1.AddButton("button2", newClick)
	butNode2.Dump()

	groupNode2 := newNode.AddGroup("group 2")
	groupNode2.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")

	gNode := newNode.AddGroup("domU")
	makeSSHbutton(gNode, "hugo@www",	"www.wit.org")
	makeSSHbutton(gNode, "check.lab",	"check.lab.wit.org")
	makeSSHbutton(gNode, "gobuild.lab",	"gobuild.lab.wit.org")
	makeSSHbutton(gNode, "gobuild2.lab",	"gobuild2.lab.wit.org")

///////////////////////////////  Column DNS  ////////////////////////////////
	gNode = newNode.AddGroup("dns")
	makeSSHbutton(gNode, "bind.wit.org",	"bind.wit.org")
	makeSSHbutton(gNode, "ns1.wit.com",	"ns1.wit.com")
	makeSSHbutton(gNode, "ns2.wit.com",	"ns2.wit.com")
	makeSSHbutton(gNode, "coredns",		"coredns.lab.wit.org")

///////////////////////////////  PHYS 530  //////////////////////////////////
	gNode = newNode.AddGroup("phys 530")
	// makeXtermButton(gNode, "openwrt",	"SUBDOMAIN",  "ssh -4 -v root@openwrt")
	gNode.AddButton("openwrt", func (*gui.Node) {
		stuff := "ssh -4 -v root@openwrt"
		xterm(stuff)
	})
	makeSSHbutton  (gNode, "mirrors",	"mirrors.wit.org")
	makeSSHbutton  (gNode, "node004",	"node004.lab.wit.org")
	makeSSHbutton  (gNode, "lenovo-z70",	"lenovo-z70.lab.wit.org")

///////////////////////////////  PHYS 522  //////////////////////////////////
	gNode = newNode.AddGroup("phys 522")
	// makeXtermButton(gNode, "openwrt2",	"SUBDOMAIN", "ssh -4 -v root@openwrt2")
	gNode.AddButton("openwrt2", func (*gui.Node) {
		stuff := "ssh -4 -v root@openwrt2"
		xterm(stuff)
	})
	makeSSHbutton  (gNode, "fire.lab",	"fire.lab.wit.org")
	makeSSHbutton  (gNode, "predator",	"predator.lab.wit.org")

///////////////////////////////  FLOAT  /////////////////////////////////////
	gNode = newNode.AddGroup("float")
	makeSSHbutton(gNode, "root@asus-n501vw",	"asus-n501vw.lab.wit.org")
}

func makeSSHbutton (n *gui.Node, name string, hostname string) {
	bNode := n.AddButton(name, func (*gui.Node) {
		var tmp []string
		if (username == "") {
			username = "root"
		}
		junk := "ssh -v " + username + "@" + hostname
		log.Println("junk = " , junk)
		log.Println("username = '" + username + "'")
		xterm(junk)
		log.Println("tmp = " , reflect.ValueOf(tmp).Kind())
		spew.Dump(tmp)
	})
	bNode.Dump()
}
