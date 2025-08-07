package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// "Gerar" returns a new CLI application instance. (Command Line Interface)
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Aplicação"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

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
			Usage:  "Busca o IP de endereços de servidor",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca os nomes de servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, err := net.LookupNS(host) // "NS" para "Name Server"
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
