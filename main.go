package main

import (
	"os"

	"github.com/codegangsta/cli"
)

// ServerNodeServiceTally records the tally of services on a server node
type ServerNodeServiceTally struct {
	Address          string
	ServiceCharCodes []string
}

var tally = make(map[string]*ServerNodeServiceTally)
var consulAddr = "localhost:8500"

func main() {
	app := cli.NewApp()
	app.Name = "visualizeservices"
	app.Usage = "Visualize Consul services being advertised"
	app.Commands = []cli.Command{
		{
			Name:      "dots",
			ShortName: "d",
			Usage:     "Visualize using dots or characters per service node",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "consul-addr, a",
					Value:  "localhost:8500",
					Usage:  "HTTP API for Consul agent/server",
					EnvVar: "CONSUL_HTTP_ADDR",
				},
			},
			Action: func(c *cli.Context) {
				consulAddr = c.String("consul-addr")
				showDotVisualization()
			},
		},
		{
			Name:      "pretty",
			ShortName: "p",
			Usage:     "Pretty-print all services being advertised on Consul",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "consul-addr, a",
					Value:  "localhost:8500",
					Usage:  "HTTP API for Consul agent/server",
					EnvVar: "CONSUL_HTTP_ADDR",
				},
				cli.StringFlag{
					Name:  "filter-name-prefix, n",
					Value: "",
					Usage: "Filter out services that commence with a prefix",
				},
				cli.StringFlag{
					Name:  "filter-tag, t",
					Value: "",
					Usage: "Filter out services that include specific tag",
				},
			},
			Action: func(c *cli.Context) {
				consulAddr = c.String("consul-addr")
				showPrettyVisualization(c.String("filter-name-prefix"), c.String("filter-tag"))
			},
		},
	}
	app.Run(os.Args)

}
