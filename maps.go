// Exercise: Maps
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	wordList := strings.Fields(s)
	for _, word := range wordList {
		_, isExists := result[word]
		if isExists {
			result[word] += 1
		} else {
			result[word] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
