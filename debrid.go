package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var api_url = url.URL{
	Scheme: "https",
	Host:   "api.real-debrid.com",
	Path:   "/rest/",
}

type urlQueries struct {
	q     string
	value string
}

func (c *Client) rdGetUser() (rdUserSchema, error) {
	resBody, err := c.getReq("/user")
	if err != nil {
		return rdUserSchema{}, fmt.Errorf("Couldnt get user")
	}

	user := rdUserSchema{}
	if err := json.Unmarshal(resBody, &user); err != nil {
		fmt.Println("Decode failed")
		return rdUserSchema{}, err
	}

	return user, nil
}

func (c *Client) rdGetTorrents() ([]rdTorrentSchema, error) {
	resBody, err := c.getReq("/torrents")
	if err != nil {
		return nil, fmt.Errorf("Couldnt get user")
	}

	torrents := []rdTorrentSchema{}
	if err := json.Unmarshal(resBody, &torrents); err != nil {
		return nil, fmt.Errorf("Decode failed")
	}

	return torrents, nil

}

func (c *Client) rdAddMagnet(magnet string) {

}
