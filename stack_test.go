package gdk

import "testing"

type Info struct {
	Name string
}

func TestNewStack(t *testing.T) {
	t.Run("StackEmpty", func(t *testing.T) {
		s := NewStack[*Info]()
		if s.IsEmpty() != true {
			t.Errorf("got %v, want true", s.IsEmpty())
		}
		ss := NewStackSize[*Info](-1)
		if ss.IsEmpty() != true {
			t.Errorf("got %v, want true", ss.IsEmpty())
		}
	})
	t.Run("StackSize", func(t *testing.T) {
		s := NewStackSize[*Info](16)
		if s.Cap() != 16 {
			t.Errorf("got %v, want %v", s.Cap(), 16)
		}
	})
}

func TestStackPushPop(t *testing.T) {
	s := NewStackSize[*Info](16)
	s.Push(&Info{Name: "matt"})
	if s.Cap() != 15 {
		t.Errorf("s.Cap() should return 15, got %v", s.Cap())
	}
	if s.IsEmpty() != false {
		t.Errorf("s.IsEmpty() should return false, got %v", s.IsEmpty())
	}
	info := s.Pop()
	if info == nil {
		t.Errorf("info should not be nil, got %v", info)
	}
	if info.Name != "matt" {
		t.Errorf("info.Name should be matt, got %v", info.Name)
	}
	info = s.Pop()
	if info != nil {
		t.Errorf("info should be nil, got %v", info)
	}
	if s.IsEmpty() != true {
		t.Errorf("s.IsEmpty() should be true, got %v", s.IsEmpty())
	}
}

func TestStackResize(t *testing.T) {
	t.Run("Resize", func(t *testing.T) {
		s := NewStackSize[*Info](2)
		if s.Size() != 0 {
			t.Errorf("s.Size() should equal 0, got %v", s.Size())
		}
		if s.Cap() != 2 {
			t.Errorf("s.Cap() should equal 2, got %v", s.Cap())
		}
		s.Push(&Info{"matt"})
		s.Push(&Info{"xml"})
		if s.Size() != 2 {
			t.Errorf("s.Size() should return 2, got %v", s.Size())
		}
		if c := s.Cap(); c != 0 {
			t.Errorf("s.Cap() should return 0, got %v", c)
		}

		s.Push(&Info{"max"})
		if s := s.Size(); s != 3 {
			t.Errorf("s.Size() should return 3, got %v", s)
		}
		if c := s.Cap(); c != 1 {
			t.Errorf("s.Cap() should return 1, got %v", c)
		}
	})
}

func TestStackCopy(t *testing.T) {
	t.Run("Copy", func(t *testing.T) {
		s := NewStackSize[*Info](2)
		s.Push(&Info{"matt"})
		s.Push(&Info{"xl"})
		cs := s.Copy()
		if cs.Cap() != s.Cap() {
			t.Errorf("got %v, want %v", cs.Cap(), s.Cap())
		}
		if cs.Size() != s.Size() {
			t.Errorf("got %v, want %v", cs.Size(), s.Size())
		}
	})
}
