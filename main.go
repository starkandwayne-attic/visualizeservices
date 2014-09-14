package main

import (
	"flag"
	"fmt"
	"net/http"

	consulapi "github.com/armon/consul-api"
)

func main() {
	var consulAddr = flag.String("consul-addr", "localhost:8500", "HTTP API for Consul agent/server")
	flag.Parse()

	config := consulapi.Config{Address: *consulAddr, HttpClient: http.DefaultClient}
	consul, _ := consulapi.NewClient(&config)
	query := consulapi.QueryOptions{}
	catalog := consul.Catalog()

	services, _, err := catalog.Services(&query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", services)
}
