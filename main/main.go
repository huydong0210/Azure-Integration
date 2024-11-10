package main

import "fmt"

func main() {

	//config in https://portal.azure.com/
	cfg := AzureConfig{
		ClientId:     "client_id",
		ClientSecret: "client_secret",
		GrantType:    "grant_type",
		Scope:        "scope", // value = "https://graph.microsoft.com/.default"
		TenantId:     "tenantId",
	}
	accessToken, err := GetAzureAccessToken(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(accessToken.AccessToken)
}
