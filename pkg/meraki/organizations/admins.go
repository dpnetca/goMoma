package organizations

import (
	"encoding/json"

	"github.com/dpnetca/gomoma/pkg/meraki"
)

type Admin struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	OrgAccess string `json:"orgAccess"`
}

func GetOrganizationAdmins(dashboard meraki.Dashboard, orgId string) ([]Admin, error) {
	endpoint := "/organizations/" + orgId + "/admins"
	data, err := meraki.SendGetRequest(dashboard, endpoint)
	if err != nil {
		return []Admin{}, err
	}
	var admins []Admin
	json.Unmarshal([]byte(data), &admins)

	return admins, nil
}

func CreateOrganizationAdmin(dashboard meraki.Dashboard, orgId string, admin Admin) error {
	endpoint := "/organizations/" + orgId + "/admins"
	body, err := json.Marshal(admin)
	if err != nil {
		return err
	}
	_, err = meraki.SendPostRequest(dashboard, endpoint, body)
	if err != nil {
		return err
	}
	return nil

}
