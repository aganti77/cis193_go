package yelp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Review struct {
	Url          string            `json:"url,omitempty"`
	Text         string            `json:"text,omitempty"`
	Rating       float64           `json:"rating,omitempty"`
	User         map[string]string `json:"user,omitempty"`
	Time_created string            `json:"time_created,omitempty"`
}

type ReviewResult struct {
	Reviews []Review `json:"reviews"`
	Total   int      `json:"total"`
}

func (c *Client) GetReviews(id string, locale string) *ReviewResult {
	url := fmt.Sprintf("https://api.yelp.com/v3/businesses/%s/reviews", id)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Unable to reach Yelp API")
	}
	r.Header.Add("Authorization", "Bearer "+c.auth)

	if len(locale) != 0 {
		q := r.URL.Query()
		q.Add("locale", locale)
		r.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(r)
	if err != nil {
		log.Println(err)
	}
	result := new(ReviewResult)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println(err)
	}
	return result
}
