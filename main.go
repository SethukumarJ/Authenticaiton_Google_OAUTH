package main

import (
	"net/http"
	"auth/controllers"
)

func main() {

	http.HandleFunc("/google/login", controllers.GoogleLogin)
	http.HandleFunc("/google/callback", controllers.GoogleCallback)
}
