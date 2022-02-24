package organizations

import (
	"encoding/json"

	"github.com/dpnetca/gomoma/pkg/meraki"
)

func GetOrganizations(dashboard meraki.Dashboard) ([]Organization, error) {
	endpoint := "/organizations"
	data, err := meraki.SendRequest(dashboard, endpoint)
	if err != nil {
		return []Organization{}, err
	}
	var organizations []Organization
	json.Unmarshal([]byte(data), &organizations)

	return organizations, nil

}
