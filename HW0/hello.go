// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"
import "strings"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, à¤¦à¥à¤¨à¤¿à¤¯à¤¾!")
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	i := 3
	for i^2 <= n {
		if n%i == 0 {
			return false
		}
		i += 2
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = strings.ToLower(s)
	var left int = 0
	var right int = len(s) - 1
	for left < right {
		if strings.Compare(string(s[left]), string(s[right])) != 0 {
			return false
		}
		left++
		right--
	}
	return true
}
