package gdk

// ArrayContains return true if this array contains the specified element.
func ArrayContains[E string | int64 | int | float64](data []E, key E) bool {
	if len(data) == 0 {
		return false
	}
	for i := 0; i < len(data); i++ {
		if data[i] == key {
			return true
		}
	}
	return false
}

// ArrayMerge  merge two or more arrays into a new array
func ArrayMerge[E string | int64 | int](a, b []E, c ...[]E) (data []E) {
	data = append(a, b...)
	for _, v := range c {
		data = append(data, v...)
	}
	return data
}

// ArraySum  sum of the given array
func ArraySum[E int64 | int | float64](data []E) (sum E) {
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

// ArrayUnique remove Duplicate item of the given array
func ArrayUnique[E int64 | int | string](in []E) (out []E) {
	m := make(map[E]struct{})
	out = []E{}
	for i := 0; i < len(in); i++ {
		if _, ok := m[in[i]]; !ok {
			out = append(out, in[i])
			m[in[i]] = struct{}{}
		}
	}
	return out
}
