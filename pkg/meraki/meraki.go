package meraki

import (
	"errors"
	"io"
	"net/http"
)

type Dashboard struct {
	ApiKey  string
	BaseURL string
}

func NewDashboard(apiKey string) (Dashboard, error) {
	if apiKey == "" {
		return Dashboard{}, errors.New("required API key not defined")
	}
	dashboard := Dashboard{
		ApiKey:  apiKey,
		BaseURL: "https://api.meraki.com/api/v1",
	}

	return dashboard, nil
}

func SendRequest(dashboard Dashboard, endpoint string) (string, error) {
	url := dashboard.BaseURL + "/organizations"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Cisco-Meraki-API-Key", dashboard.ApiKey)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)
	return string(data), nil
}
