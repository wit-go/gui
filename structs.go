package gui

import (
	"image/color"
	"log"

	"github.com/andlabs/ui"
	"golang.org/x/image/font"

	_ "github.com/andlabs/ui/winmanifest"
)

//
// All GUI Data Structures and functions that are external
// If you need cross platform support, these might only
// be the safe way to interact with the GUI
//
var Data GuiData
var Config GuiConfig

type GuiConfig struct {
	Title      string
	Width      int
	Height     int
	Exit       func(*Node)

	Debug       bool
	DebugNode   bool
	DebugTabs   bool
	DebugTable  bool
	DebugWindow bool

	depth      int
	counter    int  // used to make unique ID's
	prefix     string
}

type GuiData struct {
	// a fallback default function to handle mouse events
	// if nothing else is defined to handle them
	MouseClick func(*GuiButton)

	// A map of all the entry boxes
	AllEntries []*GuiEntry
	WindowMap  map[string]*GuiWindow

	// Store access to everything via binary tree's
	NodeMap    map[string]*Node
	NodeArray  []*Node
	NodeSlice  []*Node

	// A map of all buttons everywhere on all
	// windows, all tabs, across all goroutines
	// This is "GLOBAL"
	//
	// This has to work this way because of how
	// andlabs/ui & andlabs/libui work
	AllButtons []*GuiButton
	buttonMap  map[*ui.Button]*GuiButton
}

type GuiTab struct {
	Name   string     // field for human readable name
	Number int        // the andlabs/ui tab index
	Window *GuiWindow // the parent Window
}

//
// stores information on the 'window'
//
// This merges the concept of andlabs/ui *Window and *Tab
//
// More than one Window is not supported in a cross platform
// sense & may never be. On Windows and MacOS, you have to have
// 'tabs'. Even under Linux, more than one Window is currently
// unstable
//
// This code will make a 'GuiWindow' regardless of if it is
// a stand alone window (which is more or less working on Linux)
// or a 'tab' inside a window (which is all that works on MacOS
// and MSWindows.
//
// This struct keeps track of what is in the window so you
// can destroy and replace it with something else
//
type GuiWindow struct {
	Name      string // field for human readable name
	Width     int
	Height    int
	Axis      int  // does it add items to the X or Y axis
	TabNumber *int // the andlabs/ui tab index

	// the callback function to make the window contents
	// MakeWindow	func(*GuiBox) *GuiBox

	// the components of the window
	BoxMap   map[string]*GuiBox
	EntryMap map[string]*GuiEntry
	Area     *GuiArea

	node	*Node

	// andlabs/ui abstraction mapping
	UiWindow *ui.Window
	UiTab    *ui.Tab // if this != nil, the window is 'tabbed'
}

func (gw *GuiWindow) Dump() {
	log.Println("gui.GuiWindow.Dump() Name       = ", gw.Name)
	log.Println("gui.GuiWindow.Dump() node       = ", gw.node)
	log.Println("gui.GuiWindow.Dump() Width      = ", gw.Width)
	log.Println("gui.GuiWindow.Dump() Height     = ", gw.Height)
}

// GuiBox is any type of ui.Hbox or ui.Vbox
// There can be lots of these for each GuiWindow
type GuiBox struct {
	Name   string     // field for human readable name
	Axis   int        // does it add items to the X or Y axis
	Window *GuiWindow // the parent Window

	node	*Node

	// andlabs/ui abstraction mapping
	UiBox *ui.Box
}

func (gb *GuiBox) Dump() {
	log.Println("gui.GuiBox.Dump() Name       = ", gb.Name)
	log.Println("gui.GuiBox.Dump() Axis       = ", gb.Axis)
	log.Println("gui.GuiBox.Dump() GuiWindow  = ", gb.Window)
	log.Println("gui.GuiBox.Dump() node       = ", gb.node)
	log.Println("gui.GuiBox.Dump() UiBox      = ", gb.UiBox)
}

func (s GuiBox) SetTitle(title string) {
	log.Println("DID IT!", title)
	if s.Window == nil {
		return
	}
	if s.Window.UiWindow == nil {
		return
	}
	s.Window.UiWindow.SetTitle(title)
	return
}

func (w *GuiWindow) SetNode(n *Node) {
	if (w.node != nil) {
		w.Dump()
		panic("gui.SetNode() Error not nil")
	}
	w.node = n
	if (w.node == nil) {
		w.Dump()
		panic("gui.SetNode() node == nil")
	}
}

func (b *GuiBox) SetNode(n *Node) {
	if (b.node != nil) {
		b.Dump()
		panic("gui.SetNode() Error not nil")
	}
	b.node = n
	if (b.node == nil) {
		b.Dump()
		panic("gui.SetNode() node == nil")
	}
}

func (s GuiBox) Append(child ui.Control, x bool) {
	if s.UiBox == nil {
		return
	}
	s.UiBox.Append(child, x)
}

// Note: every mouse click is handled
// as a 'Button' regardless of where
// the user clicks it. You could probably
// call this 'GuiMouseClick'
type GuiButton struct {
	Name string  // field for human readable name
	Box  *GuiBox // what box the button click was in

	// a callback function for the main application
	Custom func(*GuiButton)
	Values interface{}
	Color  color.RGBA

	// andlabs/ui abstraction mapping
	B  *ui.Button
	FB *ui.FontButton
	CB *ui.ColorButton
}

// text entry fields
type GuiEntry struct {
	Name      string // field for human readable name
	Edit      bool
	Last      string              // the last value
	Normalize func(string) string // function to 'normalize' the data

	B   *GuiButton
	Box *GuiBox

	// andlabs/ui abstraction mapping
	UiEntry *ui.Entry
}

//
// AREA STRUCTURES START
// AREA STRUCTURES START
// AREA STRUCTURES START
//
type GuiArea struct {
	Button *GuiButton // what button handles mouse events
	Box    *GuiBox

	UiAttrstr *ui.AttributedString
	UiArea    *ui.Area
}

type FontString struct {
	S    string
	Size int
	F    font.Face
	W    font.Weight
}

//
// AREA STRUCTURES END
// AREA STRUCTURES END
// AREA STRUCTURES END
//

//
// TABLE DATA STRUCTURES START
// TABLE DATA STRUCTURES START
// TABLE DATA STRUCTURES START
//

//
// This is the structure that andlabs/ui uses to pass information
// to the GUI. This is the "authoritative" data.
//
type TableData struct {
	RowCount             int             // This is the number of 'rows' which really means data elements not what the human sees
	RowWidth             int             // This is how wide each row is
	Rows                 []RowData       // This is all the table data by row
	generatedColumnTypes []ui.TableValue // generate this dynamically

	Cells [20]CellData
	Human [20]HumanMap

	Box *GuiBox

	lastRow    int
	lastColumn int
}

//
// This maps the andlabs/ui & libui components into a "human"
// readable cell reference list. The reason is that there
// are potentially 3 values for each cell. The Text, the Color
// and an image. These are not always needed so the number
// of fields varies between 1 and 3. Internally, the toolkit
// GUI abstraction needs to list all of them, but it's then
// hard to figure out which column goes with the columns that
// you see when you visually are looking at it like a spreadsheet
//
// This makes a map so that we can say "give me the value at
// row 4 and column 2" and find the fields that are needed
//
// TODO: re-add images and the progress bar (works in andlabs/ui)
//
type HumanCellData struct {
	Name    string // what kind of row is this?
	Text    string
	TextID  int
	Color   color.RGBA
	ColorID int
	Button  *GuiButton
}

type HumanMap struct {
	Name    string // what kind of row is this?
	TextID  int
	ColorID int
}

type TableColumnData struct {
	Index    int
	CellType string
	Heading  string
	Color    string
}

type CellData struct {
	Index   int
	HumanID int
	Name    string // what type of cell is this?
}

// hmm. will this stand the test of time?
type RowData struct {
	Name   string // what kind of row is this?
	Status string // status of the row?
	/*
		// TODO: These may or may not be implementable
		// depending on if it's possible to detect the bgcolor or what row is selected
		click		func()			// what function to call if the user clicks on it
		doubleclick	func()			// what function to call if the user double clicks on it
	*/
	HumanData [20]HumanCellData
}

//
// TABLE DATA STRUCTURES END
//
