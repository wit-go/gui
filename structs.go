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
	State		string  // used like a state machine

	// a fallback default function to handle mouse events 
	// if nothing else is defined to handle them
	MouseClick	func(*GuiButton)

	// passes in all the User accounts from the cloud-control-panel config file
	Config		*pb.Config

	// general information on the App
	// move all this to Config (?)
	Version		string
	GitCommit	string
	GoVersion	string
	Buildtime	string
	HomeDir		string
	Debug		bool
	DebugTable	bool

	// official hostname and IPv6 address for this box
	// also move all this to Config (?)
	Hostname	string
	IPv6		string

	// A map of all the entry boxes
	AllEntries	[]*GuiEntry
	Windows		[]*GuiWindow

	// A map of all buttons everywhere on all
	// windows, all tabs, across all goroutines
	// This is "GLOBAL"
	AllButtons	[]*GuiButton
	ButtonMap	map[*GuiButton][]func (*GuiButton)

	EntryNick	*ui.Entry
	EntryUser	*ui.Entry
	EntryPass	*ui.Entry
}

// stores information on 'the' window

// More than one Window is not supported in a cross platform
// sense & may never be. On Windows and MacOS, you have to have
// 'tabs'. Even under Linux, more than one Window is currently
// unstable
//
// This code will keep track of if the windows is 'tabbed' or
// not. You can draw one thing in the window, then destroy
// that, then redraw the window with something else
//
// This struct keeps track of what is in the window so you
// can destroy and replace it with something else
//
type GuiWindow struct {
	Action		string
//	Area		*GuiArea	// should be moved to GuiBox
	BoxMap		map[string]*GuiBox
	Width		int
	Height		int

	UiWindow	*ui.Window
	UiTab		*ui.Tab		// if this != nil, the window is 'tabbed'
}


// Note: every mouse click is handled
// as a 'Button' regardless of where
// the user clicks it. You could probably
// call this 'GuiMouseClick'
type GuiButton struct {
	Action		string		// what type of button
	Box		*GuiBox		// what box the button click was in
	Area		*GuiArea	// indicates the button click was in an Area
	GW		*GuiWindow	// what window the button click was in (redundant?)

	Account		*pb.Account	// associated with what account?
	VM		*pb.Event_VM	// associated with which VM?

	// a callback function for the main application
	custom		func (*GuiButton)

	// andlabs/ui abstraction mapping
	B		*ui.Button
	FB		*ui.FontButton
	A		*ui.Area	// should be deprecated
	W		*ui.Window	// should be deprecated
	T		*ui.Tab		// should be deprecated
}

type GuiBox struct {
	W		*GuiWindow
	EntryMap	map[string]*GuiEntry
	Area		*GuiArea

	UiBox		*ui.Box
}

type GuiEntry struct {
	Action		string	// what type of button
	Edit		bool
	Last		string	// the last value
	Normalize	func (string) string // function to 'normalize' the data

	B		*GuiButton
	Account		*pb.Account
	VM		*pb.Event_VM

	E		*ui.Entry
	W		*ui.Window  // should be moved to *GuiWindow or GuiBox
	T		*ui.Tab     // should be moved to *GuiWindow or GuiBox

}

//
// AREA STRUCTURES START
// AREA STRUCTURES START
// AREA STRUCTURES START
//
type GuiArea struct{
	Window		*GuiWindow // what window this area is in (should be GuiBox?)
	Button		*GuiButton // what button handles mouse events

	UiAttrstr	*ui.AttributedString
	UiArea		*ui.Area
}

type FontString struct {
	S		string
	Size		int
	F		font.Face
	W		font.Weight
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
	RowCount		int			// This is the number of 'rows' which really means data elements not what the human sees
	RowWidth		int			// This is how wide each row is
	Rows			[]RowData		// This is all the table data by row
	generatedColumnTypes	[]ui.TableValue		// generate this dynamically

	Cells			[20]CellData
	Human			[20]HumanMap

	Account			*pb.Account	// what account this table is for
	Box			*GuiBox

	lastRow			int
	lastColumn		int
//	parentTab		*ui.Tab
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

type TableColumnData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
}

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
// TABLE DATA STRUCTURES END
//
