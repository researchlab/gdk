package gdk

// ArrayContains return true if this array contains the specified element.
func ArrayContains[E string | int64 | int | float64](array []E, key E) bool {
	if len(array) == 0 {
		return false
	}
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return true
		}
	}
	return false
}

// ArrayMerge  merge two or more arrays into a new array
func ArrayMerge[E string | int64 | int](arraya, arrayb []E, arrays ...[]E) (array []E) {
	array = append(arraya, arrayb...)
	for _, v := range arrays {
		array = append(array, v...)
	}
	return array
}

// ArraySum  sum of the given array
func ArraySum[E int64 | int | float64](array []E) (sum E) {
	for i := 0; i < len(array); i++ {
		sum += array[i]
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

// ArrayMax return the max one
func ArrayMax[E Ordered](array ...E) E {
	if len(array) < 1 {
		panic("target 'array' cannot be empty.")
	}
	// Finds and returns min
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

// ArrayMin return the smaller one
func ArrayMin[E Ordered](array ...E) E {
	if len(array) < 1 {
		panic("target 'array' cannot be empty.")
	}
	// Finds and returns min
	min := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}

// ArrayMap mapping the given array to a map[r]t
func ArrayMap[E, R, T any](array []E, f TFunc[E, R, T]) (out map[any]T) {
	out = make(map[any]T)
	for _, v := range array {
		r, t := f(v)
		out[r] = t
	}
	return out
}
