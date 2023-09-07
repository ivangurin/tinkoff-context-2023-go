package main

import (
	"fmt"
	"strings"
)

func main() {

	var phrase string
	fmt.Scanln(&phrase)

	tempalte := "sheriff"

	letters := make(map[string]int)

	for _, letter := range tempalte {
		letters[string(letter)]++
	}

	matchCounter := 0

	for letter, occurse := range letters {

		counter := strings.Count(phrase, string(letter)) / occurse
		if counter == 0 {
			break
		}

		if matchCounter == 0 || counter < matchCounter {
			matchCounter = counter
		}

	}

	fmt.Println(matchCounter)

}
