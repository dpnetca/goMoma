package main

import (
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

	admins, err := moma.ReadCsv("removeAdmins.csv")
	if err != nil {
		log.Fatalln(err)
	}

	removedAdmins, err := moma.RemoveAdminsFromOrgs(dashboard, orgs, admins)
	if err != nil {
		log.Fatalln(err)
	}

	fileName := moma.SetFileName(flags.OutputFile, "removedAdmins")

	err = moma.WriteCsv(flags.OutputFile.Path, fileName, removedAdmins)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}

}
