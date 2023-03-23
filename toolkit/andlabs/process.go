// myplugin/myplugin.go
package main

/*
from chatgpt:

// put this in widget.go
import (
    "fmt"
    // "toolkit"
)

type Plugin interface {
    Process(input chan string, output chan string)
}

// put this in wit/gui/toolkit/*
type myPlugin struct{}

var Plugin myPlugin

func (p *myPlugin) Process(input chan string, output chan string) {
    go func() {
        for msg := range input {
            // Your processing logic goes here
            result := fmt.Sprintf("Processed: %s", msg)
            output <- result
        }
    }()
}

// main.go put this in wit/gui
package main

import (
    "fmt"
    "plugin"
    "pluginapi"
)

func main() {
    plug, err := plugin.Open("myplugin.so")
    if err != nil {
        panic(err)
    }

    symPlugin, err := plug.Lookup("Plugin")
    if err != nil {
        panic(err)
    }

    p, ok := symPlugin.(pluginapi.Plugin)
    if !ok {
        panic("Invalid plugin type")
    }

    input := make(chan string)
    output := make(chan string)

    p.Process(input, output)

    input <- "Hello, World!"
    close(input)

    for result := range output {
        fmt.Println(result)
    }
}

*/

// func main() {}
