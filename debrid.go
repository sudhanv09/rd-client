package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

	resBody, err := c.postReq("/torrents/addMagnet", data)
	if err != nil {
		return rdAddMagnetSchema{}, err
	}

	mag := rdAddMagnetSchema{}
	if err := json.Unmarshal(resBody, &mag); err != nil {
		return rdAddMagnetSchema{}, fmt.Errorf("Decode failed")
	}

	return mag, nil
}

func (c *Client) rdGetFileInfo(id string) (rdTorrentInfoSchema, error) {

	resBody, err := c.getReq(fmt.Sprintf("/torrents/info/%s", id))
	if err != nil {
		return rdTorrentInfoSchema{}, err
	}

	fileInfo := rdTorrentInfoSchema{}
	if err := json.Unmarshal(resBody, &fileInfo); err != nil {
		return rdTorrentInfoSchema{}, fmt.Errorf("Decode failed")
	}

	return fileInfo, nil
}

func (c *Client) rdSelectFiles(id string) error {

	torrentFiles, err := c.rdGetFileInfo(id)
	if err != nil {
		fmt.Errorf("Couldnt get files from the torrent")
	}

	files := getFileIdsFromTorrent(torrentFiles)

	data := url.Values{}
	data.Set("files", files)
	req, err := c.postReq("/torrents/selectFiles/"+id, data)

	fmt.Println(string(req))

	if err != nil {
		fmt.Errorf("Couldnt make the request")
	}
	return nil
}

func getFileIdsFromTorrent(val rdTorrentInfoSchema) string {
	allowedFileTypes := []string{"mkv", "srt"}
	var fileIds []string

	for _, val := range val.Files {
		for _, id := range allowedFileTypes {
			if strings.Contains(val.Path, id) {
				fileIds = append(fileIds, strconv.Itoa(val.Id))
			}
		}
	}
	downloadFiles := strings.Join(fileIds, ",")

	return downloadFiles
}
