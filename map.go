package gdk

// MapClear remove all keys and values in map
func MapClear[K comparable, V any](data map[K]V) {
	for k := range data {
		delete(data, k)
	}
}

// MapSize return count of size
func MapSize[K comparable, V any](data map[K]V) int {
	if data == nil {
		return 0
	}
	return len(data)
}

// MapRange calls f sequentially for each key and value present in the map.
func MapRange[K comparable, V any](data map[K]V, f BiFunc[bool, K, V]) {
	for k, v := range data {
		ok := f(k, v)
		if !ok {
			break
		}
	}
}

// MapFilter 过滤出符合条件的key,value
func MapFilter[K comparable, V any](data map[K]V, f BiFunc[bool, K, V]) map[K]V {
	ret := make(map[K]V)
	for k, v := range data {
		ok := f(k, v)
		if !ok {
			continue
		}
		ret[k] = v
	}
	return ret
}

// MapValues return all value as slice in map
func MapValues[K comparable, V any](data map[K]V) []V {
	ret := make([]V, 0)
	MapRange(data, func(key K, value V) bool {
		ret = append(ret, value)
		return true
	})
	return ret
}

// MapKeys return all key as slice in map
func MapKeys[K comparable, V any](data map[K]V) []K {
	ret := make([]K, 0)
	MapRange(data, func(key K, value V) bool {
		ret = append(ret, key)
		return true
	})
	return ret
}
