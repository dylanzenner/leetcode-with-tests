package luckyInt

import "slices"

func LuckyInt(arr []int) int {
	slices.Sort(arr)
	mapping := make(map[int]int)

	for _, v := range arr {
		_, ok := mapping[v]
		if !ok {
			mapping[v] = 1
		} else {
			mapping[v] += 1
		}
	}

	pointer := len(arr) - 1

	for pointer > -1 {
		if mapping[arr[pointer]] == arr[pointer] {
			return arr[pointer]
		}
		pointer -= 1
	}

	return -1
}