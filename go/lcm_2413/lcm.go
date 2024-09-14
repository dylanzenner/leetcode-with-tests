package lcm

func Lcm(n int ) int {
	if n % 2 == 0 {
		return n
	} else {
		return n * 2
	}
}