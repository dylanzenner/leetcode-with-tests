package fizz

import "strconv"

func FizzBuzz(n int) []string{
	res := []string{}

	for i := range n {
		i += 1
		if i % 3 == 0 && i % 5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i % 3 == 0{
			res = append(res, "Fizz")
		} else if i % 5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}