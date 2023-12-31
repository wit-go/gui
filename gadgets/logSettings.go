package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui"
)

var myLogGui *LogSettings

type LogSettings struct {
	ready		bool
	hidden		bool
	err		error

	parent	*gui.Node // should be the root of the 'gui' package binary tree
	window	*gui.Node // our window for displaying the log package settings
	group	*gui.Node //
	grid	*gui.Node //

	// Primary Directives
	status		*OneLiner
	summary		*OneLiner
}

// This is initializes the main DO object
// You can only have one of these
func NewLogSettings(p *gui.Node) *LogSettings {
	if myLogGui != nil {return myLogGui}
	myLogGui = new(LogSettings)
	myLogGui.parent = p

	myLogGui.ready = false

	myLogGui.window = p.NewWindow("Log Settings")

	// make a group label and a grid
	myLogGui.group = myLogGui.window.NewGroup("droplets:").Pad()
	myLogGui.grid = myLogGui.group.NewGrid("grid", 2, 1).Pad()

	myLogGui.ready = true
	myLogGui.Hide()
	return myLogGui
}

// Returns true if the status is valid
func (d *LogSettings) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *LogSettings) Show() {
	if ! d.Ready() {return}
	log.Info("LogSettings.Show() window")
	if d.hidden {
		d.window.Show()
	}
	d.hidden = false
}

func (d *LogSettings) Hide() {
	if ! d.Ready() {return}
	log.Info("LogSettings.Hide() window")
	if ! d.hidden {
		d.window.Hide()
	}
	d.hidden = true
}

func (d *LogSettings) Update() bool {
	if ! d.Ready() {return false}
	return true
}
