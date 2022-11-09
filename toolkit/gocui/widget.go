package main

// passes information between the toolkit library (plugin)

// All Toolkit interactions should be done via a channel or Queue()

// This is the only thing that is passed between the toolkit plugin

// what names should be used? This is not part of [[Graphical Widget]]
// Event() seems like a good name.
// Could a protobuf be used here? (Can functions be passed?)
type Widget struct {
	i     int
	s     string

	Name   string
	Width  int
	Height int

	Event     func(*Widget) *Widget

	// Probably deprecate these
	OnChanged func(*Widget)
	Custom    func(*Widget)
	OnExit    func(*Widget)
}
