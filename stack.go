package gdk

import "sync"

const (
	defaultInitSize = 16
)

// Stack
type Stack[E any] struct {
	data  []E
	pos   int
	lock  sync.Mutex
	empty E
}

func NewStack[E any]() *Stack[E] {
	return NewStackSize[E](defaultInitSize)
}

func NewStackSize[E any](size int) *Stack[E] {
	if size <= 0 {
		size = defaultInitSize
	}
	return &Stack[E]{data: make([]E, size), pos: -1}
}

func (s *Stack[E]) resize() {
	l := len(s.data)
	newData := make([]E, l*2)
	copy(newData, s.data)
	s.data = newData
}

// Push push the element into the stack
func (s *Stack[E]) Push(e E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Cap() <= 0 {
		s.resize()
	}
	s.pos++
	s.data[s.pos] = e
}

// Pop pop the element from the stack
func (s *Stack[E]) Pop() (e E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.pos >= 0 {
		e = s.data[s.pos]
		s.data[s.pos] = s.empty
		s.pos--
	}
	return
}

// Cap return capacity
func (s *Stack[E]) Cap() int {
	return len(s.data) - s.pos - 1
}

// IsEmpty return true if stack has elements
func (s *Stack[E]) IsEmpty() bool {
	return s.pos < 0
}

// Size return the stack elements size
func (s *Stack[E]) Size() int {
	return s.pos + 1
}

// Copy copy to a new stack
func (s *Stack[E]) Copy() *Stack[E] {
	data := make([]E, len(s.data))
	copy(data, s.data)
	return &Stack[E]{data: data, pos: s.pos}
}
