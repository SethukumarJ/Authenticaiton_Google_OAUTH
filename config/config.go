package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func SetupConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "429584661113-tvn5jtiuam6ojr5vhp4afadv7k7hci74.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/google/callback",
	}

}