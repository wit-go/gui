package toolkit

// passes information between the toolkit library (plugin)
//
// All Toolkit interactions should be done via a channel or Queue()
// TODO: FIND THIS NOTE AND FIGURE OUT HOW TO IMPLEMENT IT
//
// This is the only thing that is passed between the toolkit plugin
//
// what names should be used? This is not part of [[Graphical Widget]]
// Event() seems like a good name.
// Could a protobuf be used here? (Can functions be passed?)
type Widget struct {
	Name   string
	Type   string	// after lots of back and forth, a simple string

	// This GUI is intended for simple things
	// We are not laying out PDF's here
	// This is used for things like a slider(0,100)
	Width  int
	Height int
	X      int
	Y      int

	// latest attempt
	Custom    func()

	// This might be useful to simplify retrieving 
	// values from things like checkboxes & dropdown's
	B	bool
	I	int
	S	string

	// other things I've tried
	// Event     func(*Widget) *Widget
//	OnChanged func(*Widget)
//	Custom    func(*Widget)
//	OnExit    func(*Widget)
}

/*
type Widget int

// https://ieftimov.com/post/golang-datastructures-trees/
const (
	Unknown Widget = iota
	Window
	Tab
	Frame
	Dropbox
	Spinner
	Label
)

func (s Widget) String() string {
	switch s {
	case Window:
		return "Window"
	case Tab:
		return "Tab"
	case Frame:
		return "Frame"
	case Label:
		return "Label"
	case Dropbox:
		return "Dropbox"
	}
	return "unknown"
}
*/
