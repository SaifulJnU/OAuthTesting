package controllers

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ssogolang *oauth2.Config
var RandomString = "random-string"

func init() {
	ssogolang = &oauth2.Config{
		RedirectURL:  "http://localhost:8081/callback",
		ClientID:     "254690736652-q3j6n6joh3u9ru63c46ruoh419pevqq3.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	url := ssogolang.AuthCodeURL(RandomString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
