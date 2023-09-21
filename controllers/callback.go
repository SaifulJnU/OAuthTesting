package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// RandomString should be defined somewhere in your code as the expected 'state' value.

func Callback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")

	data, err := getUserData(state, code)
	if err != nil {
		log.Fatal("Error getting user data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data:", string(data))
	fmt.Println(string(data))

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Fatal("Error parsing user data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	verifiedEmail, ok := dataMap["verified_email"].(bool)
	if !ok {
		log.Fatal("Error accessing 'verified_email' field")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println(verifiedEmail)
}

func getUserData(state, code string) ([]byte, error) {
	// Assuming 'RandomString' should be compared with 'state'.
	if state != RandomString {
		return nil, fmt.Errorf("invalid state value")
	}

	token, err := ssogolang.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// Construct the URL with the access token correctly.
	userInfoURL := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken

	response, err := http.Get(userInfoURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
