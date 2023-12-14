// This is a simple example
package main

import 	(
	"log"
	"strconv"
)

func loadDNS(c *configT) {
	hostname := c.domain
	log.Println("adding DNS record", hostname)

	newt := mainWindow.NewTab(hostname)
	newg := newt.NewGroup("more")

	// make a grid 6 things wide
	grid := newg.NewGrid("gridnuts", 6, gridH)

//	grid.NewButton("Type", func () {
//		log.Println("sort by Type")
//	})
	typedrop := grid.NewDropdown("type")
	typedrop.AddText("A")
	typedrop.AddText("AAAA")
	typedrop.AddText("CNAME")
	typedrop.Custom = func () {
		log.Println("custom dropdown() a =", typedrop.Name, typedrop.S)
	}
	nb := grid.NewButton("Name", func () {
		log.Println("sort by Name")
	})
	nb.Disable()

	grid.NewButton("Protection", func () {
		log.Println("sort proxied")
	})
	grid.NewButton("TTL", func () {
		log.Println("sort by TTL")
	})
	nb = grid.NewButton("Value", func () {
		log.Println("sort by Value")
	})
	nb.Disable()
	nb = grid.NewButton("Save", func () {
		log.Println("click below to save")
	})
	nb.Disable()

	masterSave = newt.NewButton("Master Save", func () {
		log.Println("save stuff to cloudflare")
	})
	masterSave.Disable()

	records := getZonefile(c)
	for _, record := range records.Result {
		var rr RRT // dns zonefile resource record

		// copy all the JSON values into the row record.
		rr.ID = record.ID
		rr.Type = record.Type
		rr.Name = record.Name
		rr.Content = record.Content
		rr.Proxied = record.Proxied
		rr.Proxiable = record.Proxiable
		rr.TTL = record.TTL

		rr.typeNode = grid.NewLabel(record.Type)
		rr.nameNode = grid.NewEntryLine(record.Name)
		rr.nameNode.SetText(record.Name)
		rr.nameNode.Disable()

		// set proxy or unproxied
		rr.proxyNode = grid.NewDropdown("proxy")
		if (record.Proxied) {
			rr.proxyNode.AddText("Proxied")
			rr.proxyNode.AddText("DNS")
		} else {
			rr.proxyNode.AddText("DNS")
			rr.proxyNode.AddText("Proxied")
		}
		rr.proxyNode.Custom = func () {
			log.Println("proxy dropdown() a =", rr.proxyNode.Name, rr.proxyNode.S, rr.ID)
			rr.saveNode.Enable()
			masterSave.Enable()
		}

		var ttl, short  string
		if (record.TTL == 1) {
			ttl = "Auto"
		} else {
			ttl = strconv.Itoa(record.TTL)
		}
		rr.ttlNode = grid.NewLabel(ttl)
		// short = fmt.Sprintf("%80s", record.Content)
		short = record.Content
		if len(short) > 40 {
			short = short[:40] // Slice the first 20 characters
		}

		rr.valueNode = grid.NewEntryLine(short)
		rr.valueNode.SetText(record.Content)

		rr.valueNode.Custom = func () {
			log.Println("value changed =", rr.valueNode.Name, rr.proxyNode.S, rr.ID)
			rr.saveNode.Enable()
			masterSave.Enable()
		}

		// fmt.Printf("ID: %s, Type: %s, Name: %s, short Content: %s\n", record.ID, record.Type, record.Name, short)
		// fmt.Printf("\tproxied: %b, %b, string TTL: %i\n", record.Proxied, record.Proxiable, ttl)

		rr.saveNode = grid.NewButton("Save", nil)
		rr.saveNode.Disable()
		rr.saveNode.Custom = func () {
			name := "save stuff to cloudflare for " + rr.ID
			log.Println(name)
			doChange(&rr)
		}
	}

	grid.Pad()
}
