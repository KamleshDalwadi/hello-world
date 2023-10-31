package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Commits struct {
	Sha string `json:"sha"`
}

func main() {
	url := "https://api.github.com/repos/KamleshDalwadi/hello-world/commits?since=2023-01-01&until=2023-10-27"
	// url := "https://api.github.com/repos/KamleshDalwadi/RESTServices/commits"

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		return
	}

	// Set headers
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	// req.Header.Set("Authorization", "<token>") //add github token here
	// scoped token
	req.Header.Set("Authorization", "<token>") //add github token here
	req.Header.Set("Accept", "application/vnd.github+json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println(string(body))
	jsn := []Commits{}
	er := json.Unmarshal(body, &jsn)
	if er != nil {
		fmt.Printf("er: %v\n", er)
	}
	fmt.Printf("jsn: %v\n", len(jsn))
	for _, js := range jsn {
		fmt.Printf("js: %v\n", js.Sha)
	}
}
