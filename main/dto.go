package main

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
}
