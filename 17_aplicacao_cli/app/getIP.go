package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func getIPs(c *cli.Context) {
	host := c.String("host")

	//net
	ipList, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP list of %s: \n", host)

	for _, ip := range ipList {
		fmt.Printf("   [%s]\n", ip)
	}
}
