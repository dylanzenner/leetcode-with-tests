package truncate

import "strings"

func Truncate(s string, k int) string {
	splitArray := strings.Split(s, " ")
	return strings.Join(splitArray[:k], " ")
}