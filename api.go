package main

import (
	"io/ioutil"
	"net/http"
)

var Api = map[string]string{ // Creating Key/Value API
	"artists":   "https://groupietrackers.herokuapp.com/api/artists",
	"locations": "https://groupietrackers.herokuapp.com/api/locations",
	"dates":     "https://groupietrackers.herokuapp.com/api/dates",
	"relation":  "https://groupietrackers.herokuapp.com/api/relation",
}

func callApi(url string) ([]byte, error) {
	var w http.ResponseWriter

	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
