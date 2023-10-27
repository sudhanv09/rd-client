package main

import (
	"fmt"
	"github.com/siku2/arigo"
	"golang.design/x/clipboard"
)

func main() {
	client := httpClient(getApiKey())
	aria, err := arigo.Dial("ws://localhost:6800/jsonrpc", "")
	if err != nil {
		fmt.Println("Aria is not running. Run using aria2c --enable-rpc")
	}

	cliperr := clipboard.Init()
	if cliperr != nil {
		fmt.Println("Couldnt init clipboard")
	}

	for {
		magLink := string(clipboard.Read(clipboard.FmtText))

		if magLink != "" && isMagnetValid(magLink) {
			fmt.Println("Found magnet. Adding....")
			magnet, err := client.rdAddMagnet(magLink)
			if err != nil {
				fmt.Println("Couldnt add magnet link")
			}

			selErr := client.rdSelectFiles(magnet.Id)
			if selErr != nil {
				fmt.Println("Failed to choose files in torrent")
			}

			getLinks, err := client.rdGetFileInfo(magnet.Id)
			if err != nil {
				fmt.Println("Couldnt get file info")
			}

			for idx, link := range getLinks.Links {
				_ = idx
				unrestrict, err := client.rdUnrestrictLinks(link)
				if err != nil {
					fmt.Println("Couldnt unrestrict links")
				}

				url := arigo.URIs(unrestrict.Download)

				status, err := aria.AddURI(url, &arigo.Options{Dir: "/hdd/media/aria"})
				if err != nil {
					fmt.Println("Couldnt add links to aria")
				}
				fmt.Println(status.GID)
			}
		}
	}
}
