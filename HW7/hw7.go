// Homework 7: Web Scraping
// Due March 28, 2017 at 11:59pm
package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	i := ScrapeHackerNews(5)
	for j := range i {
		fmt.Println(i[j])
	}
}

// News is a Hacker News article listing
type News struct {
	Points   int
	Title    string
	Username string
	URL      string
}

// NewsSlice is a slice of News pointers
type NewsSlice []*News

// ScrapeHackerNews scrapes the website "https://news.ycombinator.com/" using goquery and returns
// information on the first n posts.
//
// For each post, the attributes to be extracted are: points, title, username and url.
// This data should be returned as a NewsSlice, where NewsSlice is a custom slice of News structs.
//
// For example, for the sample image located at `https://www.cis.upenn.edu/~cis193/homeworks/hn.png`,
// the struct would look like:
// News{24, "QEMU(TCG): user-to-root privesc inside VM via bad translation caching",
// "webaholic", "https://bugs.chromium.org/p/project-zero/issues/detail?id=1122"}.
//
// If n is greater than the number of total posts available (which should be 30), return data from
// the all of the available posts (all thirty).
func ScrapeHackerNews(n int) NewsSlice {
	doc, err := goquery.NewDocument("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}
	result := make([]*News, 0)
	points := make([]int, 0)
	titles := make([]string, 0)
	usernames := make([]string, 0)
	URLs := make([]string, 0)
	titleUrls := doc.Find("tr.athing td.title a.storylink")
	for i := range titleUrls.Nodes {
		title := titleUrls.Eq(i).Text()
		titles = append(titles, title)
		url, _ := titleUrls.Eq(i).Attr("href")
		URLs = append(URLs, url)
	}
	pointsName := doc.Find("td.subtext")
	for i := range pointsName.Nodes {
		point := pointsName.Eq(i).Find("span.score").Text()
		if len(point) == 0 {
			point = "0 points"
		}
		point = point[0 : len(point)-7]
		pt, _ := strconv.Atoi(point)
		points = append(points, pt)
		username := pointsName.Eq(i).Find("a.hnuser").Text()
		if len(username) == 0 {
			username = ""
		}
		usernames = append(usernames, username)
	}
	if n > 30 {
		n = 30
	}
	for i := 0; i < n; i++ {
		adder := &News{Points: points[i], Title: titles[i], Username: usernames[i], URL: URLs[i]}
		result = append(result, adder)
	}
	return NewsSlice(result)
}

// GetEmails returns a string slice of the emails found on the given URL.
//
// Scenario: you are a student enthusiastic about spreading awareness about Go. To effectively
// market Go, you decide to email Penn CIS professors about the wonders of the Go programming
// language. In this function, use goquery to extract the email addresses from the URL
// "http://www.cis.upenn.edu/about-people/" and return them as a string slice. This will involve you
// having to investigate where and how emails are located on the webpage.
// Note: you should have 47 total emails returned.
func GetEmails() []string {
	doc, err := goquery.NewDocument("http://www.cis.upenn.edu/about-people/")
	if err != nil {
		log.Fatal(err)
	}
	result := make([]string, 0)
	nodes := doc.Find("table tbody tr td a")
	for i := range nodes.Nodes {
		href, _ := nodes.Eq(i).Attr("href")
		if len(href) > 1 && href[0:7] == "mailto:" {
			result = append(result, href[7:len(href)])
		}
	}
	return result
}

// CountryData has GDP information on a country
type CountryData struct {
	Country string
	GDP     string
}

// GetCountryGDP takes in a string country name and returns the GDP (in millions) as
// an integer. Information on the country is found by concurrently scraping a hidden website with
// data on countries scattered on many pages.
//
// Scenario: imagine you are a spy and you have discovered a URL with top secret GDP information:
// "https://www.cis.upenn.edu/~cis193/scraping/9828772efc2bd314a277c8880695dea2.html". This webpage
// has a country name and the GDP (in millions of US Dollars). It also has links to two other
// country's webpages. Based on intelligence you've received, every country has a webpage on this
// website with information about it, but you do not know the URL for each page. You can assume that
// none of the page links lead you to a cycle and every country can be reached from a path from the
// initial URL that you are given. So, for this function, you will need to traverse from the initial
// url to every webpage link you encounter in order to find information on the target `country`
// string. Since time is of the essence, you want to use concurrency to scrape webpages
// simultaneously. Note that for this function, we only care about getting the GDP for the input
// `country` string. You may find it useful to use the CountryData struct to send country
// information between goroutines.
//
// To prevent the function from getting stuck if an invalid `country` string is entered,
// you should also implement a timeout that will automatically return an error after 10 seconds
// if the program hasn't already finished terminating.
//
// Feel free to make and use helper functions for this function. To help with testing this
// function, we know from intelligence reports that the GDP for "Canada" is 1532343 and
// the GDP for "Colombia" is 274135.
func helper(search string, url string, c chan CountryData) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	country := doc.Find("div h3.country").Eq(0).Text()
	gdp := doc.Find("div h3.gdp").Eq(0).Text()
	if country == search {
		c <- CountryData{Country: country, GDP: gdp}
	} else {
		links := doc.Find("ul li a")
		for i := range links.Nodes {
			link, _ := links.Eq(i).Attr("href")
			go func() {
				helper(search, link, c)
			}()
		}
	}
}

func GetCountryGDP(country string) (int, error) {
	data := make(chan CountryData)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(10 * time.Second)
		timeout <- true
	}()
	go func() {
		helper(country, "https://www.cis.upenn.edu/~cis193/scraping/9828772efc2bd314a277c8880695dea2.html", data)
	}()
	select {
	case d := <-data:
		gdp, _ := strconv.Atoi(strings.Replace(d.GDP, ",", "", -1))
		return gdp, nil
	case <-timeout:
		return 0, errors.New("Timeout occurred")
	}
}
