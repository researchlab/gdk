package slice

import (
	"reflect"
)

// UniqueInt64 reports the UniqueInt64 slice
func UniqueInt64(nums []int64) (ret []int64) {
	if len(nums) == 0 {
		return
	}
	tmp := make(map[int64]int64)
	for _, num := range nums {
		if _, ok := tmp[num]; !ok {
			ret = append(ret, num)
			tmp[num] = num
		}
	}
	return ret

	return ret
}

// UniqueInt
func UniqueInt(nums []int) (ret []int) {
	if len(nums) == 0 {
		return
	}
	tmp := make(map[int]int)
	for _, num := range nums {
		if _, ok := tmp[num]; !ok {
			ret = append(ret, num)
			tmp[num] = num
		}
	}
	return ret
}

// UniqueString
func UniqueString(strs []string) (ret []string) {
	if len(strs) == 0 {
		return
	}
	tmp := make(map[string]string)
	for _, str := range strs {
		if _, ok := tmp[str]; !ok {
			ret = append(ret, str)
			tmp[str] = str
		}
	}
	return ret
}

// Unique  reports any type slice to Unique
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
