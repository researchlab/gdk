package generator

import (
	"testing"
)

func TestRandomInts(t *testing.T) {
	tests := []struct {
		Min         int
		Max         int
		Count       int
		expectedLen int
	}{
		{10, 30, 10, 10},
		{1, 30, 5, 5},
		{10, 300, 6, 6},
		{10, 3, 6, 0},
		{10, 13, 6, 0},
	}
	for _, test := range tests {
		nums := RandomInts(test.Min, test.Max, test.Count)
		if len(nums) != test.expectedLen {
			t.Errorf("RandomInts invalid, test:%v, len(nums):%v", test, len(nums))
		}

		for _, num := range nums {
			if num > test.Max || num < test.Min {
				t.Errorf("RandomInts invalid, test:%v, num:%v", test, num)
			}
		}
	}
}
