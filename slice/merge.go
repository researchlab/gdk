package slice

// SliceMerge merges interface slices to one slice.
func Merge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

// MergeInt  merges the int slice to one
func MergeInt(slice1, slice2 []int) (c []int) {
	c = append(slice1, slice2...)
	return
}

// MergeInt64  merges the int64 slice to one
func MergeInt64(slice1, slice2 []int64) (c []int64) {
	c = append(slice1, slice2...)
	return
}

// MergeString  merges the string slice to one
func MergeString(slice1, slice2 []string) (c []string) {
	c = append(slice1, slice2...)
	return
}
