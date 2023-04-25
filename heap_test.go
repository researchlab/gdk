package gdk

import (
	"fmt"
	"testing"
)

type player struct {
	level int
	name  string
}

func TestNewHeap(t *testing.T) {
	h := NewHeap([]player{}, func(p1, p2 player) int {
		return p1.level - p2.level
	})
	t.Run("TestNewHeap EmptyCheck", func(t *testing.T) {
		//want := player{}
		var want player
		if got := h.Pop(); got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	for i := 100; i > 0; i-- {
		h.Push(player{i, fmt.Sprintf("name%d", i)})
	}
	t.Run("TestNewHeap Sort", func(t *testing.T) {
		if got := h.Pop(); got.level != 1 {
			t.Errorf("got %v, want 1", got.level)
		}
	})
}

func TestCopy(t *testing.T) {
	h := NewHeap([]*player{}, func(p1, p2 *player) int {
		return p1.level - p2.level
	})
	t.Run("nil check", func(t *testing.T) {
		if got := h.Pop(); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
	h.Push(&player{1, "math"})
	h.Push(&player{2, "mat"})

	h2 := h.Copy()
	h2.Push(&player{3, "ma"})
	t.Run("Len Check", func(t *testing.T) {
		if h.Len() != 2 {
			t.Errorf("got %v, want 2", h.Len())
		}
		if h2.Len() != 3 {
			t.Errorf("got %v, want 3", h2.Len())
		}
	})
	t.Run("Pop check", func(t *testing.T) {
		data := h.Pop()
		if data.level != 1 {
			t.Errorf("got %v, want 1", data.level)
		}
		data = h2.Pop()
		if data.level != 1 {
			t.Errorf("got %v, want 1", data.level)
		}
		data = h.Pop()
		if data.level != 2 {
			t.Errorf("got %v, want 2", data.level)
		}
		data = h.Pop()
		if data != nil {
			t.Errorf("got %v, want nil", data)
		}
	})
}
