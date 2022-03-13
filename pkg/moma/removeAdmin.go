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
		"AdminName",
		"ErrorMessage",
	},
	)
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, org := range orgs {
		// TODO Get Admins for Org
		for _, admin := range admins {
			// TODO check if admin  email in admin list, get admin ID and continue if not exit loop
			wg.Add(1)
			go func(org organizations.Organization, admin []string, wg *sync.WaitGroup, m *sync.Mutex) {
				defer wg.Done()
				res, err := organizations.DeleteOrganizationAdmin(dashboard, org.Id, admin[0])
				if err != nil {
					fmt.Printf("error deleting admin %s from %s: %s", admin[0], org.Name, err)
				}
				m.Lock()
				removedAdmins = append(removedAdmins, []string{
					org.Id,
					org.Name,
					admin[0],
					strings.Join(res.ErrorMessage, ", "),
				})
			}(org, admin, wg, m)
		}
	}
	wg.Wait()
	sortedRemovedAdmins := SortSlicesWithHeader(removedAdmins, 1)

	return sortedRemovedAdmins, nil
}
