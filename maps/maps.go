package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Takes a string splits it at every whitespace then loops the splitted slice (loops all the words)
// and  adds to the word count associated with that key (the word) in a dictionary
// where the word is the key and the wordcount the value
func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	Ssplitted := strings.Split(s, " ")

	for _, v := range Ssplitted {
		wordMap[v]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
