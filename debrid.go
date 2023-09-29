package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const api_url = "https://api.real-debrid.com/rest/1.0"

type urlQueries struct {
	q     string
	value string
}

// Load API KEY
func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading keys")
	}

	apiKey := os.Getenv("rd_api")
	return apiKey
}

func urlBuilder(endpoint string, query []*urlQueries) string {
	if query != nil {
		var qArray []string
		for _, i := range query {
			q := fmt.Sprintf("%s=%s", i.q, i.value)
			qArray = append(qArray, q)
		}
		q := strings.Join(qArray, "&")
		return api_url + "/" + endpoint + "?" + q
	}

	return api_url + "/" + endpoint
}

func rdGetUser() {
	client := &http.Client{}
	reqUrl := urlBuilder("user", []*urlQueries{{q: "auth_token", value: getApiKey()}})
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request")
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

func rdGetTorrents() {

}

func rdAddMagnet() {

}
