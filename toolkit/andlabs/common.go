package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) commonChange(tw *toolkit.Widget) {
	log(debugChange, "commonChange() START widget   =", t.tw.Name, t.tw.Type)
	if (tw == nil) {
		log(true, "commonChange() What the fuck. there is no widget t.tw == nil")
		return
	}
	if (tw.Custom == nil) {
		log(debugChange, "commonChange() END    Widget.Custom() = nil", t.tw.Name, t.tw.Type)
		return
	}
	tw.Custom()
	log(debugChange, "commonChange() END   Widget.Custom()", t.tw.Name, t.tw.Type)
}

func dump(p *toolkit.Widget, c *toolkit.Widget, b bool) {
	log(b, "Parent:")
	pt := mapToolkits[p]
	if (pt == nil) {
		log(b, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	pt.Dump(b)

	log(b, "Child:")
	ct := mapToolkits[c]
	if (ct == nil) {
		log(b, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	ct.Dump(b)
}
