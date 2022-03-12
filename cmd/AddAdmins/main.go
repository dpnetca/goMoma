package main

import (
	"fmt"
	"log"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
	"github.com/dpnetca/gomoma/pkg/moma"
)

func main() {
	flags := moma.HandleFlags()

	dashboard, err := meraki.NewDashboard(flags.ApiKey)
	if err != nil {
		log.Fatalf("error creating dashboard: %v\n", err)
	}

	orgs, err := organizations.GetOrganizations(dashboard)
	if err != nil {
		log.Fatalf("error getting organizations: %v\n", err)
	}

	admins, err := moma.ReadAdminCsv("newAdmins.csv")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(admins)

	err = moma.AddAdminsToOrgs(dashboard, orgs, admins)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO Add Error Handling
	// TODO Add output logging

}
