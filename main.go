package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	consulapi "github.com/armon/consul-api"
)

// ServerNodeServiceTally records the tally of services on a server node
type ServerNodeServiceTally struct {
	Address          string
	ServiceCharCodes string
}

var tally = make(map[string]*ServerNodeServiceTally)

func main() {
	var consulAddr = flag.String("consul-addr", "localhost:8500", "HTTP API for Consul agent/server (or $CONSUL_HTTP_ADDR)")
	flag.Parse()
	if os.Getenv("CONSUL_HTTP_ADDR") != "" {
		env := os.Getenv("CONSUL_HTTP_ADDR")
		consulAddr = &env
	}

	config := consulapi.Config{Address: *consulAddr, HttpClient: http.DefaultClient}
	consul, _ := consulapi.NewClient(&config)
	query := consulapi.QueryOptions{}
	catalog := consul.Catalog()

	catalogServices, _, err := catalog.Services(&query)
	if err != nil {
		fmt.Println(err)
		return
	}
	for catalogService, tags := range catalogServices {
		if len(tags) > 0 {
			// fmt.Printf("ServiceName: %s Tags: %#v\n", catalogService, tags)
			services, _, err := catalog.Service(catalogService, "", &query)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, service := range services {
				appendServiceToServerNodes(service.Node, serviceCharCodeForService(service))
				// fmt.Printf("%#v\n", *service)
			}

		}
	}
	for nodeName, nodeTally := range tally {
		fmt.Printf("%s: %s\n", nodeName, nodeTally.ServiceCharCodes)
	}
}

func appendServiceToServerNodes(nodeName string, serviceCharCode string) {
	if tally[nodeName] == nil {
		tally[nodeName] = &ServerNodeServiceTally{}
	}
	serverNode := tally[nodeName]
	serverNode.ServiceCharCodes = serverNode.ServiceCharCodes + serviceCharCode
}

func serviceCharCodeForService(service *consulapi.CatalogService) string {
	// mid-dot http://www.fileformat.info/info/unicode/char/b7/index.htm
	return "·"
}
