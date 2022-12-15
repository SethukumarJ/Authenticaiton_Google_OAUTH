package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func SetupConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "1011798649072-1u1m8152g461di4glig0hr30vfgpcj3o.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "https://localhost:3000/callback",

		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},

		Endpoint: google.Endpoint,

	}

}