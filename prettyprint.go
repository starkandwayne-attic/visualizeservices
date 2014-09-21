package main

import (
	"fmt"
	"net/http"

	"github.com/armon/consul-api"
)

func showPrettyVisualization() {
	config := consulapi.Config{Address: consulAddr, HttpClient: http.DefaultClient}
	consul, _ := consulapi.NewClient(&config)
	query := consulapi.QueryOptions{}
	catalog := consul.Catalog()

	catalogServices, _, err := catalog.Services(&query)
	if err != nil {
		fmt.Println(err)
		return
	}
	for catalogService, _ := range catalogServices {
		fmt.Printf("%s:\n", catalogService)

		services, _, err := catalog.Service(catalogService, "", &query)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, service := range services {
			fmt.Printf("  %v\n", *service)
		}
	}

}
