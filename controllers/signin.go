package controllers

import (
	"fmt"
	"net/http"

	"github.com/saifuljnu/testingOauth/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ssogolang *oauth2.Config
var RandomString = "random-string"

func SetSSOGolang() {
	ssogolang = &oauth2.Config{
		RedirectURL:  config.URL,
		ClientID:     config.ID,
		ClientSecret: config.Secret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	url := ssogolang.AuthCodeURL(RandomString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
