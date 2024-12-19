package shortener

import (
	"math/rand/v2"
	"strings"
)

const URL_LEN = 6

var characters = getCharacters()

func getRandomIntInRange(max, min int) int {
	return rand.IntN(max-min+1) + min
}

func getCharacters() []string {
	letters := []string{}

	for i := 'A'; i <= 'Z'; i++ {
		letters = append(letters, string(i), strings.ToLower(string(i)))
	}

	return letters
}

func GenerateString() string {
	min := 0
	max := len(characters) - 1

	url := ""

	for range URL_LEN {
		url += characters[getRandomIntInRange(max, min)]
	}

	return url
}
