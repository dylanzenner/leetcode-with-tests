package convert

func Convert(c float64) []float64 {
	res := []float64{}

	res = append(res, c + 273.15)
	res = append(res, c * 1.80 + 32.00)
	return res
}