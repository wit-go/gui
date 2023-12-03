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

	// more2.NewButton(name, func () {
	// 	log.Println(name, "ip =", ip)
	// })

	newt := mainWindow.NewTab(hostname)
	newg := newt.NewGroup("more")
	more2 := newg.NewGrid("gridnuts", 5, gridH)

	records := getRecords()
	for _, record := range records.Result {
		more2.NewLabel(record.Type)
		more2.NewLabel(record.Name)
		if (record.Proxied) {
			more2.NewLabel("Proxied")
		} else {
			more2.NewLabel("DNS")
		}
		var ttl, short  string
		if (record.TTL == 1) {
			ttl = "Auto"
		} else {
			ttl = strconv.Itoa(record.TTL)
		}
		more2.NewLabel(ttl)
		// short = fmt.Sprintf("%80s", record.Content)
		short = record.Content
		if len(short) > 40 {
			short = short[:40] // Slice the first 20 characters
		}
		more2.NewLabel(short)

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

    // Process the records as needed
    /*
    for _, record := range records.Result {
        fmt.Printf("ID: %s, Type: %s, Name: %s, Content: %s\n", record.ID, record.Type, record.Name, record.Content)
	fmt.Printf("\tproxied: %b, %b, TTL: %i\n", record.Proxied, record.Proxiable, record.TTL)
    }
    */

	return &records
}
