package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/armon/consul-api"
)

func showDotVisualization() {
	config := consulapi.Config{Address: consulAddr, HttpClient: http.DefaultClient}
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
				appendServiceToServerNodes(service.Node, serviceCharCodeForServiceNode(service))
				// fmt.Printf("%#v\n", *service)
			}

		}
	}
	// TODO: sort tally by nodeName
	orderedNodeNames := []string{}
	maxNodeNameSize := 0
	for nodeName := range tally {
		orderedNodeNames = append(orderedNodeNames, nodeName)
		if len(nodeName) > maxNodeNameSize {
			maxNodeNameSize = len(nodeName)
		}
	}
	sort.Strings(orderedNodeNames)
	nodeNameFormatter := fmt.Sprintf("%%%ds: ", maxNodeNameSize)

	// TODO: fixed-width for displaying nodeName: so all .... are aligned
	for _, nodeName := range orderedNodeNames {
		fmt.Printf(nodeNameFormatter, nodeName)

		nodeTally := tally[nodeName]
		sort.Strings(nodeTally.ServiceCharCodes)
		for _, charCode := range nodeTally.ServiceCharCodes {
			fmt.Print(charCode)
		}
		fmt.Println("")
	}
}

func appendServiceToServerNodes(nodeName string, serviceCharCode string) {
	if tally[nodeName] == nil {
		tally[nodeName] = &ServerNodeServiceTally{}
	}
	serverNode := tally[nodeName]
	serverNode.ServiceCharCodes = append(serverNode.ServiceCharCodes, serviceCharCode)
}

func serviceCharCodeForServiceNode(service *consulapi.CatalogService) string {
	// default character is the mid-dot http://www.fileformat.info/info/unicode/char/b7/index.htm
	charCode := "·"
	primaryNode := false

	for _, tag := range service.ServiceTags {
		if strings.HasPrefix(tag, "char-code-") {
			charCode = strings.TrimPrefix(tag, "char-code-")
		}
		if tag == "master" || tag == "primary" {
			primaryNode = true
		}
	}

	if primaryNode {
		charCode = strings.ToUpper(charCode)
	}

	return charCode
}
