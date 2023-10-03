package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

var api_url = url.URL{
	Scheme: "https",
	Host:   "api.real-debrid.com",
	Path:   "/rest/",
}

func (c *Client) rdGetUser() (rdUserSchema, error) {
	resBody, err := c.getReq("/user")
	if err != nil {
		fmt.Println("Couldnt get user")
		return rdUserSchema{}, fmt.Errorf("Couldnt get user")
	}

	user := rdUserSchema{}
	if err := json.Unmarshal(resBody, &user); err != nil {
		fmt.Println("Decode failed")
		return rdUserSchema{}, fmt.Errorf("decode failed %w", err)
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

	fmt.Println(torrents)
	return torrents, nil

}

func (c *Client) rdAddMagnet(magnet string) (rdAddMagnetSchema, error) {

	data := url.Values{}
	data.Set("magnet", magnet)
	data.Set("host", "real-debrid.com")

	resBody, err := c.postReq("/torrents/addMagnet", data)
	if err != nil {
		return rdAddMagnetSchema{}, err
	}

	fmt.Println(string(resBody))

	mag := rdAddMagnetSchema{}
	if err := json.Unmarshal(resBody, &mag); err != nil {
		return rdAddMagnetSchema{}, fmt.Errorf("Decode failed")
	}

	return mag, nil
}

func (c *Client) rdGetFileInfo(id string) (rdAddMagnetSchema, error) {

	params := url.Values{}
	params.Set("id", id)

	resBody, err := c.getReqWithParams("/torrents/info", params)
	if err != nil {
		return rdAddMagnetSchema{}, err
	}

	mag := rdAddMagnetSchema{}
	if err := json.Unmarshal(resBody, &mag); err != nil {
		return rdAddMagnetSchema{}, fmt.Errorf("Decode failed")
	}

	return mag, nil
}
