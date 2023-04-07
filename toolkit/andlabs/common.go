package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) doUserEvent() {
	if (callback == nil) {
		log(debugError, "doUserEvent() callback == nil", t.wId)
		return
	}
	var a toolkit.Action
	a.WidgetId = t.wId
	a.Name = t.Name
	a.S = t.s
	a.I = t.i
	a.B = t.b
	a.ActionType = toolkit.User
	log(logInfo, "doUserEvent() START: send a user event to the callback channel")
	callback <- a
	log(logInfo, "doUserEvent() END:   sent a user event to the callback channel")
	return
}
