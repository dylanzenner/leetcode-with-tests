package commonwords

func CommonWords(words1, words2 []string) int {
	mappingOne := make(map[string]int)
	mappingTwo := make(map[string]int)
	count := 0

	for _, val := range words1 {
		_, ok := mappingOne[val]
		if !ok {
			mappingOne[val] = 1
		} else {
			mappingOne[val] += 1
		}
	}

	for _, val := range words2 {
		_, ok := mappingTwo[val]
		if !ok {
			mappingTwo[val] = 1
		} else {
			mappingTwo[val] += 1
		}
	}

	for k, v := range mappingOne {
		if v == 1{
			_, ok := mappingTwo[k]
			if ok && mappingTwo[k] == 1 {
				count += 1
			}
		}
	}
	return count
}