package moma

import (
	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
)

func AddAdminsToOrgs(
	dashboard meraki.Dashboard, orgs []organizations.Organization, admins [][]string,
) ([][]string, error) {

	var addedAdmins [][]string
	addedAdmins = append(addedAdmins, []string{
		"OrgID",
		"OrgName",
		"AdminID",
		"AdminName",
		"AdminEmail",
		"ErrorMessage",
	},
	)
	for _, org := range orgs {
		for _, admin := range admins {
			newAdmin := organizations.Admin{
				Name:      admin[0],
				Email:     admin[1],
				OrgAccess: admin[2],
			}
			res, err := organizations.CreateOrganizationAdmin(dashboard, org.Id, newAdmin)
			if err != nil {
				return [][]string{}, err
			}
			if res.Success {
				addedAdmins = append(addedAdmins, []string{
					org.Id,
					org.Name,
					res.Admin.Id,
					res.Admin.Name,
					res.Admin.Email,
					"",
				},
				)
			} else {
				for _, errMsg := range res.ErrorMessage {
					addedAdmins = append(addedAdmins, []string{
						org.Id,
						org.Name,
						"",
						newAdmin.Name,
						newAdmin.Email,
						errMsg,
					},
					)
				}
			}
		}
	}
	return addedAdmins, nil
}
