package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) commonChange(tw *toolkit.Widget, wId int) {
	log(debugChange, "commonChange() START widget   =", t.Name, t.WidgetType)
	if (sendToChan(wId)) {
		log(debugChange, "commonChange() END attempted channel worked", t.Name, t.WidgetType)
		return
	}
	if (tw == nil) {
		log(true, "commonChange() What the fuck. there is no widget t.tw == nil")
		return
	}
	if (tw.Custom == nil) {
		log(debugChange, "commonChange() END    Widget.Custom() = nil", t.Name, t.WidgetType)
		return
	}
	tw.Custom()

	if (andlabs[wId] == nil) {
		log(debugError, "commonChange() ERROR: wId map == nil", wId)
		return
	}

	log(debugChange, "commonChange() END   Widget.Custom()", t.Name, t.WidgetType)
}

func sendToChan(i int) bool {
	if (callback == nil) {
		log(debugError, "commonChange() SHOULD SEND int back here, but callback == nil", i)
		return false
	}
	log(debugError, "commonChange() Running callback() i =", i)
	return callback(i)
}
