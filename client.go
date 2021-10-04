package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeRequest(start int, stop int) ([]int, error) {
	return []int{1, 2}, nil
}

func TestClient() {
	url := "https://ifconfig.co"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		FatalOnErr(err)
	}

	req.Header = http.Header{
		"Accept": []string{"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		FatalOnErr(err)
	}

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
