package oop

import (
	"fmt"
	"testing"
)

func TestNewFuncs01(t *testing.T) {
	nfuncs := NewFuncs()
	tests := []struct {
		Name    string
		F       interface{}
		inputs  []interface{}
		outputs []interface{}
	}{
		{"Add", Add, []interface{}{1, 2}, []interface{}{3}},
		{"sum", sum, []interface{}{1, 2, 3}, []interface{}{6, true}},
	}
	for _, test := range tests {
		if err := nfuncs.Bind(test.Name, test.F); err != nil {
			t.Fatalf("Bind invalid, test:%v, err:%v", test, err)
		}
		if outputs, err := nfuncs.Call(test.Name, test.inputs...); err != nil || len(outputs) != len(test.outputs) {
			t.Fatalf("Call invalid, test:%v, err:%v", test, err)
		} else {
			for i, v := range outputs {
				if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", test.outputs[i]) {
					t.Fatalf("Call invalid, out:%v, expected:%v", v, test.outputs[i])
				}
			}
		}
	}
	// negative cases
	if err := nfuncs.Bind("Add", nil); err == nil {
		t.Fatal("Bind invalid")
	}
	if outputs, err := nfuncs.Call("unknown", []interface{}{}); err == nil {
		t.Fatalf("Call invalid, outputs:%v", outputs)
	}
	if outputs, err := nfuncs.Call("Add", []interface{}{}); err == nil {
		t.Fatalf("Call invalid, outputs:%v", outputs)
	}
}

func Add(a, b int) int {
	return a + b
}

func sum(c, d, e int) (int, bool) {
	return c + d + e, true
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) ShowName() string {
	return p.Name
}

func (p Person) ShowAge() int {
	return p.Age
}

func TestNewFuncs02(t *testing.T) {
	persons := []Person{
		{"mike", 10},
		{"mike", 20},
		{"jack", 10},
	}
	nfuncs := NewFuncs()
	for _, person := range persons {
		nfuncs.Bind("name", person.ShowName)
		nfuncs.Bind("age", person.ShowAge)
		if out, err := nfuncs.Call("name"); err != nil || fmt.Sprintf("%v", out[0]) != person.Name {
			t.Fatalf("Call invalid, out:%v, expected:%v, err:%v", out, person.Name, err)
		}
		if out, err := nfuncs.Call("age"); err != nil || fmt.Sprintf("%v", out[0]) != fmt.Sprintf("%v", person.Age) {
			t.Fatalf("Call invalid, out:%v, expected:%v, err:%v", out, person.Age, err)
		}
	}
}
