/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gui_test

import (
	"git.wit.org/wit/gui"
)

// This example demonstrates how to create a NewWindow()
//
// Interacting with a GUI in a cross platform fashion adds some
// unusual problems. To obvuscate those, andlabs/ui starts a
// goroutine that interacts with the native gui toolkits
// on the Linux, MacOS, Windows, etc.
//
// Because of this oddity, to initialize a new window, the
// function is not passed any arguements and instead passes
// the information via the Config type.
//
func ExampleNewWindow() {
	// Define the name and size
	gui.Config.Title = "WIT GUI Window 1"
        gui.Config.Width = 640
        gui.Config.Height = 480

	// Create the Window
	gui.NewWindow()

	// Output:
	// You get a window
}
