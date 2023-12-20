package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"

	"go.wit.com/gui/toolkit"
)

func simpleStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSuffix(s, "\n")
		switch s {
		case "l":
			log(true, "list widgets")
			me.rootNode.listWidgets()
		case "b":
			log(true, "show buttons")
			me.rootNode.showButtons()
		case "d":
			var a toolkit.Action
			a.ActionType = toolkit.EnableDebug
			callback <- a
		case "":
			fmt.Println("")
			fmt.Println("Enter:")
			fmt.Println("'l': list all widgets")
			fmt.Println("'b': for buttons")
			fmt.Println("'d': enable debugging")
		default:
			i, _ := strconv.Atoi(s)
			log(true, "got input:", i)
			n := me.rootNode.findWidgetId(i)
			if (n != nil) {
				n.dumpWidget("found node")
				n.doUserEvent()
			}
		}
	}
}

func (n *node) showButtons() {
	if n.WidgetType == toolkit.Button {
		n.dumpWidget("Button:")
	}

	for _, child := range n.children {
		child.showButtons()
	}
}

func (n *node) dumpWidget(pad string) {
	log(true, "node:", pad, n.WidgetId, ",", n.WidgetType, ",", n.Name)
}

var depth int = 0

func (n *node) listWidgets() {
	if (n == nil) {
		return
	}

	var pad string
	for i := 0; i < depth; i++ {
		pad = pad + "    "
	}
	n.dumpWidget(pad)

	for _, child := range n.children {
		depth += 1
		child.listWidgets()
		depth -= 1
	}
	return
}
