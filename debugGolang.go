package gui

import (
	"fmt"
	"bytes"
	// "os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
)

func (n *Node) DebugGolangWindow() {
	var w, g, og, outputTextbox *Node

	w = n.NewWindow("GO")
	w.Custom = w.StandardClose

	g = w.NewGroup("Language Internals")

	g.NewButton("ReadModuleInfo()", func () {
		tmp, _ := debug.ReadBuildInfo()
		outputTextbox.SetText(tmp.String())
	})
	g.NewButton("runtime.NumGoroutine()", func () {
		buf := new(bytes.Buffer)
		pprof.Lookup("goroutine").WriteTo(buf, 1)
		outputTextbox.SetText(buf.String())

		outputTextbox.AppendText(fmt.Sprintln("runtime.NumGoroutine() = ", runtime.NumGoroutine()))
	})
	g.NewButton("pprof.Lookup(heap)", func () {
		buf := new(bytes.Buffer)
		pprof.Lookup("heap").WriteTo(buf, 1)
		outputTextbox.SetText(buf.String())
	})
	g.NewButton("debug.PrintStack(current)", func () {
		outputTextbox.SetText(string(debug.Stack()))
	})
	g.NewButton("pprof.Lookup(goroutine)", func () {
		buf := new(bytes.Buffer)
		pprof.Lookup("goroutine").WriteTo(buf, 1)
		outputTextbox.SetText(buf.String())
	})
	g.NewButton("pprof.Lookup(block)", func () {
		buf := new(bytes.Buffer)
		pprof.Lookup("block").WriteTo(buf, 1)
		outputTextbox.SetText(buf.String())
	})
	g.NewButton("pprof.Lookup threadcreate", func () {
		buf := new(bytes.Buffer)
		pprof.Lookup("threadcreate").WriteTo(buf, 1)
		outputTextbox.SetText(buf.String())
	})

	g.NewButton("runtime.ReadMemStats()", func () {
		outputTextbox.SetText(runtimeReadMemStats())
	})

	g.NewButton("debug.FreeOSMemory()", func () {
		var out string = "Before debug.FreeOSMemory():\n\n"
		out += runtimeReadMemStats()
		debug.FreeOSMemory()
		out += "\n\nAfter debug.FreeOSMemory():\n\n"
		out += runtimeReadMemStats()
		outputTextbox.SetText(out)
	})

	g.NewButton("debug.ReadGCStats()", func () {
		var tmp debug.GCStats
		var out string
		debug.ReadGCStats(&tmp)
		log(tmp)
		out += fmt.Sprintln("LastGC:", tmp.LastGC, "// time.Time time of last collection")
		out += fmt.Sprintln("NumGC:", tmp.NumGC, "// number of garbage collections")
		out += fmt.Sprintln("PauseTotal:", tmp.PauseTotal, "// total pause for all collections")
		out += fmt.Sprintln("Pause:", tmp.Pause, "// []time.Duration pause history, most recent first")
		out += fmt.Sprintln("PauseEnd:", tmp.Pause, "// []time.Time pause history, most recent first")
		out += fmt.Sprintln("PauseQuantiles:", tmp.PauseQuantiles, "// []time.Duration")
		outputTextbox.SetText(out)
	})

	g.NewButton("debug.SetTraceback('all')", func () {
		debug.SetTraceback("all")
	})

	g.NewButton("panic()", func () {
		panic("test")
	})

	g = w.NewGroup("TODO: finish these")

	// g.NewLabel("TODO:")

	g.NewButton("runtime.Stack(true)", func () {
		// TODO: https://stackoverflow.com/questions/61127053/how-to-list-all-the-running-goroutines-in-a-go-program
		// func Stack(buf []byte, all bool) int
	})

	g.NewButton("debug.SetMemoryLimit(int)", func () {
		// TODO:
		//debug.SetMemoryLimit(1024 * 1024 * 100)
	})

	g.NewButton("debug.SetMaxStack(int bytes)", func () {
		// default is apparently 1GB
	})

	g.NewButton("debug.SetMaxThreads(int)", func () {
		// default is apparently 10,000
	})

	g.NewButton("debug.SetTraceback('all')", func () {
		debug.SetTraceback("all")
	})

	// deprecated (probably) by String() implementation within golang
	g.NewButton("dumpModuleInfo() (deprecate)", func () {
		outputTextbox.SetText(dumpModuleInfo())
	})

	og = w.NewGroup("output")
	outputTextbox = og.NewTextbox("outputBox")
	outputTextbox.Custom = func () {
		log("custom TextBox() for golang output a =", outputTextbox.S, outputTextbox.id)
	}
}

func runtimeReadMemStats() string {
	var s runtime.MemStats
	var out string
	runtime.ReadMemStats(&s)
	out += fmt.Sprintln("alloc:", s.Alloc, "bytes")
	out += fmt.Sprintln("total-alloc:", s.TotalAlloc, "bytes")
	out += fmt.Sprintln("sys:", s.Sys, "bytes")
	out += fmt.Sprintln("lookups:", s.Lookups)
	out += fmt.Sprintln("mallocs:", s.Mallocs)
	out += fmt.Sprintln("frees:", s.Frees)
	out += fmt.Sprintln("heap-alloc:", s.HeapAlloc, "bytes")
	out += fmt.Sprintln("heap-sys:", s.HeapSys, "bytes")
	out += fmt.Sprintln("heap-idle:", s.HeapIdle,"bytes")
	out += fmt.Sprintln("heap-in-use:", s.HeapInuse, "bytes")
	out += fmt.Sprintln("heap-released:", s.HeapReleased, "bytes")
	out += fmt.Sprintln("heap-objects:", s.HeapObjects)
	out += fmt.Sprintln("stack-in-use:", s.StackInuse, "bytes")
	out += fmt.Sprintln("stack-sys", s.StackSys, "bytes")
	out += fmt.Sprintln("next-gc: when heap-alloc >=", s.NextGC, "bytes")
	out += fmt.Sprintln("last-gc:", s.LastGC, "ns")
	out += fmt.Sprintln("gc-pause:", s.PauseTotalNs, "ns")
	out += fmt.Sprintln("num-gc:", s.NumGC)
	out += fmt.Sprintln("enable-gc:", s.EnableGC)
	out += fmt.Sprintln("debug-gc:", s.DebugGC)
	return out
}

func dumpModuleInfo() string {
	var out string
	tmp, _ := debug.ReadBuildInfo()
	if tmp == nil {
		out += fmt.Sprintln("This wasn't compiled with go module support")
		return ""
	}
	out += fmt.Sprintln("mod.Path         = ", tmp.Path)
	out += fmt.Sprintln("mod.Main.Path    = ", tmp.Main.Path)
	out += fmt.Sprintln("mod.Main.Version = ", tmp.Main.Version)
	out += fmt.Sprintln("mod.Main.Sum     = ", tmp.Main.Sum)
	for _, value := range tmp.Deps {
		out += fmt.Sprintln("\tmod.Path    = ", value.Path)
		out += fmt.Sprintln("\tmod.Version = ", value.Version)
	}
	return out
}
