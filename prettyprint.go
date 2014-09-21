package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/armon/consul-api"
)

func showPrettyVisualization(filterNamePrefix, filterTag string) {
	config := consulapi.Config{Address: consulAddr, HttpClient: http.DefaultClient}
	consul, _ := consulapi.NewClient(&config)
	query := consulapi.QueryOptions{}
	catalog := consul.Catalog()

	catalogServices, _, err := catalog.Services(&query)
	if err != nil {
		fmt.Println(err)
		return
	}
	for catalogService, serviceTags := range catalogServices {
		skipService := strings.HasPrefix(catalogService, filterNamePrefix)
		if skipService {
			continue
		}

		for _, tag := range serviceTags {
			if tag == filterTag {
				skipService = true
				break
			}
		}
		if skipService {
			continue
		}

		fmt.Printf("%s:\n", catalogService)

		services, _, err := catalog.Service(catalogService, "", &query)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, service := range services {
			fmt.Printf("  %s:%d - %v\n", service.Node, service.ServicePort, service.ServiceTags)
		}
	}

}
