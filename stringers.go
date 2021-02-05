// Exercise: Stringers
package main

import "fmt"

type IPAddr [4]byte

func (host IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", host[0], host[1], host[2], host[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
