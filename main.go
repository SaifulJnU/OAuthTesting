package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/saifuljnu/testingOauth/config"
	"github.com/saifuljnu/testingOauth/controllers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config.SetEnvionment()

	// Calling SetSSOGolang()
	controllers.SetSSOGolang()
}

func main() {

	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)
	http.HandleFunc("/signin", controllers.Signin)
	http.HandleFunc("/callback", controllers.Callback)
	http.ListenAndServe(":8081", nil)
}
