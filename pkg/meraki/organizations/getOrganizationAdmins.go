package organizations

import (
	"encoding/json"

	"github.com/dpnetca/gomoma/pkg/meraki"
)

type Admin struct {
	Id        string
	Name      string
	Email     string
	OrgAccess string
}

func GetOrganizationAdmins(dashboard meraki.Dashboard, orgId string) ([]Admin, error) {
	endpoint := "/organizations/" + orgId + "/admins"
	data, err := meraki.SendRequest(dashboard, endpoint)
	if err != nil {
		return []Admin{}, err
	}
	var admins []Admin
	json.Unmarshal([]byte(data), &admins)

	return admins, nil
}
