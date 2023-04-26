package gdk

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayContains(t *testing.T) {
	R := func(t *testing.T, got, want bool) {
		if got != want {
			t.Errorf("ArrayContains()=%v, want %v", got, want)
		}
	}
	t.Run("String ArrayContains true", func(t *testing.T) {
		R(t,
			ArrayContains([]string{"mike", "jack"}, "mike"), // got
			true, // want
		)
	})
	t.Run("String ArrayContains false", func(t *testing.T) { R(t, ArrayContains([]string{"mike", "jack"}, "hello"), false) })
	t.Run("int ArrayContains true", func(t *testing.T) {
		R(t,
			ArrayContains([]int{1, 2}, 1), // got
			true,                          // want
		)
	})
	t.Run("int ArrayContains false", func(t *testing.T) {
		R(t,
			ArrayContains([]int{1, 2}, 3), // got
			false,                         // want
		)
	})
	t.Run("float64 ArrayContains true", func(t *testing.T) {
		R(t,
			ArrayContains([]float64{1.1, 2.0}, 1.1), // got
			true,                                    // want
		)
	})
	t.Run("int ArrayContains false", func(t *testing.T) {
		R(t,
			ArrayContains[float64]([]float64{1, 2.2}, 3.0), // got
			false, // want
		)
	})

}

func TestArrayMerge(t *testing.T) {
	t.Run("nil merge array", func(t *testing.T) {
		if got := ArrayMerge(nil, []int{1, 2}); !reflect.DeepEqual(got, []int{1, 2}) {
			t.Errorf("ArrayMerge(nil, int{1,2})=%+v, want %v", got, "[]int{1, 2}")
		}
	})
	t.Run("array merge nil", func(t *testing.T) {
		if got := ArrayMerge([]int{1, 2}, nil); !reflect.DeepEqual(got, []int{1, 2}) {
			t.Errorf("ArrayMerge(int{1,2},nil)=%+v, want %v", got, "[]int{1, 2}")
		}
	})
	t.Run("large merge small", func(t *testing.T) {
		if got := ArrayMerge([]int{1, 2, 3}, []int{1, 5}); !reflect.DeepEqual(got, []int{1, 2, 3, 1, 5}) {
			t.Errorf("got %v, want %v", got, "[]int{1,2,3,1,5}")
		}
	})
	t.Run("small merge large", func(t *testing.T) {
		if got := ArrayMerge([]int{1, 2}, []int{1, 5, 6, 7}); !reflect.DeepEqual(got, []int{1, 2, 1, 5, 6, 7}) {
			t.Errorf("got %v, want %v", got, "[]int{1,2,1,5,6,7}")
		}
	})
	t.Run("more than 2 array merged", func(t *testing.T) {
		if got := ArrayMerge([]int{1, 2}, []int{1}, []int{1, 5, 7}, []int{9, 0}, []int{1, 1}); !reflect.DeepEqual(got, []int{1, 2, 1, 1, 5, 7, 9, 0, 1, 1}) {
			t.Errorf("got %v, want %v", got, "[]int{1,2,1,1,5,7,9,0,1,1}")
		}
	})
	t.Run("merge effect", func(t *testing.T) {
		a := []string{"a", "b", "cc"}
		b := []string{"d", "e"}
		c := []string{}
		if got := ArrayMerge(a, b, c, nil); !reflect.DeepEqual(got, []string{"a", "b", "cc", "d", "e"}) {
			t.Errorf("got %v, want %v", got, "[]string{a,b,cc,d,e}")
		} else {
			got[0] = "111"
			got[3] = "222"
			if !reflect.DeepEqual(a, []string{"a", "b", "cc"}) {
				t.Errorf("got %v, want %v", a, "[]string{a,b,cc}")
			}
			if !reflect.DeepEqual(b, []string{"d", "e"}) {
				t.Errorf("got %v, want %v", b, "[]string{d,e}")
			}
			if !reflect.DeepEqual(got, []string{"111", "b", "cc", "222", "e"}) {
				t.Errorf("got %v, want %v", got, "[]string{111,b,cc,222,e}")
			}
		}

	})
}

func TestArraySum(t *testing.T) {
	t.Run("sum of array int", func(t *testing.T) {
		if got := ArraySum([]int{1, 2, 3}); got != 6 {
			t.Errorf("got %v want 6", got)
		}
	})

	t.Run("sum of array int64", func(t *testing.T) {
		if got := ArraySum([]int64{1, 2, 3}); got != 6 {
			t.Errorf("got %v want 6", got)
		}
	})
	t.Run("sum of array float64", func(t *testing.T) {
		if got := ArraySum([]float64{1.0, 2.1, 3.5}); got != 6.6 {
			t.Errorf("got %v want 6.6", got)
		}
	})

	t.Run("sum of array nil", func(t *testing.T) {
		if got := ArraySum[int](nil); got != 0 {
			t.Errorf("got %v want 0", got)
		}
	})
	t.Run("sum of empty array", func(t *testing.T) {
		if got := ArraySum([]int{}); got != 0 {
			t.Errorf("got %v want 0", got)
		}
	})
}

func TestArrayUnique(t *testing.T) {
	t.Run("remove duplicate item of array int64", func(t *testing.T) {
		if got := ArrayUnique([]int64{1, 1, 1, 2, 2}); !reflect.DeepEqual(got, []int64{1, 2}) {
			t.Errorf("got %v, want []int64[1,2]", got)
		}
	})

	t.Run("remove duplicate item of array nil", func(t *testing.T) {
		if got := ArrayUnique([]int64{}); !reflect.DeepEqual(got, []int64{}) {
			t.Errorf("got %v, want []int64{}", got)
		}
	})

	t.Run("remove duplicate item of array string", func(t *testing.T) {
		if got := ArrayUnique([]string{"a", "a", "aa", "ab"}); !reflect.DeepEqual(got, []string{"a", "aa", "ab"}) {
			t.Errorf("got %v, want []string{a,aa,ab}", got)
		}
	})

	t.Run("remove duplicate item of array int", func(t *testing.T) {
		if got := ArrayUnique([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
			t.Errorf("got %v, want []int{1,2,3}", got)
		}
	})
}

func TestArrayMax(t *testing.T) {
	t.Run("max float64", func(t *testing.T) {
		if got := ArrayMax(1.1, 2.1, 2.2, 2.3); got != 2.3 {
			t.Errorf("got %v, want 2.3", got)
		}
	})
	t.Run("max array", func(t *testing.T) {
		if got := ArrayMax([]int{1, 2, 3, 0}...); got != 3 {
			t.Errorf("got %v, want 3", got)
		}
	})
}

func TestArrayMin(t *testing.T) {
	t.Run("min float64", func(t *testing.T) {
		if got := ArrayMin(1.1, 2.1, 2.2, 2.3); got != 1.1 {
			t.Errorf("got %v, want 1.1", got)
		}
	})
	t.Run("min array", func(t *testing.T) {
		if got := ArrayMin([]int{1, 2, 3, 0}...); got != 0 {
			t.Errorf("got %v, want 0", got)
		}
	})
}

type stu struct {
	name  string
	alias string
}

func (s *stu) Name() string {
	return s.name
}
func TestArrayToMap(t *testing.T) {
	t.Run("mapping basic array to a map", func(t *testing.T) {
		a := []int{1, 2, 3}
		got := ArrayToMap(a, func(i int) (int, struct{}) {
			return i, struct{}{}
		})
		for _, v := range a {
			if vv, ok := got[v]; !ok {
				t.Errorf("got false,  want true, v %v", v)
			} else if vv != struct{}{} {
				t.Errorf("got %v, want struct{}{}", vv)
			}
		}
	})

	stus := []stu{
		stu{"mike", "mike"},
		stu{"jack", "jack"},
	}
	t.Run("mapping struct array to a map", func(t *testing.T) {
		got := ArrayToMap(stus, func(s stu) (r string, t stu) {
			return s.name, s
		})
		for _, s := range stus {
			v, ok := got[s.name]
			if !ok {
				t.Errorf("got %v, want %v", v, s)
			}
			if v.name != v.alias {
				t.Errorf("map value invalid, got %v", v)
			}
		}
	})
	t.Run("mapping struct array to a map", func(t *testing.T) {
		got := ArrayToMap(stus, func(s stu) (r string, t *stu) {
			return s.name, &s
		})
		for _, s := range stus {
			v, ok := got[s.name]
			if !ok {
				t.Errorf("got %v, want %v", *v, s)
			}
			// 注意是(*v).name 不是*v.name, 也不是v.name
			t.Log(fmt.Printf("%T, %v, %v, %s", v, v, *v, (*v).name))
			t.Log("name=", v.Name(), " name=", (*v).Name())

			//if *v.name != *v.alias {
			//	t.Errorf("map value invalid, got %v", *v)
			//}
			if reflect.TypeOf(v).Kind() != reflect.Pointer {
				t.Errorf("got %v, want pointer", reflect.TypeOf(v))
			}
		}
	})

}
