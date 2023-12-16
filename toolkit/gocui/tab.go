package main

// implements widgets 'Window' and 'Tab'

import (
	"strings"
	"git.wit.org/wit/gui/toolkit"
)

func (w *guiWidget) Width() int {
	if w.frame {
		return w.gocuiSize.w1 - w.gocuiSize.w0
	}
	return w.gocuiSize.w1 - w.gocuiSize.w0 - 1
}

func (w *guiWidget) Height() int {
	if w.frame {
		return w.gocuiSize.h1 - w.gocuiSize.h0
	}
	return w.gocuiSize.h1 - w.gocuiSize.h0 - 1
}

func (n *node) gocuiSetWH(sizeW, sizeH int) {
	w := len(n.Text)
	lines := strings.Split(n.Text, "\n")
	h := len(lines)

	tk := n.tk
	if tk.isFake {
		tk.gocuiSize.w0 = sizeW
		tk.gocuiSize.h0 = sizeH
		tk.gocuiSize.w1 = tk.gocuiSize.w0 + w + me.FramePadW
		tk.gocuiSize.h1 = tk.gocuiSize.h0 + h + me.FramePadH
		return
	}

	if tk.frame {
		tk.gocuiSize.w0 = sizeW
		tk.gocuiSize.h0 = sizeH
		tk.gocuiSize.w1 = tk.gocuiSize.w0 + w + me.FramePadW
		tk.gocuiSize.h1 = tk.gocuiSize.h0 + h + me.FramePadH
	} else {
		tk.gocuiSize.w0 = sizeW - 1
		tk.gocuiSize.h0 = sizeH - 1
		tk.gocuiSize.w1 = tk.gocuiSize.w0 + w + 1
		tk.gocuiSize.h1 = tk.gocuiSize.h0 + h + 1
	}
}

func redoWindows(nextW int, nextH int) {
	for _, n := range me.rootNode.children {
		if n.WidgetType != toolkit.Window {
			continue
		}
		w := n.tk
		var tabs bool
		for _, child := range n.children {
			if (child.WidgetType == toolkit.Tab) {
				tabs = true
			}
		}
		if (tabs) {
			// window is tabs. Don't show it as a standard button
			w.frame = false
			n.hasTabs = true
		} else {
			w.frame = false
			n.hasTabs = false
		}

		n.gocuiSetWH(nextW, nextH)
		n.deleteView()
		n.showView()

		sizeW := w.Width() + me.WindowPadW
		sizeH := w.Height()
		nextW += sizeW
		log(logNow, "redoWindows() start nextW,H =", nextW, nextH, "gocuiSize.W,H =", sizeW, sizeH, n.Name)

		if n.hasTabs {
			n.redoTabs(me.TabW, me.TabH)
		}
	}
}

func (p *node) redoTabs(nextW int, nextH int) {
	for _, n := range p.children {
		if n.WidgetType != toolkit.Tab {
			continue
		}
		w := n.tk
		w.frame = true

		n.gocuiSetWH(nextW, nextH)
		n.deleteView()
		// setCurrentTab(n)
		// if (len(w.cuiName) < 4) {
		// 	w.cuiName = "abcd"
		// }

		n.showView()

		sizeW := w.Width() + me.TabPadW
		sizeH := w.Height()
		log(logNow, "redoTabs() start nextW,H =", nextW, nextH, "gocuiSize.W,H =", sizeW, sizeH, n.Name)
		nextW += sizeW
	}
}
