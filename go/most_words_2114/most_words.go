package mostwords

import "strings"

func MostWords(words []string) int {
	count := 0
	for _, val := range words {
		length := strings.Split(val, " ")
		if len(length) > count {
			count = len(length)
		}
	}
	return count
}