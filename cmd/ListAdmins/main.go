package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
	outfile "github.com/dpnetca/gomoma/pkg/outFile"
)

func main() {
	apiKey := handleFlags()

	dashboard, err := meraki.NewDashboard(apiKey)
	if err != nil {
		log.Fatalf("error creating dashboard: %v\n", err)
	}

	orgs, _ := organizations.GetOrganizations(dashboard)

	var adminList [][]string
	// set csv headers
	adminList = append(
		adminList,
		[]string{
			"org_id",
			"org_name",
			"admin_id",
			"admin_name",
			"admin_email",
			"admin_access",
		},
	)
	for _, org := range orgs {
		orgAdmins, _ := organizations.GetOrganizationAdmins(dashboard, org.Id)
		for _, admin := range orgAdmins {
			l := []string{org.Id, org.Name, admin.Id, admin.Name, admin.Email, admin.OrgAccess}
			adminList = append(adminList, l)
		}

	}

	outfile.WriteOutput("listAdmins", adminList)

}

func handleFlags() string {
	apiKey := flag.String("key", "", "Meraki API Key")
	// TODO: add flag for output directory
	// TODO: add flag for output filename

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
	return *apiKey

}
