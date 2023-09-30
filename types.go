package main

import "net/http"

type Client struct {
	c      *http.Client
	apiKey string
}

type rdUserSchema struct {
	Id         int    `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Points     int    `json:"points,omitempty"`
	Locale     string `json:"locale,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	User_type  string `json:"type,omitempty"`
	Premium    int    `json:"premium,omitempty"`
	Expiration string `json:"expiration,omitempty"`
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
