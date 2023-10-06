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

/* endpoint /user
* Params: 
* returns: All the details about the user in Json
*/

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

/* endpoint /torrents
* Params: 
* returns: all the torrents added by the user 
*/

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

/* endpoint /torrents/addMagnet
* Params: magnet link string
* returns: id and url of the torrent added
*/

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

/* endpoint /torrents/info/{id}.
* Params: Id of the torrrent whose info is needed
* returns: all the details of the torrent in json format
*/

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

/* endpoint /torrents/selectFiles/{id}.
* Params: Id of the torrent, we can get id from /torrents/info
* returns: Nothing
*/
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

/* endpoint /unrestrict/link.
* Params: Link to the torrent
* returns: Unrestricted real-debrid link which can be downloaded by aria
*/
func (c *Client) rdUnrestrictLinks(link string) (UnrestrictLink, error){
  data := url.Values{}
  data.Set("link", link)

  resp, err := c.postReq("/unrestrict/link", data)
  if err != nil {
    return UnrestrictLink{}, err
  }

  getLink := UnrestrictLink{}
	if err := json.Unmarshal(resp, &getLink); err != nil {
		return UnrestrictLink{}, fmt.Errorf("Decode failed")
	}

	return getLink, nil
}

// Helper function. TODO move to helper.go
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
