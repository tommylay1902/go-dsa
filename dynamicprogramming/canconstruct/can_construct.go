package canconstruct

import (
	"strings"
)

func CanConstructBrute(word string, wordbank []string) bool {
	if word == "" {
		return true
	}

	for _, value := range wordbank {

		if strings.HasPrefix(word, value) {
			resultString := strings.TrimPrefix(word, value)
			return CanConstructBrute(resultString, wordbank)
		}
	}

	return false
}

func CanConstruct(word string, wordbank []string) bool {
	return helper(word, wordbank, make(map[string]bool))
}

func helper(word string, wordbank []string, memo map[string]bool) bool {

	if value, ok := memo[word]; ok {
		return value
	}
	if word == "" {
		return true
	}

	for _, value := range wordbank {

		if strings.HasPrefix(word, value) {
			resultString := strings.TrimPrefix(word, value)
			memo[resultString] = helper(resultString, wordbank, memo)
			return memo[resultString]
		}
	}
	memo[word] = false
	return false
}
