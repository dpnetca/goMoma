package moma

import (
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
			err := organizations.CreateOrganizationAdmin(dashboard, org.Id, newAdmin)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
