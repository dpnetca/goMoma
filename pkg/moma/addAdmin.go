package moma

import (
	"fmt"
	"sort"
	"strings"
	"sync"

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
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, org := range orgs {
		for _, admin := range admins {
			wg.Add(1)
			go func(org organizations.Organization, admin []string, wg *sync.WaitGroup, m *sync.Mutex) {
				defer wg.Done()
				newAdmin := organizations.Admin{
					Name:      admin[0],
					Email:     admin[1],
					OrgAccess: admin[2],
				}
				res, err := organizations.CreateOrganizationAdmin(dashboard, org.Id, newAdmin)
				if err != nil {
					fmt.Printf("error creating admin %s in %s: %s", newAdmin.Name, org.Name, err)
					// return [][]string{}, err
				}
				// if res.Success {
				// 	adminId := res.Admin.Id
				// 	adminName := res.Admin.Name
				// 	adminEmail := res.Admin.Email
				// 	errMsg := ""

				// } else {

				// 	adminId := res.Admin.Id
				// 	adminName := res.Admin.Name
				// 	adminEmail := res.Admin.Email
				// 	errMsg := res.ErrorMessage)

				// }
				m.Lock()
				addedAdmins = append(addedAdmins, []string{
					org.Id,
					org.Name,
					res.Admin.Id,
					newAdmin.Name,
					newAdmin.Email,
					strings.Join(res.ErrorMessage, ", "),
				},
				)
				m.Unlock()

			}(org, admin, wg, m)
		}
	}
	wg.Wait()
	sortedSlice := [][]string{addedAdmins[0]}
	addedAdmins = addedAdmins[1:]
	sort.Slice(addedAdmins, func(p, q int) bool {
		return addedAdmins[p][1] < addedAdmins[q][1]
	})
	sortedSlice = append(sortedSlice, addedAdmins...)

	return sortedSlice, nil
}
