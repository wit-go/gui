package toolkit

type WidgetType int
type ActionType int

// passes information between the toolkit library (plugin)
//
// All Toolkit interactions should be done via a channel or Queue()
// TODO: FIGURE OUT HOW TO IMPLEMENT THIS
// https://ieftimov.com/post/golang-datastructures-trees/
// TODO: protobuf ?
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
	// Name   string
	Type  WidgetType

	// This function is how you interact with the toolkit
	// latest attempt. seems to work so far (2023/02/28)
	// Hopefully this will be the barrier between the goroutines
	// TODO: move this interaction to channels
	Custom    func()

	// re-adding an id to test channels
	Id     int

	// This is how the values are passed back and forth
	// values from things like checkboxes & dropdown's
	// The string is also used to set the button name
	B	bool
	I	int
	// maybe safe if there is correctly working Custom() between goroutines?
	// (still probably not, almost certainly not. not possible. layer violation?)
	S	string // not safe to have 'S'
}

type Action struct {
	ActionType ActionType
	WidgetType WidgetType

	WidgetId int
	ParentId int

	Text string  // what is visable to the user
	Name string  // a name useful for programming

	// this should be the widget
	// if the action is New, Hide, Enable, etc
	// Widget *Widget

	// This is how the values are passed back and forth
	// values from things like checkboxes & dropdown's
	// The string is also used to set the button name
	B	bool
	I	int
	// maybe safe if there is correctly working Custom() between goroutines?
	// (still probably not, almost certainly not. not possible. layer violation?)
	S	string // not safe to have 'S'

	A	any

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

const (
	Unknown WidgetType = iota
	Root	// the master 'root' node of the binary tree
	Flag	// used to send configuration values to plugins
	Window	// in certain gui's (ncurses), these are tabs
	Tab	// internally, this is a window
	Frame	// deprecate?
	Grid	// like drawers in a chest
	Group	// like the 'Appetizers' section on a menu
	Box	// a vertical or horizontal stack of widgets
	Button
	Checkbox // select 'on' or 'off'
	Dropdown
	Combobox // dropdown with edit=true
	Label
	Textbox	// is this a Label with edit=true
	Slider // like a progress bar
	Spinner // like setting the oven temperature
	Image // TODO
	Area // TODO
	Form // TODO
	Font // TODO
	Color // TODO
	Dialog // TODO
)

const (
	Add ActionType = iota
	User // the user did something (mouse, keyboard, etc)
	Delete
	Get
	Set
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
	Quit
)

func (s WidgetType) String() string {
	switch s {
	case Root:
		return "Root"
	case Flag:
		return "Flag"
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
	return "WidgetType.String() Error"
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
	case Dump:
		return "Dump"
	}
	return "ActionType.String() Error"
}
