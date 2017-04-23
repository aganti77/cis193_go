package yelp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BusinessResult struct {
	Id            string                   `json:"id,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Image_url     string                   `json:"image_url,omitempty"`
	Is_claimed    bool                     `json:"is_claimed,omitempty"`
	Is_closed     bool                     `json:"is_closed,omitempty"`
	Url           string                   `json:"url,omitempty"`
	Phone         string                   `json:"phone,omitempty"`
	Display_phone string                   `json:"display_phone,omitempty"`
	Review_count  int                      `json:"review_count,omitempty"`
	Categories    []map[string]string      `json:"categories,omitempty"`
	Rating        float64                  `json:"rating,omitempty"`
	Location      map[string]interface{}   `json:"location,omitempty"`
	Coordinates   map[string]float64       `json:"coordinates,omitempty"`
	Photos        []string                 `json:"photos,omitempty"`
	Price         string                   `json:"price,omitempty"`
	Hours         []map[string]interface{} `json:"hours,omitempty"`
	Transactiosn  []string                 `json:"transactions,omitempty"`
}

func (c *Client) GetBusinessById(id string) *BusinessResult {
	url := fmt.Sprintf("https://api.yelp.com/v3/businesses/%s", id)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Unable to reach Yelp API")
	}
	r.Header.Add("Authorization", "Bearer "+c.auth)

	resp, err := c.client.Do(r)
	if err != nil {
		log.Println(err)
	}
	result := new(BusinessResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println(err)
	}
	return result
}
