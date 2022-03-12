package main

import (
	"log"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
	"github.com/dpnetca/gomoma/pkg/moma"
	"github.com/dpnetca/gomoma/pkg/outfile"
)

func main() {
	apiKey := moma.HandleFlags()

	dashboard, err := meraki.NewDashboard(apiKey)
	if err != nil {
		log.Fatalf("error creating dashboard: %v\n", err)
	}

	orgs, err := organizations.GetOrganizations(dashboard)
	if err != nil {
		log.Fatalf("error getting organizations: %v\n", err)
	}

	admins := moma.GetAdminList(dashboard, orgs)

	outfile.WriteOutput("listAdmins", admins)

}
