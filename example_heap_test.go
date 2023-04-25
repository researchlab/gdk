package gdk_test

import (
	"fmt"

	"github.com/researchlab/gdk"
)

type student struct {
	name string
	age  int
}

func ExampleNewHeap() {
	h := gdk.NewHeap([]student{}, func(s1, s2 student) int {
		return s1.age - s2.age // age 小的先出
	})
	for i := 100; i > 0; i-- {
		h.Push(student{fmt.Sprintf("name%d", i), i})
	}
	// 每次获取最小年龄的student
	student := h.Pop()
	fmt.Println(student.age)
	student = h.Pop()
	fmt.Println(student.age)
	// Output:
	// 1
	// 2
}

// An IntHeap is a min-heap of ints
type IntHeap []int

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	h := gdk.NewHeap(IntHeap{2, 1, 5}, func(i1, i2 int) int { return i1 - i2 })
	h.Push(3)
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	// Output:
	// 1 2 3 5
}
