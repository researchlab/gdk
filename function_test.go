package gdk

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFunctions(t *testing.T) {
	want := make(Functions, 0)
	if f := NewFunctions(); !reflect.DeepEqual(f, want) {
		t.Fatalf("NewFunctions() = %v, want:%v", f, want)
	}
}

func TestFunctionsBind(t *testing.T) {
	type args struct {
		name string
		fn   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "bind success", args: args{name: "Add", fn: func(a, b int) int { return a + b }}, wantErr: false},
		{name: "panics if the fn type's Kind is not Func", args: args{name: "nil", fn: nil}, wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := NewFunctions()
			if err := f.Bind(test.args.name, test.args.fn); (err != nil) != test.wantErr {
				t.Errorf("f.Bind() err =%v, wantErr:%v", err, test.wantErr)
			}
		})
	}
}

func TestFunctionsCall(t *testing.T) {
	type args struct {
		name   string
		params []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantR   []reflect.Value
		wantErr bool
	}{
		{name: "call success", args: args{name: "Add", params: []interface{}{1, 2}}, wantR: []reflect.Value{reflect.ValueOf(3)}, wantErr: false},
		{name: "function name doesn't exist", args: args{name: "Sum"}, wantErr: true},
		{name: "not enough params", args: args{name: "Add", params: []interface{}{1}}, wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := NewFunctions()
			f.Bind("Add", func(a, b int) int { return a + b })
			gotR, err := f.Call(test.args.name, test.args.params...)
			if (err != nil) != test.wantErr {
				t.Errorf("f.Call() err=%v, wantErr:%v", err, test.wantErr)
			}
			if fmt.Sprintf("%v", gotR) != fmt.Sprintf("%v", test.wantR) {
				t.Errorf("f.Call() gotR:%v, wantR:%v", gotR, test.wantR)
			}
		})
	}
}

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p People) ShowName() string {
	return p.Name
}

func (p People) ShowAge() int {
	return p.Age
}

func TestFunctionsCall02(t *testing.T) {
	persons := []People{
		{"mike", 10},
		{"mike", 20},
		{"jack", 10},
	}
	nfuncs := NewFunctions()
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
