package gdk

import (
	"testing"
)

func TestMapClear(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m := map[string]string{"key": "value"}
		MapClear(m)
		if len(m) != 0 {
			t.Errorf("len(m)=%v, want 0", len(m))
		}
	})

	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		MapClear(m)
		if m != nil {
			t.Errorf("m=%+v, want nil", m)
		}
	})
}

func TestMapSize(t *testing.T) {
	m := make(map[string]struct{})
	t.Run("nil map", func(t *testing.T) {
		if got := MapSize(m); got != 0 {
			t.Errorf("got %v, want 0", got)
		}
	})
	t.Run("map[string]struct", func(t *testing.T) {
		m["test"] = struct{}{}
		if got := MapSize(m); got != 1 {
			t.Errorf("got %v, want 1", got)
		}
	})
}

func TestMapRange(t *testing.T) {
	t.Run("value add number 1", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		MapRange(m, func(key string, value int) bool {
			m[key] = value + 1
			return true
		})
		want := map[string]int{"a": 2, "b": 3}
		for k, v := range m {
			if v != want[k] {
				t.Errorf("got %v, want %v", v, want[k])
			}
		}
	})
}

func TestMapFilter(t *testing.T) {
	t.Run("filter < 5", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 6}
		got := MapFilter(m, func(key string, value int) bool {
			if value > 5 {
				return false
			}
			return true
		})
		for k, v := range got {
			if v > 5 {
				t.Errorf("got %v, want %v, key %v", v, "<5", k)
			}
		}
		if len(got) != 2 {
			t.Errorf("got %v, want 2", len(got))
		}
	})
}
func TestMapValues(t *testing.T) {
	t.Run("map values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		want := map[int]bool{1: true, 2: true, 3: true}
		got := MapValues(m)
		for _, v := range got {
			if _, ok := want[v]; !ok {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}

func TestMapKeys(t *testing.T) {
	t.Run("map keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		want := map[string]bool{"a": true, "b": true}
		got := MapKeys(m)
		for _, v := range got {
			if _, ok := want[v]; !ok {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})

}
