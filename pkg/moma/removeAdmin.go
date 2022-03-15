package moma

import (
	"fmt"
	"strings"
	"sync"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
)

func RemoveAdminsFromOrgs(
	dashboard meraki.Dashboard, orgs []organizations.Organization, admins [][]string,
) ([][]string, error) {
	var removedAdmins [][]string
	removedAdmins = append(removedAdmins, []string{
		"OrgID",
		"OrgName",
		"AdminId",
		"AdminName",
		"AdminEmail",
		"ErrorMessage",
	},
	)
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, org := range orgs {
		adminList, err := organizations.GetOrganizationAdmins(dashboard, org.Id)
		if err != nil {
			fmt.Printf("error getting admins from %s\n", org.Name)
		}
		for _, admin := range admins {
			var removeAdmin organizations.Admin
			for _, existingAdmin := range adminList {
				if existingAdmin.Email == admin[0] {
					removeAdmin = existingAdmin
					break
				}
			}
			if removeAdmin.Id == "" {
				continue
			}

			wg.Add(1)
			go func(org organizations.Organization, removeAdmin organizations.Admin, wg *sync.WaitGroup, m *sync.Mutex) {
				defer wg.Done()
				res, err := organizations.DeleteOrganizationAdmin(dashboard, org.Id, removeAdmin.Id)
				if err != nil {
					fmt.Printf("error deleting admin %s from %s: %s", removeAdmin, org.Name, err)
				}
				m.Lock()
				removedAdmins = append(removedAdmins, []string{
					org.Id,
					org.Name,
					removeAdmin.Id,
					removeAdmin.Name,
					removeAdmin.Email,
					strings.Join(res.ErrorMessage, ", "),
				})
				m.Unlock()
			}(org, removeAdmin, wg, m)
		}
	}
	wg.Wait()
	sortedRemovedAdmins := SortSlicesWithHeader(removedAdmins, 1)

	return sortedRemovedAdmins, nil
}
