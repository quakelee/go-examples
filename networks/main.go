// DNS resolve example in go

package main

import (
	"fmt"
	"net"
)

func main() {

	// DNS resolve
	namelist := []string{"www.google.com", "www.amazon.com", "www.facebook.com"}
	for _, n := range namelist {
		ip, _ := net.LookupHost(n)
		fmt.Println(net.LookupAddr(ip[0]))
	}
}
