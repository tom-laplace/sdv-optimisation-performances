package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	result := wordsCount("fox fox cat dog, dog again !")
	fmt.Printf("%v", result)
}

func wordsCount(texte string) map[string]int {
	texte = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(texte, "")
	texte = strings.ToLower(texte)
	texte = strings.TrimSpace(texte)
	wordsArray := strings.Split(texte, " ")

	mapResult := make(map[string]int)

	for i := 0; i < len(wordsArray); i++ {
		value, exist := mapResult[wordsArray[i]]

		if exist {
			mapResult[wordsArray[i]] = value + 1
		} else {
			mapResult[wordsArray[i]] = 1
		}
	}

	return mapResult
}
