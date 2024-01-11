package bestsum

func BestSumBrute(target int, numbers []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var result []int = nil
	for _, value := range numbers {
		resultRemainder := BestSumBrute(target-value, numbers)

		if resultRemainder != nil {
			currentSum := append(resultRemainder, value)
			if result == nil || len(currentSum) < len(result) {
				result = currentSum
			}
		}
	}

	return result
}

func BestSum(target int, numbers []int) []int {
	return helper(target, numbers, make(map[int][]int))
}

func helper(target int, numbers []int, memo map[int][]int) []int {
	if value, ok := memo[target]; ok {
		return value
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var result []int = nil
	for _, value := range numbers {
		resultRemainder := helper(target-value, numbers, memo)

		if resultRemainder != nil {
			currentSum := append(resultRemainder, value)
			if result == nil || len(currentSum) < len(result) {
				result = currentSum
			}
		}
	}

	memo[target] = result
	return result
}
