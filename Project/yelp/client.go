package yelp

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	auth   string
}

func GetAuthClient(client_id string, client_secret string) *Client {
	authUrl := "https://api.yelp.com/oauth2/token"
	data := url.Values{}
	data.Add("client_id", client_id)
	data.Add("client_secret", client_secret)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", authUrl, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	token := string(body)
	return &Client{
		client: client,
		auth:   token[18 : len(token)-50],
	}
}
