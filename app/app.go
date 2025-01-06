package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate new CLI application for execute
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP Shower"
	app.Usage = "Search for a domain IP and nameservers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find Ip's",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search nameservers",
			Flags:  flags,
			Action: searchNameservers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchNameservers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
