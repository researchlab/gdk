package slice

// SumInt64
func SumInt64(s []int64) (sum int64) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumInt
func SumInt(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumFloat64
func SumFloat64(s []float64) (sum float64) {
	for _, v := range s {
		sum += v
	}
	return
}
