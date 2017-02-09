// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var c Cart
	c.AddItem("eggs")
	c.AddItem("bread")
	c.Checkout()
	fmt.Println(c)
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	strPrice := strconv.FormatInt(int64(p), 10)
	if len(strPrice) == 1 {
		strPrice = "00" + strPrice
	} else if len(strPrice) == 2 {
		strPrice = "0" + strPrice
	}
	return fmt.Sprintf("$%s.%s", strPrice[:len(strPrice)-2], strPrice[len(strPrice)-2:len(strPrice)])
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if _, ok := Prices[item]; ok {
		log.Println("Item already in map, but price now overwritten")
	}
	Prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	i := 0
	for i < len(c.Items) {
		if c.Items[i] == "milk" {
			return true
		}
		i++
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	i := 0
	for i < len(c.Items) {
		if c.Items[i] == item {
			return true
		}
		i++
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	if val, ok := Prices[item]; ok {
		c.Items = append(c.Items, item)
		c.TotalPrice += Price(val)
	} else {
		log.Fatalln("Item cannot be found")
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Printf("Final cart balance: %s\n", c.TotalPrice)
	c.Items = []string{}
	c.TotalPrice = 0
}
