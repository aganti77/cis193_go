package yelp

import (
	"encoding/json"
	"log"
	"net/http"
)

type Business struct {
	Id            string                 `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Image_url     string                 `json:"image_url,omitempty"`
	Is_closed     bool                   `json:"is_closed,omitempty"`
	Url           string                 `json:"url,omitempty"`
	Review_count  int                    `json:"review_count,omitempty"`
	Categories    []map[string]string    `json:"categories,omitempty"`
	Rating        float64                `json:"rating,omitempty"`
	Coordinates   map[string]float64     `json:"coordinates,omitempty"`
	Transactions  []string               `json:"transactions,omitempty"`
	Price         string                 `json:"price,omitempty"`
	Location      map[string]interface{} `json:"location,omitempty"`
	Phone         string                 `json:"phone,omitempty"`
	Display_phone string                 `json:"display_phone,omitempty"`
	Distance      float64                `json:"distance,omitempty"`
}

type SearchResult struct {
	Businesses []Business                    `json:"businesses"`
	Total      int                           `json:"total"`
	Region     map[string]map[string]float64 `json:"region,omitempty"`
}

func (c *Client) Search(params map[string]string) *SearchResult {
	r, err := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search", nil)
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
	result := new(SearchResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func (c *Client) PhoneSearch(phone string) *SearchResult {
	r, err := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search/phone", nil)
	if err != nil {
		log.Println("Unable to reach Yelp API")
	}
	r.Header.Add("Authorization", "Bearer "+c.auth)

	q := r.URL.Query()
	q.Add("phone", phone)
	r.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(r)
	if err != nil {
		log.Println(err)
	}
	result := new(SearchResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func (c *Client) WhoDelivers(params map[string]string) *SearchResult {
	r, err := http.NewRequest("GET", "https://api.yelp.com/v3/transactions/delivery/search", nil)
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
	result := new(SearchResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println(err)
	}
	return result
}
