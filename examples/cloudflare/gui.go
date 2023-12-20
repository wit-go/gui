// This is a simple example
package main

import 	(
	"log"
	"strconv"

	"git.wit.org/jcarr/control-panel-dns/cloudflare"
)

func loadDNS(c *configT) {
	hostname := c.domain
	log.Println("adding DNS record", hostname)

	newt := mainWindow.NewTab(hostname)
	vb := newt.NewBox("vBox", false)
	newg := vb.NewGroup("more zoneID = " + c.zoneID)

	// make a grid 6 things wide
	grid := newg.NewGrid("gridnuts", 6, gridH)

//	grid.NewButton("Type", func () {
//		log.Println("sort by Type")
//	})
	grid.NewLabel("RR type")
	grid.NewLabel("hostname")

	grid.NewLabel("Proxy")
	grid.NewLabel("TTL")
	grid.NewLabel("Value")
	grid.NewLabel("Save")

	masterSave = vb.NewButton("Master Save", func () {
		log.Println("save stuff to cloudflare")
	})
	masterSave.Disable()

	records := getZonefile(c)
	for _, record := range records.Result {
		var rr cloudflare.RRT // dns zonefile resource record

		// copy all the JSON values into the row record.
		rr.ID = record.ID
		rr.Type = record.Type
		rr.Name = record.Name
		rr.Content = record.Content
		rr.Proxied = record.Proxied
		rr.Proxiable = record.Proxiable
		// rr.Ttl = record.TTL

		grid.NewLabel(record.Type)
		grid.NewLabel(record.Name)

		proxy := grid.NewLabel("proxy")
		if (record.Proxied) {
			proxy.SetText("On")
		} else {
			proxy.SetText("Off")
		}

		var ttl  string
		if (record.TTL == 1) {
			ttl = "Auto"
		} else {
			ttl = strconv.Itoa(record.TTL)
		}
		grid.NewLabel(ttl)

		val := grid.NewLabel("Value")
		val.SetText(record.Content)

		load := grid.NewButton("Load", nil)
		load.Custom = func () {
			name := "save stuff to cloudflare for " + rr.ID
			log.Println(name)
			// doChange(&rr)
		}
	}

	grid.Pad()
}
