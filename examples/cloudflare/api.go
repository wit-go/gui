// This is a simple example
package main

import 	(
	"os"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
//	"strconv"
	"bytes"

	"github.com/davecgh/go-spew/spew"
)

func doChange(dnsRow *RRT) {
	log.Println("Look for changes in row", dnsRow.ID)
	log.Println("Proxy", dnsRow.Proxied, "vs", dnsRow.proxyNode.S)
	log.Println("Content", dnsRow.Content, "vs", dnsRow.valueNode.S)
	if (dnsRow.Content != dnsRow.valueNode.S) {
		log.Println("UPDATE VALUE", dnsRow.nameNode.Name, dnsRow.typeNode.Name, "to", dnsRow.valueNode.S)
		stuff, result := httpPut(dnsRow)
		if (dnsRow.curlNode != nil) {
			pretty, _ := formatJSON(stuff)
			log.Println("http PUT curl =", pretty)
			dnsRow.curlNode.SetText(pretty)
		}
		if (dnsRow.resultNode != nil) {
			pretty, _ := formatJSON(result)
			log.Println("http PUT result =", pretty)
			dnsRow.resultNode.SetText(pretty)
		}
	}
	dnsRow.saveNode.Disable()
}

func getZonefile(c *configT) *DNSRecords {
	var url = cloudflareURL + c.zoneID + "/dns_records/"
	log.Println("getZonefile()", c.domain, url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http.NewRequest error:", err)
		return nil
	}

	// Set headers
	req.Header.Set("X-Auth-Key", c.auth)
	req.Header.Set("X-Auth-Email", c.email)

	log.Println("getZonefile() auth, email", c.auth, c.email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("http.Client error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() error", err)
		return nil
	}

	var records DNSRecords
	if err := json.Unmarshal(body, &records); err != nil {
		log.Println("json.Unmarshal() error", err)
		return nil
	}

	log.Println("getZonefile() worked", records)
	return &records
}

/*
	pass in a DNS Resource Records (the stuff in a zonefile)

	This will talk to the cloudflare API and generate a resource record in the zonefile:

	For example:
	gitea.wit.com. 3600 IN CNAME git.wit.com.
	go.wit.com. 3600 IN A 1.1.1.9
	test.wit.com. 3600 IN NS ns1.wit.com.
*/
func httpPut(dnsRow *RRT) (string, string) {
	var url string = cloudflareURL + os.Getenv("CF_API_ZONEID") + "/dns_records/" + dnsRow.ID
	var authKey string = os.Getenv("CF_API_KEY")
	var email string = os.Getenv("CF_API_EMAIL")

	// make a json record to send on port 80 to cloudflare
	var tmp string
	tmp = `{"content": "` + dnsRow.valueNode.S + `", `
	tmp += `"name": "` + dnsRow.Name + `", `
	tmp += `"type": "` + dnsRow.Type + `", `
	tmp+= `"ttl": "` +  "1" + `", `
	tmp += `"comment": "WIT DNS Control Panel"`
	tmp +=  `}`
	data := []byte(tmp)

	log.Println("http PUT url =", url)
	// log.Println("http PUT data =", data)
	// spew.Dump(data)
	pretty, _ := formatJSON(string(data))
	log.Println("http PUT data =", pretty)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Key", authKey)
	req.Header.Set("X-Auth-Email", email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return tmp, fmt.Sprintf("blah err =", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return tmp, fmt.Sprintf("blah err =", err)
	}
	// log.Println("http PUT body =", body)
	// spew.Dump(body)

	return tmp, string(body)
}

// https://api.cloudflare.com/client/v4/zones
func getZones(auth, email string) *DNSRecords {
	var url = "https://api.cloudflare.com/client/v4/zones"
	log.Println("getZones()", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http.NewRequest error:", err)
		return nil
	}

	// Set headers
	req.Header.Set("X-Auth-Key", auth)
	req.Header.Set("X-Auth-Email", email)

	log.Println("getZones() auth, email", auth, email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("getZones() http.Client error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("getZones() ioutil.ReadAll() error", err)
		return nil
	}

	var records DNSRecords
	if err := json.Unmarshal(body, &records); err != nil {
		log.Println("getZones() json.Unmarshal() error", err)
		return nil
	}

	/* Cloudflare API returns struct[] of:
	  struct { ID string "json:\"id\""; Type string "json:\"type\""; Name string "json:\"name\"";
		Content string "json:\"content\""; Proxied bool "json:\"proxied\"";
		Proxiable bool "json:\"proxiable\""; TTL int "json:\"ttl\"" }
	*/

	// log.Println("getZones() worked", records)
	// log.Println("spew dump:")
	spew.Dump(records)
	for _, record := range records.Result {
		log.Println("spew record:", record)
		log.Println("record:", record.Name, record.ID)

		var newc *configT
		newc = new(configT)

		newc.domain = record.Name
		newc.zoneID = record.ID
		newc.auth = auth
		newc.email = email

		config[record.Name] = newc
		zonedrop.AddText(record.Name)
		log.Println("zonedrop.AddText:", record.Name, record.ID)
	}
	for d, _ := range config {
		log.Println("config entry:", d)
	}

	return &records
}

// formatJSON takes an unformatted JSON string and returns a formatted version.
func formatJSON(unformattedJSON string) (string, error) {
	var jsonData interface{}

	// Decode the JSON string into an interface
	err := json.Unmarshal([]byte(unformattedJSON), &jsonData)
	if err != nil {
		return "", err
	}

	// Re-encode the JSON with indentation for formatting
	formattedJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return "", err
	}

	return string(formattedJSON), nil
}
