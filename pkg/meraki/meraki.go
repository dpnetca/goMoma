package meraki

import (
	"bytes"
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

func SendPostRequest(dashboard Dashboard, endpoint string, body []byte) (string, error) {
	url := dashboard.BaseURL + endpoint
	data, err := sendRequest(http.MethodPost, url, dashboard.ApiKey, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	return data, nil
}

func SendGetRequest(dashboard Dashboard, endpoint string) (string, error) {
	url := dashboard.BaseURL + endpoint
	data, err := sendRequest(http.MethodGet, url, dashboard.ApiKey, nil)
	if err != nil {
		return "", err
	}
	return data, nil
}

func sendRequest(method, url, apiKey string, body io.Reader) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Cisco-Meraki-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	// this is a breaking change, needs more thought...
	// if response.StatusCode < 200 || response.StatusCode > 299 {
	// 	return "", fmt.Errorf("request failed %s", data)
	// }

	return string(data), nil

}
