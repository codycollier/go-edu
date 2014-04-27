//
// host - resolve a dns name to an ip
//
// Lightly mimics the unix host command
//

// host resolves a dns name to ip(s)
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var host string

func main() {

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Printf("Usage: %s <hostname>\n", os.Args[0])
		os.Exit(1)
	}

	host := flag.Arg(0)

	iplist, err := net.LookupHost(host)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	for _, ip := range iplist {
		fmt.Printf("%s has address %s\n", host, ip)
	}

}
