package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
)

func main() {
	apiKey := flag.String("key", "", "Meraki API Key")

	flag.Parse()

	if *apiKey == "" {
		*apiKey = os.Getenv("MERAKI_API_KEY")
	}
	if *apiKey == "" {
		fmt.Println("API Key not defined, Meraki API Key must either be passed as a runtime flag or an Environmental Variable")
		fmt.Println("   To pass as a flag include `-key <api key>`")
		fmt.Println("   To pass as an environmental variable set `MERAKI_API_KEY` to your API key")
		os.Exit(1)
	}

	dashboard, err := meraki.NewDashboard(*apiKey)
	if err != nil {
		log.Fatalf("error creating dashboard: %v\n", err)
	}
	// fmt.Printf("Dashboard: %v\n", dashboard)

	orgs, _ := organizations.GetOrganizations(dashboard)

	for _, org := range orgs {
		orgAdmins, _ := organizations.GetOrganizationAdmins(dashboard, org.Id)
		for _, admin := range orgAdmins {
			// TODO not yielding expected results I don't think...
			fmt.Printf("ORG  %v    admin: %v\n", org.Name, admin.Name)
		}

	}

}
