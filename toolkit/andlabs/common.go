package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) commonChange(tw *toolkit.Widget, wId int) {
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

	if (andlabs[wId] == nil) {
		log(debugError, "commonChange() ERROR: wId map == nil", wId)
		return
	}
	sendToChan(wId)

	log(debugChange, "commonChange() END   Widget.Custom()", t.tw.Name, t.tw.Type)
}

func sendToChan(i int) {
	if (callback == nil) {
		log(debugError, "commonChange() SHOULD SEND int back here, but callback == nil", i)
		return
	}
	log(debugError, "commonChange() Running callback() i =", i)
	callback(i)
}
