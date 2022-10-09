package gui

import (
	"image/color"
//	"log"

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
	Stretchy   bool
	Menu       bool
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
	MouseClick func(*Node)

	// A map of all the entry boxes
	AllEntries []*GuiEntry

	// Store access to everything via binary tree's
	NodeMap    map[string]*Node
	NodeArray  []*Node
	NodeSlice  []*Node
}

// text entry fields
type GuiEntry struct {
	Name      string // field for human readable name
	Edit      bool
	Last      string              // the last value
	Normalize func(string) string // function to 'normalize' the data

	N   *Node

	// andlabs/ui abstraction mapping
	UiEntry *ui.Entry
}

type GuiArea struct {
	N *Node	// what node to pass mouse events

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

	n *Node

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
	N  *Node
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
