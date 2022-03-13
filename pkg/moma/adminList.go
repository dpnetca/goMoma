package moma

import (
	"log"
	"sync"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
)

func GetAdminList(dashboard meraki.Dashboard, orgs []organizations.Organization) [][]string {
	admins := getAdminListSync(dashboard, orgs)
	return admins
}

func getAdminListSync(dashboard meraki.Dashboard, orgs []organizations.Organization) [][]string {
	var adminList [][]string
	adminList = append(adminList, []string{
		"OrgID",
		"OrgName",
		"AdminID",
		"AdminName",
		"AdminEmail",
		"AdminOrgAccess",
	},
	)
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for _, org := range orgs {
		wg.Add(1)
		// TODO consider using channel isntead of mutex
		go getOrgAdminList(dashboard, org, &adminList, wg, m)
	}
	wg.Wait()
	sortedAdminList := SortSlicesWithHeader(adminList, 1)
	return sortedAdminList
}

func getOrgAdminList(
	dashboard meraki.Dashboard,
	org organizations.Organization,
	adminList *[][]string,
	wg *sync.WaitGroup,
	m *sync.Mutex,
) {

	admins, err := organizations.GetOrganizationAdmins(dashboard, org.Id)
	if err != nil {
		log.Printf("Error getting admins from %v: %v\n", org.Name, err)
	}
	var al [][]string
	for _, admin := range admins {
		al = append(al, []string{org.Id, org.Name, admin.Id, admin.Name, admin.Email, admin.OrgAccess})

	}
	m.Lock()
	*adminList = append(*adminList, al...)
	m.Unlock()
	wg.Done()
}
