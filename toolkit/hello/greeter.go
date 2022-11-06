package main

import (
//	"errors"
	"fmt"
//	"log"
//	"strings"

//	"github.com/awesome-gocui/gocui"
)

type greeting string


// func main() {
func (g greeting) Greet() {
	fmt.Println("Hello Universe")
	Init()
	// ToolkitMain()
}

// this is exported
var Greeter greeting
