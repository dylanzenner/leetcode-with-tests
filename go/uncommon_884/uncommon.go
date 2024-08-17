package uncommon

import "strings"

func Uncommon(s, t string) []string {
	mapping := make(map[string]bool)
	res := []string{}

	for _, val := range strings.Split(s, " "){
		_, ok := mapping[val]
		if !ok {
			mapping[val] = false
		} else {
			mapping[val] = true
		}
	}

	for _, val := range strings.Split(t, " "){
		_, ok := mapping[val]
		if !ok {
			mapping[val] = false
		} else {
			mapping[val] = true
		}
	}

	for key, val := range mapping {
		if val == false {
			res = append(res, key)
		}
	}

	return res
}