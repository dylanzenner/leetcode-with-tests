package count

func Count(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	pointer := 0

	for pointer < len(items) {
		if ruleKey == "type" && ruleValue == items[pointer][0] {
			count += 1
		} else if ruleKey == "color" && ruleValue == items[pointer][1] {
			count += 1
		} else if ruleKey == "name" && ruleValue == items[pointer][2] {
			count += 1
		}
		pointer += 1
	}

	return count
}