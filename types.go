package main

import (
	"net/http"
	"time"
)

type Client struct {
	c      *http.Client
	apiKey string
}

type rdUserSchema struct {
	Id         int       `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	Points     int       `json:"points,omitempty"`
	Locale     string    `json:"locale,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	User_type  string    `json:"type,omitempty"`
	Premium    int       `json:"premium,omitempty"`
	Expiration time.Time `json:"expiration,omitempty"`
}

type rdTorrentSchema struct {
	Id       string   `json:"id,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Hash     string   `json:"hash,omitempty"`
	Bytes    int      `json:"bytes,omitempty"`
	Host     string   `json:"host,omitempty"`
	Split    int      `json:"split,omitempty"`
	Progress int      `json:"progress,omitempty"`
	Status   string   `json:"status,omitempty"`
	Added    string   `json:"added,omitempty"`
	Links    []string `json:"link,omitempty"`
	Ended    string   `json:"ended,omitempty"`
	Speed    int      `json:"speed,omitempty"`
	Seeders  int      `json:"seeders,omitempty"`
}

type rdAddMagnetSchema struct {
	Id  string `json:"id,omitempty"`
	Uri string `json:"uri,omitempty"`
}

type rdTorrentInfoSchema struct {
	Id               string         `json:"id,omitempty"`
	Filename         string         `json:"filename,omitempty"`
	OriginalFilename string         `json:"original_filename,omitempty"`
	Hash             string         `json:"hash,omitempty"`
	Bytes            int            `json:"bytes,omitempty"`
	OriginalBytes    int            `json:"original_bytes,omitempty"`
	Host             string         `json:"host,omitempty"`
	Split            int            `json:"split,omitempty"`
	Progress         int            `json:"progress,omitempty"`
	Status           string         `json:"status,omitempty"`
	Added            string         `json:"added,omitempty"`
	Files            []TorrentFiles `json:"files,omitempty"`
	Links            []string       `json:"links,omitempty"`
	Ended            string         `json:"ended,omitempty"`
	Speed            int            `json:"speed,omitempty"`
	Seeders          int            `json:"seeders,omitempty"`
}

type TorrentFiles struct {
	Id       int    `json:"id,omitempty"`
	Path     string `json:"path,omitempty"`
	Bytes    int    `json:"bytes,omitempty"`
	Selected int    `json:"selected,omitempty"`
}

type UnrestrictLink struct {
  Id         string `json:"id,omitempty"`
  Filename   string `json:"filename,omitempty"`
  MimeType   string `json:"mimeType,omitempty"`
  Filesize   int    `json:"filesize,omitempty"`
  Link       string `json:"link,omitempty"`
  Host       string `json:"host,omitempty"`
  Chunks     int    `json:"chunks,omitempty"`
  Crc        int    `json:"crc,omitempty"`
  Download   string `json:"download,omitempty"`
  Streamable int    `json:"streamable,omitempty"`
}
