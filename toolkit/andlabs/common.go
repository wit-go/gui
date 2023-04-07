package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) commonChange(tw *toolkit.Widget, wId int) {
	log(debugChange, "commonChange() START widget   =", t.Name, t.WidgetType)
//	if (sendToChan(wId)) {
//		log(debugChange, "commonChange() END attempted channel worked", t.Name, t.WidgetType)
//		return
//	}
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

func (t *andlabsT) doUserEvent() {
	if (callback == nil) {
		log(debugError, "douserEvent() callback == nil", t.wId)
		return
	}
	var a toolkit.Action
	a.WidgetId = t.wId
	a.Name = t.Name
	a.S = t.s
	a.I = t.i
	a.B = t.b
	a.ActionType = toolkit.User
	log(logNow, "START: send a user event to the callback channel")
	callback <- a
	log(logNow, "END:   sent a user event to the callback channel")
	return
}
