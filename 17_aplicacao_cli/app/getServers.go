package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func getServers(c *cli.Context) {
	host := c.String("host")

	serversList, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Servers list of %s: \n", host)

	for _, server := range serversList {
		fmt.Printf("   [%s]\n", server.Host)
	}
}
