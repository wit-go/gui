package main

import (
//	"github.com/awesome-gocui/gocui"
	"go.wit.com/gui/toolkit"
)

func (n *node) setCheckbox(b bool) {
	w := n.tk
	if (n.WidgetType != toolkit.Checkbox) {
		return
	}
	if (b) {
		n.B = b
		n.Text = "X " + n.Name
	} else {
		n.B = b
		n.Text = "  " + n.Name
	}
	t := len(n.Text) + 1
	w.gocuiSize.w1 = w.gocuiSize.w0 + t

//	w.realWidth = w.gocuiSize.Width() + me.PadW
//	w.realHeight = w.gocuiSize.Height() + me.PadH

//	if w.frame {
//		w.realWidth += me.FramePadW
//		w.realHeight += me.FramePadH
//	}

	n.deleteView()
	n.showView()
}
