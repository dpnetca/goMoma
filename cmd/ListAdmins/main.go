package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dpnetca/gomoma/pkg/meraki"
	"github.com/dpnetca/gomoma/pkg/meraki/organizations"
	"github.com/dpnetca/gomoma/pkg/outfile"
)

func main() {
	apiKey := handleFlags()

	dashboard, err := meraki.NewDashboard(apiKey)
	if err != nil {
		log.Fatalf("error creating dashboard: %v\n", err)
	}

	orgs, err := organizations.GetOrganizations(dashboard)
	if err != nil {
		log.Fatalf("error getting organizations: %v\n", err)
	}

	admins := getAdminList(dashboard, orgs)

	outfile.WriteOutput("listAdmins", admins)

}

func handleFlags() string {
	apiKey := flag.String("key", "", "Meraki API Key")
	// TODO: add flag for output directory
	// TODO: add flag for output filename

	flag.Parse()

	if *apiKey == "" {
		*apiKey = os.Getenv("MERAKI_API_KEY")
	}
	if *apiKey == "" {
		fmt.Println("API Key not defined, Meraki API Key must either be passed as a runtime flag or an Environmental Variable")
		fmt.Println("   To pass as a flag include `-key <api key>`")
		fmt.Println("   To pass as an environmental variable set `MERAKI_API_KEY` to your API key")
		os.Exit(1)
	}
	return *apiKey

}

func getAdminList(dashboard meraki.Dashboard, orgs []organizations.Organization) [][]string {
	admins := getAdminSyncMap(dashboard, orgs)
	return admins
}

func getAdminSyncMap(dashboard meraki.Dashboard, orgs []organizations.Organization) [][]string {
	var adminList [][]string
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for _, org := range orgs {
		wg.Add(1)
		go func(org organizations.Organization, wg *sync.WaitGroup, m *sync.Mutex) {
			admins, err := organizations.GetOrganizationAdmins(dashboard, org.Id)
			if err != nil {
				log.Printf("Error getting admins from %v: %v\n", org.Name, err)
			}
			var al [][]string
			for _, admin := range admins {
				l := []string{org.Id, org.Name, admin.Id, admin.Name, admin.Email, admin.OrgAccess}
				al = append(al, l)

			}
			m.Lock()
			adminList = append(adminList, al...)
			m.Unlock()
			wg.Done()
		}(org, wg, m)
	}
	wg.Wait()
	fmt.Printf("%v\n", adminList)
	return adminList
}
