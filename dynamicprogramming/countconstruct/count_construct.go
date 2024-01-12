package countconstruct

import "strings"

func CountConstructBrute(word string, wordbank []string) int {
	if word == "" {
		return 1
	}
	sum := 0
	for _, value := range wordbank {
		if strings.HasPrefix(word, value) {
			trimWord := strings.TrimPrefix(word, value)
			resultingRemainder := CountConstructBrute(trimWord, wordbank)
			sum += resultingRemainder
		}
	}
	return sum
}

func CountConstruct(word string, wordbank []string) int {
	return helper(word, wordbank, make(map[string]int))
}

func helper(word string, wordbank []string, memo map[string]int) int {
	if value, ok := memo[word]; ok {
		return value
	}
	if word == "" {
		return 1
	}
	sum := 0
	for _, value := range wordbank {
		if strings.HasPrefix(word, value) {
			trimWord := strings.TrimPrefix(word, value)
			resultingRemainder := helper(trimWord, wordbank, memo)
			memo[word] += resultingRemainder
			sum += resultingRemainder
		}
	}

	memo[word] = sum
	return sum
}
