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

	fmt.Println(flags)

	dashboard, err := meraki.NewDashboard(flags.ApiKey)
	if err != nil {
		log.Fatalf("error creating dashboard: %v\n", err)
	}

	orgs, err := organizations.GetOrganizations(dashboard)
	if err != nil {
		log.Fatalf("error getting organizations: %v\n", err)
	}

	admins := moma.GetAdminList(dashboard, orgs)

	fileName := moma.SetFileName(flags.OutputFile, "listAdmins")

	err = moma.WriteCsv(flags.OutputFile.Path, fileName, admins)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}

}
