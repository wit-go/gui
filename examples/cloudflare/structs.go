// This is a simple example
package main

import 	(
	"git.wit.org/wit/gui"
)

var cloudflareURL string = "https://api.cloudflare.com/client/v4/zones/"

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

var masterSave *gui.Node

var domainWidget *gui.Node
var zoneWidget *gui.Node
var authWidget *gui.Node
var emailWidget *gui.Node

var loadButton *gui.Node
var saveButton *gui.Node
var zonedrop *gui.Node

// Resource Record (used in a DNS zonefile)
type RRT struct {
	typeNode *gui.Node	// CNAME, A, AAAA, ...
	nameNode *gui.Node	// www, mail, ...
	proxyNode *gui.Node	// If cloudflare is a port 80 & 443 proxy
	ttlNode *gui.Node	// just set to 1 which means automatic to cloudflare
	valueNode *gui.Node	// 4.2.2.2, "dkim stuff", etc
	curlNode *gui.Node	// shows you what you could run via curl
	resultNode *gui.Node	// what the cloudflare API returned
	saveNode *gui.Node	// button to send it to cloudflare

	ID     string
	Type   string
	Name   string
	Content string
	ProxyS string
	Proxied bool
	Proxiable bool
	Ttl string
}

/*
	This is a structure of all the RR's (Resource Records)
	in the DNS zonefiile for a hostname. For example:

	For the host test.wit.com:

	test.wit.com A 127.0.0.1
	test.wit.com AAAA
	test.wit.com TXT email test@wit.com
	test.wit.com TXT phone 212-555-1212
	test.wit.com CNAME real.wit.com
*/
type hostT struct {
	hostname string
	RRs []configT
}

type configT struct {
	domain string
	zoneID string
	auth string
	email string
}

var config map[string]*configT
