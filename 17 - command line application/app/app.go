package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// "Generate" returns a new CLI application instance. (Command Line Interface)
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search IPs and Server Names on the Internet"

	// Define the flags for the application
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	// Slice of commands to be added to the application
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP addresses of servers",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "servers",
			Usage:  "Search server names",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host) // "NS" for "Name Server"
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
