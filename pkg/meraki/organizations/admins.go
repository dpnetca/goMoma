package organizations

import (
	"encoding/json"
	"fmt"

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
	response, err := meraki.SendGetRequest(dashboard, endpoint)
	if err != nil {
		return []Admin{}, err
	}
	var admins []Admin
	json.Unmarshal([]byte(response.Data), &admins)

	return admins, nil
}

func CreateOrganizationAdmin(dashboard meraki.Dashboard, orgId string, admin Admin) (string, error) {
	endpoint := "/organizations/" + orgId + "/admins"
	body, err := json.Marshal(admin)
	if err != nil {
		return "", err
	}
	response, err := meraki.SendPostRequest(dashboard, endpoint, body)
	if err != nil {
		return "", err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		var errorMessage map[string]interface{}
		json.Unmarshal([]byte(response.Data), &errorMessage)
		return fmt.Sprintf("error %d - %s", response.StatusCode, errorMessage["errors"]), nil
	}
	return "success", nil

}
