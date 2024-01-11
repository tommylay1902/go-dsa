package cansum

func CanSum(target int, numbers []int) bool {
	return helper(target, numbers, make(map[int]bool))

}

func BruteCanSum(target int, numbers []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, value := range numbers {
		if BruteCanSum(target-value, numbers) {
			return true
		}
	}

	return false
}

func helper(target int, numbers []int, memo map[int]bool) bool {
	if value, ok := memo[target]; ok {
		return value
	}
	if target == 0 {
		memo[target] = true
		return true
	}
	if target < 0 {
		memo[target] = false
		return false
	}

	for _, value := range numbers {
		if helper(target-value, numbers, memo) {
			memo[target] = true
			return true
		}
	}

	memo[target] = false
	return false
}
