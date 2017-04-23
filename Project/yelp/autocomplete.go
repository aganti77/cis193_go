package yelp

import (
	"encoding/json"
	"log"
	"net/http"
)

type AutocompleteResult struct {
	Categories []map[string]string `json:"categories,omitempty"`
	Businesses []map[string]string `json:"businesses,omitempty"`
	Terms      []map[string]string `json:"terms,omitempty"`
}

func (c *Client) Autocomplete(params map[string]string) *AutocompleteResult {
	r, err := http.NewRequest("GET", "https://api.yelp.com/v3/autocomplete", nil)
	if err != nil {
		log.Println("Unable to reach Yelp API")
	}
	r.Header.Add("Authorization", "Bearer "+c.auth)

	q := r.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(r)
	if err != nil {
		log.Println(err)
	}
	result := new(AutocompleteResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println(err)
	}
	return result
}
