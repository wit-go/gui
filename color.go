package gui

//
// convert between 'standard' golang Color and andlabs/ui Color
//

// import "log"
// import "fmt"
import "image/color"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func libuiColorToGOlangColor(rgba color.RGBA) ui.TableColor {
	/* a hack to see if colors work differently on macos or windows 
	if (rgba.R == 72) {
		log.Println("SETTING COLOR TO NIL")
		log.Println("SETTING COLOR TO NIL")
		log.Println("SETTING COLOR TO NIL")
		return ui.TableColor{}
	}
	*/
	return ui.TableColor{float64(rgba.R) / 256, float64(rgba.G) / 256, float64(rgba.B) / 256, float64(rgba.A) / 256}
}

/*
func golangColorGOlibuiColorTo (ui.TableColor) (rgba color.RGBA) {
	color.RGBA{float64(, 100, 200, 100}
	return ui.TableColor{float64(rgba.R) / 256, float64(rgba.G) / 256, float64(rgba.B) / 256, float64(rgba.A) / 256}
}
*/
