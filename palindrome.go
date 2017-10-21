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

	numPalindromes := 0
	longestPalindrome := ""

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
		if IsPalindrome(text) {
			//fmt.Println(text + " is a palindrome!\n")
			numPalindromes++
			if len(text) > len(longestPalindrome) {
				longestPalindrome = text
			}
		}
	}
	fmt.Printf("Longest sentence: \"%v\"\n\n", longestSentence)
	fmt.Printf("Number of palindromes: %v\n\n", numPalindromes)
	fmt.Printf("Longest palindrome: \"%v\"\n", longestPalindrome)

}

func IsPalindrome(sentence string) bool {
	cleanSentence := strings.ToLower(sentence)
	reg := regexp.MustCompile("[^a-z]+")
	cleanestSentence := reg.ReplaceAllString(cleanSentence, "")
	reverseSentence := Reverse(cleanestSentence)

	return reverseSentence == cleanestSentence
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
