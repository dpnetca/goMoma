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

type Response struct {
	StatusCode int
	Data       string
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

func SendPostRequest(dashboard Dashboard, endpoint string, body []byte) (Response, error) {
	url := dashboard.BaseURL + endpoint
	response, err := sendRequest(http.MethodPost, url, dashboard.ApiKey, bytes.NewBuffer(body))
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func SendGetRequest(dashboard Dashboard, endpoint string) (Response, error) {
	url := dashboard.BaseURL + endpoint
	response, err := sendRequest(http.MethodGet, url, dashboard.ApiKey, nil)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func sendRequest(method, url, apiKey string, body io.Reader) (Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{}, err
	}
	req.Header.Add("X-Cisco-Meraki-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	response := Response{
		StatusCode: res.StatusCode,
		Data:       string(data),
	}

	return response, nil

}
