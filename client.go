// For making request to the server. This will mostly be used
// by the CLI

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var baseUrl = "http://localhost:4000"

// makeRequest actually make the request to the server using the provided url
// returns the JSON object as a dictionary
func makeRequest(url string) (map[string]interface{}, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// add headers
	req.Header = http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// hold json response
	var results map[string]interface{}

	json.NewDecoder(res.Body).Decode(&results)

	return results, nil
}

// MakeUpRequest makes a request to the server url path /up?stop=
// with the stop param
func MakeUpRequest(stop int) (map[string]interface{}, error) {
	// default stop to 5
	if stop < 1 {
		stop = 5
	}

	// build url
	stopStr := strconv.Itoa(stop)
	url := baseUrl + "/up?stop=" + stopStr

	return makeRequest(url)
}

// MakeDownRequest makes a request to the server url path /down?start=
// with the stop param
func MakeDownRequest(start int) (map[string]interface{}, error) {
	return nil, nil
}

func TestClient() {
	url := "https://ifconfig.co"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	FatalOnErr(err)

	req.Header = http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json"},
	}

	res, err := client.Do(req)
	FatalOnErr(err)

	defer res.Body.Close()

	// variable to hold json
	var results map[string]interface{}

	// decode into variable
	json.NewDecoder(res.Body).Decode(&results)

	// indent json
	j, _ := json.MarshalIndent(results, "", "    ")
	fmt.Printf("%s\n", results["country"])

	fmt.Print(string(j))
}
