/*
	A slider that goes between a High and Low time
*/

package gadgets

import 	(
	"log"
	"fmt"
	"time"

	"go.wit.com/gui"
)

type Duration struct {
	p	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	s	*gui.Node	// slider widget

	Label	string
	Low	time.Duration
	High	time.Duration
	Duration	time.Duration

	Custom	func()
}

func (n *Duration) Set(d time.Duration) {
	var timeRange, step, offset time.Duration

	if (d > n.High) {
		d = n.High
	}
	if (d < n.Low) {
		d = n.Low
	}

	// set the duration
	n.Duration = d

	// figure out the integer offset for the Slider GUI Widget
	timeRange = n.High - n.Low
	step = timeRange / 1000
	if (step == 0) {
		log.Println("duration.Set() division by step == 0", n.Low, n.High, timeRange, step)
		n.s.Set(0)
		return
	}
	offset = d - n.Low
	i := int(offset / step)
	log.Println("duration.Set() =", n.Low, n.High, d, "i =", i)
	n.s.I = i
	n.s.Set(i)
	n.s.Custom()
}

func NewDurationSlider(n *gui.Node, label string, low time.Duration, high time.Duration) *Duration {
	d := Duration {
		p: n,
		Label: label,
		High: high,
		Low: low,
	}

	// various timeout settings
	d.l = n.NewLabel(label)
	d.s = n.NewSlider(label, 0, 1000)
	d.s.Custom = func () {
		d.Duration = low + (high - low) * time.Duration(d.s.I) / 1000
		log.Println("d.Duration =", d.Duration)
		s := fmt.Sprintf("%s (%v)", d.Label, d.Duration)
		d.l.SetText(s)
		if (d.Custom != nil) {
			d.Custom()
		}
	}

	return &d
}
