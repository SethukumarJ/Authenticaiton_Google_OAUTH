package controllers

import (
	"auth/config"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GoogleController operations for Google

func GoogleLogin(res http.ResponseWriter, req *http.Request) {

	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	// redirect to google login page
	http.Redirect(res, req, url, http.StatusSeeOther)

}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {

	state := req.URL.Query()["state"][0]

	if state != "randomstate" {
		http.Error(res, "State is not valid", http.StatusBadRequest)
		return
	}

	code := req.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)

	if err != nil {
		fmt.Fprintln(res, err.Error())
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Fprintln(res, "user data fetch failed")
		return
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "user data read failed")
		return
	}

	fmt.Fprintln(res, string(userData))

}
