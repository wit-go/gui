package toolkit

type WidgetType int
type ActionType int

// passes information between the toolkit library (plugin)
//
// All Toolkit interactions should be done via a channel or Queue()
// TODO: FIGURE OUT HOW TO IMPLEMENT THIS
//
// This is the only thing that is passed between the toolkit plugin
//
// what names should be used? This is not part of [[Graphical Widget]]
// Event() seems like a good name. 
//	Event is used too much: web dev, cloud, etc
//	I'm using "Action". Maybe it should really be
//	"Interaction" as per wikipedia [[User interface]]
// Could a protobuf be used here? (Can functions be passed?)
type Widget struct {
	Name   string
	// Action string		// "New", "Delete", "Set", aka something to do
	Type  WidgetType

	// This function is how you interact with the toolkit
	// latest attempt. seems to work so far (2023/02/28)
	// Hopefully this will be the barrier between the goroutines
	// TODO: move this interaction to channels
	Custom    func()
	Callback  func()

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

type Action struct {
	Type ActionType

	// this should be the widget
	// if the action is New, Hide, Enable, etc
	Widget *Widget

	// this is the widget 
	// where the other one should be put on New, Move, etc
	Where *Widget

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


// https://ieftimov.com/post/golang-datastructures-trees/
// TODO: protobuf ?
const (
	Unknown WidgetType = iota
	Window
	Tab	// internally, this should be a window (?)
	Frame	// should windows and tab's be frames (?)
	Grid	// a grid of frames ?
	Group	// internally, this should be a grid (?)
	Box	// internally, this should be a grid (?)
	Button
	Checkbox
	Dropdown
	Combobox // dropdown with edit=true (?)
	Label
	Textbox	// is this a Label with edit=true?
	Slider
	Spinner
	Image
	Area
	Form
	Font
	Color
	Dialog
)

const (
	Add ActionType = iota
	Delete
	Get
	Set
	SetFlag
	GetText
	SetText
	AddText
	Show
	Hide
	Enable
	Disable
	Margin
	Unmargin
	Pad
	Unpad
	Append
	Move
	Dump
	Flag
)

func (s WidgetType) String() string {
	switch s {
	case Window:
		return "Window"
	case Tab:
		return "Tab"
	case Frame:
		return "Frame"
	case Grid:
		return "Grid"
	case Group:
		return "Group"
	case Box:
		return "Box"
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
	case Image:
		return "Image"
	case Area:
		return "Area"
	case Form:
		return "Form"
	case Font:
		return "Font"
	case Color:
		return "Color"
	case Dialog:
		return "Dialog"
	case Unknown:
		return "Unknown"
	}
	return "Widget.Type.String() Error"
}

func (s ActionType) String() string {
	switch s {
	case Add:
		return "Add"
	case Delete:
		return "Delete"
	case Get:
		return "Get"
	case Set:
		return "Set"
	case SetFlag:
		return "SetFlag"
	case GetText:
		return "GetText"
	case SetText:
		return "SetText"
	case AddText:
		return "AddText"
	case Show:
		return "Show"
	case Hide:
		return "Hide"
	case Enable:
		return "Enable"
	case Disable:
		return "Disable"
	case Margin:
		return "Margin"
	case Unmargin:
		return "Unmargin"
	case Pad:
		return "Pad"
	case Unpad:
		return "Unpad"
	case Append:
		return "Append"
	case Move:
		return "Move"
	case Flag:
		return "Flag"
	case Dump:
		return "Dump"
	}
	return "Action.Type.String() Error"
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
