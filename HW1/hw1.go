// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	// Feel free to use the main function for testing your functions
	s := "Hello, 世界"
	fmt.Println(TopCharacters(s, 0))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var phoneRepl string = strings.Replace(phone, " ", "", -1)
	phoneRepl = strings.Replace(phoneRepl, "-", "", -1)
	return fmt.Sprintf("(%s) %s-%s", phoneRepl[0:3], phoneRepl[3:6], phoneRepl[6:])
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		if _, ok := s1Map[r]; !ok {
			s1Map[r] = 1
		} else {
			s1Map[r]++
		}
		i += size
	}
	for i := 0; i < len(s2); {
		r, size := utf8.DecodeRuneInString(s2[i:])
		if _, ok := s2Map[r]; !ok {
			s2Map[r] = 1
		} else {
			s2Map[r]++
		}
		i += size
	}
	return reflect.DeepEqual(s1Map, s2Map)
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var result []int
	for i := range e {
		if e[i]%2 == 0 {
			result = append(result, e[i])
		}
	}
	return result
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	if len(e) == 0 {
		return 0
	}
	prod := 1
	for i := range e {
		prod *= e[i]
	}
	return prod
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)
	var result []int
	for i := range e {
		if _, ok := seen[e[i]]; !ok {
			seen[e[i]] = true
		} else {
			duplicates[e[i]] = true
		}
	}
	for i := range e {
		if _, ok := duplicates[e[i]]; !ok {
			result = append(result, e[i])
		}
	}
	return result
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
func InvertMap(kv map[string]int) map[int]string {
	result := make(map[int]string)
	for k, v := range kv {
		result[v] = k
	}
	return result
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	count := make(map[rune]int)
	result := make(map[rune]int)
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if _, ok := count[rune(r)]; !ok {
			count[rune(r)] = 1
		} else {
			count[rune(r)]++
		}
		i += size
	}
	for key, val := range count {
		if val > k {
			result[key] = val
		}
	}
	return result
}
