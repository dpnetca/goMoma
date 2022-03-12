package moma

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
)

func AddAdminsToOrgs(dashboard meraki.Dashboard, orgs []organizations.Organization, admins [][]string) error {
	for _, org := range orgs {
		for _, admin := range admins {
			newAdmin := organizations.Admin{
				Name:      admin[0],
				Email:     admin[1],
				OrgAccess: admin[2],
			}
			fmt.Printf("add %s to %s\n", admin[0], org.Name)
			err := organizations.CreateOrganizationAdmin(dashboard, org.Id, newAdmin)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ReadAdminCsv(file string) ([][]string, error) {
	adminFile, err := os.Open(file)
	if err != nil {
		return [][]string{}, fmt.Errorf("unable to open file %s", err)
	}
	defer adminFile.Close()

	r := csv.NewReader(adminFile)
	r.Comment = '#'
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, fmt.Errorf("unable to parse file %s", err)
	}

	return records, nil
}
