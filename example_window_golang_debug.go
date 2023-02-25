package gui

import (
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
)

func GolangDebugWindow() {
	var w, t *Node

	Config.Title = "Go Language Debug Window"
	Config.Width = 400
	Config.Height = 400
	Config.Exit = StandardClose
	w = NewWindow()

	t = w.NewTab("Debug Tab")
	log("debugWindow() START")


	///////////////////////////////  Column DEBUG GOLANG   //////////////////////
	g := t.NewGroup("GO Language")

	g.NewButton("runtime.Stack()", func () {
		log("\tSTART")
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		log("\t %s", buf)
		log("\tEND")
	})
	g.NewButton("dumpModuleInfo()", func () {
		log("\tSTART")
		dumpModuleInfo()
		log("\tEND")
	})
	g.NewButton("debug.PrintStack()", func () {
		log("\tSTART")
		debug.PrintStack()
		log("\tEND")
	})
	g.NewButton("pprof.Lookup(goroutine)", func () {
		log("\tSTART")
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		log("\tEND")
	})
	g.NewButton("pprof.Lookup(heap)", func () {
		log("\tSTART")
		pprof.Lookup("heap").WriteTo(os.Stdout, 1)
		log("\tEND")
	})
	g.NewButton("pprof.Lookup(block)", func () {
		log("\tSTART")
		pprof.Lookup("block").WriteTo(os.Stdout, 1)
		log("\tEND")
	})
	g.NewButton("pprof.Lookup threadcreate", func () {
		log("\tSTART")
		pprof.Lookup("threadcreate").WriteTo(os.Stdout, 1)
		log("\tEND")
	})
	g.NewButton("runtime.ReadMemStats", func () {
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		log("alloc: %v bytes\n", s.Alloc)
		log("total-alloc: %v bytes\n", s.TotalAlloc)
		log("sys: %v bytes\n", s.Sys)
		log("lookups: %v\n", s.Lookups)
		log("mallocs: %v\n", s.Mallocs)
		log("frees: %v\n", s.Frees)
		log("heap-alloc: %v bytes\n", s.HeapAlloc)
		log("heap-sys: %v bytes\n", s.HeapSys)
		log("heap-idle: %v bytes\n", s.HeapIdle)
		log("heap-in-use: %v bytes\n", s.HeapInuse)
		log("heap-released: %v bytes\n", s.HeapReleased)
		log("heap-objects: %v\n", s.HeapObjects)
		log("stack-in-use: %v bytes\n", s.StackInuse)
		log("stack-sys: %v bytes\n", s.StackSys)
		log("next-gc: when heap-alloc >= %v bytes\n", s.NextGC)
		log("last-gc: %v ns\n", s.LastGC)
		log("gc-pause: %v ns\n", s.PauseTotalNs)
		log("num-gc: %v\n", s.NumGC)
		log("enable-gc: %v\n", s.EnableGC)
		log("debug-gc: %v\n", s.DebugGC)
	})
}

func dumpModuleInfo() {
	tmp, _ := debug.ReadBuildInfo()
	if tmp == nil {
		log("This wasn't compiled with go module support")
		return
	}
	log("mod.Path         = ", tmp.Path)
	log("mod.Main.Path    = ", tmp.Main.Path)
	log("mod.Main.Version = ", tmp.Main.Version)
	log("mod.Main.Sum     = ", tmp.Main.Sum)
	for _, value := range tmp.Deps {
		log("\tmod.Path    = ", value.Path)
		log("\tmod.Version = ", value.Version)
	}
}

