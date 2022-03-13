package moma

import (
	"fmt"

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
			res, err := organizations.CreateOrganizationAdmin(dashboard, org.Id, newAdmin)
			if err != nil {
				return err
			}
			if res.Success {
				fmt.Printf("Added org: %s, admin %s, ID: %s\n", org.Name, newAdmin.Name, res.Admin.Id)
			} else {
				for _, errMsg := range res.ErrorMessage {
					fmt.Printf("ERROR org: %s, admin %s, %s\n", org.Name, newAdmin.Name, errMsg)
				}
			}
		}
	}
	return nil
}
