package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// Load API KEY
func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading keys")
	}

	apiKey := os.Getenv("rd_api")
	return apiKey
}

func httpClient(api_key string) *Client {
	client := &http.Client{}
	return &Client{
		c:      client,
		apiKey: api_key,
	}
}

func (c *Client) getReq(endpoint string) ([]byte, error) {

	reqUrl := api_url.ResolveReference(&url.URL{Path: "1.0" + endpoint})
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.apiKey)

	resp, err := c.c.Do(req)
	if err != nil {
		fmt.Println("Error when sending get request")
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldnt read body")
	}

	return resBody, nil
}

func (c *Client) getReqWithParams(endpoint string, data url.Values) ([]byte, error) {

	reqUrl := api_url.ResolveReference(&url.URL{Path: "1.0" + endpoint, RawQuery: data.Encode()})
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.apiKey)

	resp, err := c.c.Do(req)
	if err != nil {
		fmt.Println("Error when sending get request")
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldnt read body")
	}

	return resBody, nil
}

func (c *Client) postReq(endpoint string, data url.Values) ([]byte, error) {
	reqUrl := api_url.ResolveReference(&url.URL{Path: "1.0" + endpoint, RawQuery: data.Encode()})

	mod, err := url.QueryUnescape(reqUrl.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(mod)

	req, err := http.NewRequest("POST", mod, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.c.Do(req)
	if err != nil {
		fmt.Println("Error when sending post request")
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldnt read body")
	}

	return resBody, nil
}
