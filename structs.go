package gui

import "image/color"
import "golang.org/x/image/font"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

//
// All GUI Data Structures and functions that are external
// If you need cross platform support, these might only
// be the safe way to interact with the GUI
//
var Data	GuiData

type GuiData struct {
	State		string
	// Width		int
	// Height		int

	// a fallback default function to handle mouse events 
	// if nothing else is defined to handle them
	MouseClick	func(*GuiButton)

	// account entry textboxes
	Config		*pb.Config

	// general information on the App
	Version		string
	GitCommit	string
	GoVersion	string
	Buildtime	string
	HomeDir		string
	Debug		bool
	DebugTable	bool

	// official hostname and IPv6 address for this box
	Hostname	string
	IPv6		string

	// A map of all buttons everywhere on all
	// windows, all tabs, across all goroutines
	// This is "GLOBAL"
	AllButtons	[]GuiButton

	// A map of all the entry boxes
	AllEntries	[]*GuiEntry

	Windows		[]*GuiWindow

	EntryNick	*ui.Entry
	EntryUser	*ui.Entry
	EntryPass	*ui.Entry
}

type TableColumnData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
}

type GuiEntry struct {
	E		*ui.Entry
	Edit		bool
	Last		string	// the last value
	Normalize	func (string) string // function to 'normalize' the data

	Account		*pb.Account
	VM		*pb.Event_VM

	B		*GuiButton
	FB		*ui.FontButton
	A		*ui.Area
	W		*ui.Window
	T		*ui.Tab

	Action		string	// what type of button
}

type GuiWindow struct {
	W		*ui.Window
	T		*ui.Tab
	Box1		*ui.Box
	Box2		*ui.Box

	EntryMap	map[string][]*GuiEntry

	C		*pb.Config

	AH		*AreaHandler
	Action		string
}

type FontString struct {
	S		string
	Size		int
	F		font.Face
	W		font.Weight
}

type GuiButton struct {
	// andlabs/ui stuff
	B		*ui.Button
	FB		*ui.FontButton
	A		*ui.Area
	W		*ui.Window
	T		*ui.Tab

	AH		*AreaHandler

	// git.wit.com/wit/gui stuff
	WM		*GuiWindow
	Account		*pb.Account
	VM		*pb.Event_VM
	Action		string	// what type of button

	// a callback function for the main application
	custom		func (*GuiButton)
}


// AREA STRUCTURES START
type AreaHandler struct{
	Button		*GuiButton
	Attrstr		*ui.AttributedString
	Area		*ui.Area
	WM		*GuiWindow
}
// AREA STRUCTURES END

//
// TABLE DATA STRUCTURES START
//

type CellData struct {
	Index		int
	HumanID		int
	Name		string			// what type of cell is this?
}

// hmm. will this stand the test of time?
type RowData struct {
	Name		string			// what kind of row is this?
	Status		string			// status of the row?
/*
	// TODO: These may or may not be implementable
	click		func()			// what function to call if the user clicks on it
	doubleclick	func()			// what function to call if the user double clicks on it
*/
	HumanData	[20]HumanCellData

	// The VM from the protobuf
	VM		*pb.Event_VM
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
// TODO: add back image support and the progress bar
//
type HumanCellData struct {
	Name		string			// what kind of row is this?
	Text		string
	TextID		int
	Color		color.RGBA
	ColorID		int
	VM		*pb.Event_VM
	Button		*GuiButton
}

type HumanMap struct {
	Name		string			// what kind of row is this?
	TextID		int
	ColorID		int
}

//
// This is the structure that andlabs/ui uses to pass information
// to the GUI. This is the "authoritative" data.
//
type TableData struct {
	RowCount		int			// This is the number of 'rows' which really means data elements not what the human sees
	RowWidth		int			// This is how wide each row is
	Rows			[]RowData		// This is all the table data by row
	generatedColumnTypes	[]ui.TableValue		// generate this dynamically

	Cells			[20]CellData
	Human			[20]HumanMap

	Account			*pb.Account	// what account this table is for

	lastRow			int
	lastColumn		int
	parentTab		*ui.Tab
}

//
// TABLE DATA STRUCTURES END
//
