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
	Name   string   // "New", "Delete", "Set", aka something to do
	Action string   // "New", "Delete", "Set", aka something to do
	// Type   string	// after lots of back and forth, a simple string
	Type  WidgetType

	// This is how the values are passed back and forth
	// values from things like checkboxes & dropdown's
	// The string is also used to set the button name
	B	bool
	I	int
	// maybe safe if there is correctly working Custom() between goroutines
	// (still probably not, almost certainly not)
	S	string // not safe to have 'S'

	// This GUI is intended for simple things
	// We are not laying out PDF's here
	// This is used for things like a slider(0,100)
	Width  int
	Height int
	X      int
	Y      int

	// Put space around elements to improve look & feel
	Margin	bool

	// Make widgets fill up the space available
	Expand	bool

	// latest attempt. seems to work so far (2023/02/28)
	// Hopefully this will be the barrier between the goroutines
	Custom    func()
}

type WidgetType int

// https://ieftimov.com/post/golang-datastructures-trees/
const (
	Unknown WidgetType = iota
	Window
	Tab
	Group
	Frame
	Button
	Checkbox
	Dropdown
	Combobox
	Label
	Textbox
	Slider
	Spinner
	Grid
	Flag
)

func (s WidgetType) String() string {
	switch s {
	case Window:
		return "Window"
	case Tab:
		return "Tab"
	case Group:
		return "Group"
	case Frame:
		return "Frame"
	case Button:
		return "Button"
	case Checkbox:
		return "Checkbox"
	case Dropdown:
		return "Dropdown"
	case Combobox:
		return "Combobox"
	case Label:
		return "Label"
	case Textbox:
		return "Textbox"
	case Slider:
		return "Slider"
	case Spinner:
		return "Spinner"
	case Grid:
		return "Grid"
	case Flag:
		return "Flag"
	case Unknown:
		return "Unknown"
	}
	return "GuiToolkitTUndefinedType"
}
