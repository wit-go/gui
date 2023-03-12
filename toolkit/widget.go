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
	Action string		// "New", "Delete", "Set", aka something to do
	Type  WidgetType

	// This function is how you interact with the toolkit
	// latest attempt. seems to work so far (2023/02/28)
	// Hopefully this will be the barrier between the goroutines
	// TODO: move this interaction to channels
	Custom    func()

	// re-adding an id to test channels
	id     int

	// This is how the values are passed back and forth
	// values from things like checkboxes & dropdown's
	// The string is also used to set the button name
	B	bool
	I	int
	// maybe safe if there is correctly working Custom() between goroutines?
	// (still probably not, almost certainly not. not possible. layer violation?)
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
	Box
	Image
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
	case Box:
		return "Box"
	case Image:
		return "Image"
	case Flag:
		return "Flag"
	case Unknown:
		return "Unknown"
	}
	return "GuiToolkitTUndefinedType"
}

// this is hopefully just used in a very few places for
// debuging the interaction between go apps and the underlying
// toolkits. Hopefully this is less prone to problems and can
// detect memory leaks, threading problems, memory allocation & mapping errors, etc
func (w *Widget) GetId() int {
	return w.id
}

func (w *Widget) SetId(i int) {
	w.id = i
}
