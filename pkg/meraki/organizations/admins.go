package organizations

import (
	"encoding/json"

	"github.com/dpnetca/gomoma/pkg/meraki"
)

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

func CreateOrganizationAdmin(dashboard meraki.Dashboard, orgId string, admin Admin) (AdminResponse, error) {
	endpoint := "/organizations/" + orgId + "/admins"
	adminResponse := AdminResponse{Success: false}

	body, err := json.Marshal(admin)
	if err != nil {
		return AdminResponse{}, err
	}
	response, err := meraki.SendPostRequest(dashboard, endpoint, body)
	if err != nil {
		return AdminResponse{}, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = json.Unmarshal([]byte(response.Data), &adminResponse)
		if err != nil {
			return AdminResponse{}, err
		}
		return adminResponse, nil
	}

	err = json.Unmarshal([]byte(response.Data), &adminResponse.Admin)
	if err != nil {
		return AdminResponse{}, err
	}

	adminResponse.Success = true
	return adminResponse, nil

}
