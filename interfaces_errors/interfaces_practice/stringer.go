package main

import "fmt"

type IPAddr [4]byte
type Hosts map[string]IPAddr

func (h Hosts) String() string {
	text := ""
	for name, ip := range h {
		text += fmt.Sprintf("%s: %d.%d.%d.%d\n", name, ip[0], ip[1], ip[2], ip[3])
	}
	return text
}

func main() {
	hosts := Hosts{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Print(hosts.String())
}
