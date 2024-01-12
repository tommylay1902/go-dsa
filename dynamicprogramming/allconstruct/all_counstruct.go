package allconstruct

import (
	"strings"
)

func AllConstructBrute(target string, wordbank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range wordbank {
		if strings.HasPrefix(target, word) {
			trimmedWord := strings.TrimPrefix(target, word)
			comboArr := AllConstructBrute(trimmedWord, wordbank)
			for _, combo := range comboArr {
				result = append(result, append([]string{word}, combo...))
			}
		}
	}

	return result
}

func AllConstruct(target string, wordbank []string) [][]string {
	return helper(target, wordbank, make(map[string][][]string))
}

func helper(target string, wordbank []string, memo map[string][][]string) [][]string {

	if value, ok := memo[target]; ok {
		return value
	}

	if target == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range wordbank {
		if strings.HasPrefix(target, word) {
			trimmedWord := strings.TrimPrefix(target, word)
			comboArr := AllConstructBrute(trimmedWord, wordbank)
			for _, combo := range comboArr {
				resultCombo := append(result, append([]string{word}, combo...))
				memo[target] = resultCombo
				result = resultCombo
			}
		}
	}
	memo[target] = result
	return result
}
