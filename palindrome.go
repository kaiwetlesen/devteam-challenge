package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	longestSentence := ""
	longestSentenceLen := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if longestSentenceLen < len(text) {
			longestSentence = text
			longestSentenceLen = len(text)
		}
		//fmt.Printf(strings.ToLower(text) + "\n")
		IsPalindrome(text)
	}
	fmt.Printf("Longest sentence: %v\n", longestSentence)
}

func IsPalindrome(sentence string) bool {
	newSentence := strings.ToLower(sentence)
	reg := regexp.MustCompile("[^a-z]+")
	fmt.Printf(reg.ReplaceAllString(newSentence, "") + "\n")
	//reverse := Reverse(sentence)

	return false
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
