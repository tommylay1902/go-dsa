package howsum

func HowSumBrute(target int, numbers []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, value := range numbers {
		result := HowSumBrute(target-value, numbers)
		if result != nil {
			return append(result, value)
		}
	}

	return nil
}

func HowSum(target int, numbers []int, memo map[int][]int) []int {
	if value, ok := memo[target]; ok {
		return value
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, value := range numbers {
		resultRemainder := HowSum(target-value, numbers, memo)
		if resultRemainder != nil {
			memo[target] = append(resultRemainder, value)
			return append(resultRemainder, value)
		}
	}
	memo[target] = nil
	return nil
}
