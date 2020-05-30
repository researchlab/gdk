package slice

import (
	"reflect"
)

func UniqueInt64(s []int64) []int64 {
	size := len(s)
	if size == 0 {
		return []int64{}
	}

	m := make(map[int64]struct{})
	for i := 0; i < size; i++ {
		m[s[i]] = struct{}{}
	}

	realLen := len(m)
	ret := make([]int64, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

func UniqueInt(s []int) []int {
	size := len(s)
	if size == 0 {
		return []int{}
	}

	m := make(map[int]struct{})
	for i := 0; i < size; i++ {
		m[s[i]] = struct{}{}
	}

	realLen := len(m)
	ret := make([]int, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

func UniqueString(s []string) []string {
	size := len(s)
	if size == 0 {
		return []string{}
	}

	m := make(map[string]struct{})
	for i := 0; i < size; i++ {
		m[s[i]] = struct{}{}
	}

	realLen := len(m)
	ret := make([]string, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// unique slice
func Unique(data interface{}) bool {

	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() != reflect.Ptr {
		return false
	}

	tmpData := unique(dataVal.Elem().Interface())
	tmpDataVal := reflect.ValueOf(tmpData)

	dataVal.Elem().Set(tmpDataVal)
	return true
}

func unique(data interface{}) interface{} {
	inArr := reflect.ValueOf(data)
	if inArr.Kind() != reflect.Slice && inArr.Kind() != reflect.Array {
		return data
	}

	existMap := make(map[interface{}]bool)
	outArr := reflect.MakeSlice(inArr.Type(), 0, inArr.Len())

	for i := 0; i < inArr.Len(); i++ {
		iVal := inArr.Index(i)

		if _, ok := existMap[iVal.Interface()]; !ok {
			outArr = reflect.Append(outArr, inArr.Index(i))
			existMap[iVal.Interface()] = true
		}
	}

	return outArr.Interface()
}
