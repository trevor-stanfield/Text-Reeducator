/*
A small program to help learn some basics of the Go programming language including
in-built collections, file io, the extended library, and string formatting.

Author (full): Trevor Stanfield
Date Published: Tuesday 2-22-2022
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

/*
A function to make a censored string the same length as the word

Input:
	badWord: 		a word of type string that needs to be centered because it's naughty

Output:
	censoredWord: 	a string of the word that is censored by replacing all characters with '*'
*/
func makeCensoredString(badWord string) string {
	censoredWord := ""
	for i := 0; i < len(badWord); i++ {
		censoredWord = fmt.Sprint(censoredWord, "*")
	}
	return censoredWord
}

/*
A function to make a hash map thats keys are forbidden words and whose values
are censor strings the same length as the words

Input:
	scanner: 		an interface of the type *bufio.Scanner with the file filled with filthy words

Output:
	filthyWordMap: 	a hashmap of the type map[string]string thats keys are forbidden words and
					values are censored strings of identical length to their keys
*/
func makeFilthyWordMap(scanner *bufio.Scanner) map[string]string {
	filthyWordMap := make(map[string]string)
	scanner.Split(bufio.ScanWords)
	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		text = strings.ToLower(text)
		filthyWordMap[text] = makeCensoredString(text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return filthyWordMap
}

/*
A function to make a string slice of the alphanumeric sequences in the input phrase

Input:
	wordSlice: 		an empty string slice of type []string to be filled
	phraseToCheck: 	the input phrase to be made into an alphanumeric slice

Output:
	wordSlice: 		the now-filled string slice of the alphanumeric sequences in the input phrase
*/
func makeWordSlice(wordSlice []string, phraseToCheck string) []string {
	//a function to take in a rune and determine if it is alphanumeric, returns bool
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	wordSlice = strings.FieldsFunc(phraseToCheck, f)

	return wordSlice

}

/*
A function that repeatedly takes in a user input until "exit" is input. The input is repeated
back to the user, and if the input includes a forbidden word it is returned with that word censored.

Input:
	filthyWords: 	a hashmap of the type map[string]string thats keys are forbidden words and
					values are censored strings of identical length to their keys
*/
func doAllTheCensorship(filthyWords map[string]string) {
	promptStart := "Please enter"
	promptEnd := "phrase that may require reeducating:"
	phraseToCheck := ""
	more := "a"
	again := false
	for {
		fmt.Println(promptStart, more, promptEnd)
		if !again {
			again = !again
			more = "another"
		}

		in := bufio.NewReader(os.Stdin)

		phrase, err := in.ReadString('\n')
		if err != nil {
			return
		}
		phraseToCheck = phrase

		var wordSlice []string

		wordSlice = makeWordSlice(wordSlice, phraseToCheck)

		// for every word in our wordSLice check if it is forbidden, and
		// if so, censor all instances of it in our input phrase
		for _, word := range wordSlice {
			if word == "exit" {
				return
			}
			if censor, ok := filthyWords[strings.ToLower(word)]; ok {
				phraseToCheck = strings.ReplaceAll(phraseToCheck, word, censor)
			}
		}
		fmt.Println(phraseToCheck)

	}
}

/*CensorySlope
A function to begin rolling the ball down the hill to authoritarianism.
First a file is taken in. Then helper functions are called to do all the
censorship if necessary.
*/
func CensorySlope() {
	fileName := ""
	fmt.Println("Please enter a text file full of filthy, foul utterances:")
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	filthyWords := makeFilthyWordMap(scanner)
	doAllTheCensorship(filthyWords)
}

func main() {
	// write your code here
	CensorySlope()

	fmt.Println("Bye!")

}
