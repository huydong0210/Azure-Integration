package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type AzureConfig struct {
	ClientId     string
	ClientSecret string
	TenantId     string
	GrantType    string
	Scope        string
}

func GetAzureAccessToken(azureConfig AzureConfig) (AccessTokenResponse, error) {
	formData := url.Values{}
	formData.Set("grant_type", azureConfig.GrantType)
	formData.Set("client_id", azureConfig.ClientId)
	formData.Set("client_secret", azureConfig.ClientSecret)
	formData.Set("scope", azureConfig.Scope)

	reqBody := strings.NewReader(formData.Encode())
	url := BuildAzureAccessTokenEndpoint(azureConfig)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer resp.Body.Close()

	var accessTokenResponse AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenResponse); err != nil {
		return AccessTokenResponse{}, err
	}
	return accessTokenResponse, nil
}

func BuildAzureAccessTokenEndpoint(azureConfig AzureConfig) string {
	return "https://login.microsoftonline.com/" + azureConfig.TenantId + "/oauth2/v2.0/token"

}
