// This is a simple example
package main

import 	(
	"os"
	"log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Define a struct to match the JSON structure of the response.
// This structure should be adjusted based on the actual format of the response.
type DNSRecords struct {
	Result []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Content string `json:"content"`
		Proxied bool `json:"proxied"`
		Proxiable bool `json:"proxiable"`
		TTL int `json:"ttl"`
	} `json:"result"`
}

// var domain string = "wit.org"
// var os.Getenv("CLOUDFLARE_DOMAIN")

func loadDNS(hostname string) {
	log.Println("adding DNS record")

	newt := mainWindow.NewTab(hostname)
	newg := newt.NewGroup("more")
	grid := newg.NewGrid("gridnuts", 5, gridH)

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
	grid.NewButton("Name", func () {
		log.Println("sort by Name")
	})
	grid.NewButton("Protection", func () {
		log.Println("sort proxied")
	})
	grid.NewButton("TTL", func () {
		log.Println("sort by TTL")
	})
	grid.NewButton("Value", func () {
		log.Println("sort by Value")
	})

	newt.NewButton("Save", func () {
		log.Println("save stuff to cloudflare")
	})

	records := getRecords()
	for _, record := range records.Result {
		grid.NewLabel(record.Type)
		textbox := grid.NewTextbox(record.Name)
		textbox.SetText(record.Name)
		if (record.Proxied) {
			grid.NewLabel("Proxied")
		} else {
			grid.NewLabel("DNS")
		}
		var ttl, short  string
		if (record.TTL == 1) {
			ttl = "Auto"
		} else {
			ttl = strconv.Itoa(record.TTL)
		}
		grid.NewLabel(ttl)
		// short = fmt.Sprintf("%80s", record.Content)
		short = record.Content
		if len(short) > 40 {
			short = short[:40] // Slice the first 20 characters
		}

		namebox := grid.NewTextbox(short)
		namebox.SetText(short)

		fmt.Printf("ID: %s, Type: %s, Name: %s, short Content: %s\n", record.ID, record.Type, record.Name, short)
		fmt.Printf("\tproxied: %b, %b, string TTL: %i\n", record.Proxied, record.Proxiable, ttl)
	}
}

func getRecords() *DNSRecords {
	var url string = os.Getenv("CLOUDFLARE_URL")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var authKey string = os.Getenv("CLOUDFLARE_AUTHKEY")
	var email string = os.Getenv("CLOUDFLARE_EMAIL")

	// Set headers
	req.Header.Set("X-Auth-Key", authKey)
	req.Header.Set("X-Auth-Email", email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var records DNSRecords
	if err := json.Unmarshal(body, &records); err != nil {
		fmt.Println(err)
		return nil
	}

	return &records
}
