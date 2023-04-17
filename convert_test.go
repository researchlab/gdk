package gdk

import (
	"reflect"
	"testing"
)

func TestToInt64(t *testing.T) {
	tests := []struct {
		in       interface{}
		expected int64
		err      error
	}{
		{int8(15), 15, nil},
		{15, 15, nil},
		{uint(15), 15, nil},
		{uint8(15), 15, nil},
		{uint16(15), 15, nil},
		{uint32(15), 15, nil},
		{uint64(15), 15, nil},
		{"one", 0, ERR_NEED_NUMERIC},
	}
	for _, test := range tests {
		if out, err := ToInt64(test.in); out != test.expected || (err != nil && err != test.err) {
			t.Errorf("ToInt64 invalid, test:%v, out:%v, err:%v", test, out, err)
		}
	}
}

func TestBytesToFloat64(t *testing.T) {
	tests := []struct {
		in       []byte
		expected float64
		err      error
	}{
		{Float64ToBytes(10.21), 10.21, nil},
		{[]byte("20.10"), 0, ERR_BYTES_INVALILD},
	}
	for _, test := range tests {
		if out, err := BytesToFloat64(test.in); out != test.expected || err != nil && err != ERR_BYTES_INVALILD {
			t.Errorf("BytesToFloat64 invalid, test:%v, out:%v, err:%v", test, out, err)
		}
	}
}

func TestFloat64ToStr(t *testing.T) {
	tests := []struct {
		in       float64
		prec     int
		expected string
	}{
		{10.21, 3, "10.210"},
		{10.21, 1, "10.2"},
		{10.21, 2, "10.21"},
		{10.215, 3, "10.215"},
		{10.215, 2, "10.21"},
	}
	for _, test := range tests {
		if out := Float64ToStr(test.in, test.prec); out != test.expected {
			t.Errorf("Float64ToStr invalid, test:%v, out:%v", test, out)
		}
	}
}

func TestStrToFloat64(t *testing.T) {
	tests := []struct {
		in        string
		precision int
		expected  float64
	}{
		{"10.211", 2, 10.21},
		{"10.218", 3, 10.218},
		{"10.211", 1, 10.2},
		{"10.211", 0, 10},
		{"10.211", 5, 10.21100},
	}
	for _, test := range tests {
		if out := StrToFloat64(test.in, test.precision); out != test.expected {
			t.Errorf("StrToFloat64 invalid, test:%v, out:%v", test, out)
		}
	}
}

func TestStrToFloat64Round(t *testing.T) {
	tests := []struct {
		in        string
		precision int
		round     bool
		expected  float64
	}{
		{"10.211", 2, false, 10.21},
		{"10.218", 3, true, 10.218},
		{"10.213", 2, true, 10.21},
		{"10.218", 2, true, 10.22},
		{"10.218", 1, true, 10.2},
		{"10.211", 0, true, 10},
		{"10.211", 5, false, 10.21100},
	}
	for _, test := range tests {
		if out := StrToFloat64Round(test.in, test.precision, test.round); out != test.expected {
			t.Errorf("StrToFloat64Round invalid, test:%v, out:%v", test, out)
		}
	}
}

func TestFloatPrecision(t *testing.T) {
	tests := []struct {
		in        float64
		precision int
		round     bool
		expected  float64
	}{
		{10.211, 2, false, 10.21},
		{10.218, 3, true, 10.218},
		{10.213, 2, true, 10.21},
		{10.218, 2, true, 10.22},
		{10.218, 1, true, 10.2},
		{10.211, 0, true, 10},
		{10.211, 5, false, 10.21100},
	}
	for _, test := range tests {
		if out := Float64Precision(test.in, test.precision, test.round); out != test.expected {
			t.Errorf("Float64Precision invalid, test:%v, out:%v", test, out)
		}
	}
}

type Person struct {
	Name   string
	Age    int
	Height int
}

type Human struct {
	name   string
	Age    int
	Height int
	Score  *int
}

func TestStructToMap(t *testing.T) {
	tests := []struct {
		in       Person
		expected map[string]interface{}
	}{
		{in: Person{"mike", 10, 10}, expected: map[string]interface{}{
			"Age":    10,
			"Height": 10,
			"Name":   "mike",
		}},
		{in: Person{"jack", 1, 10}, expected: map[string]interface{}{
			"Age":    1,
			"Height": 10,
			"Name":   "jack",
		}},
	}
	for _, test := range tests {
		if out := StructToMap(test.in); !reflect.DeepEqual(out, test.expected) {
			t.Errorf("StructToMap invalid, test:%v, out:%v", test, out)
		}
	}
}

func TestMapToStruct(t *testing.T) {
	tests := []struct {
		in       map[string]interface{}
		expected *Person
	}{
		{expected: &Person{"mike", 10, 10}, in: map[string]interface{}{
			"Age":    10,
			"Height": 10,
			"Name":   "mike",
		}},
		{expected: &Person{"jack", 10, 102}, in: map[string]interface{}{
			"Age":    10,
			"Height": 102,
			"Name":   "jack",
		}},
		{expected: &Person{"mike", 0, 10}, in: map[string]interface{}{
			"Age":    nil,
			"Height": 10,
			"Name":   "mike",
		}},
	}
	for _, test := range tests {
		if out, err := MapToStruct(test.in, &Person{}); err != nil || !reflect.DeepEqual(out, test.expected) {
			t.Errorf("MapToStruct invalid, test:%v, out:%v, err:%v", test, out, err)
		}
	}
	negativetests := []struct {
		in       map[string]interface{}
		data     interface{}
		expected *Person
	}{
		{data: &Person{}, expected: nil, in: map[string]interface{}{
			"Ages":   10,
			"Height": 102,
			"Address": map[string]interface{}{
				"Address": "china",
				"code":    00000,
			},
		}},
		{data: &Human{}, expected: nil, in: map[string]interface{}{
			"name":   "mike",
			"Age":    10,
			"Height": 102,
		}},
		{data: &Human{}, expected: nil, in: map[string]interface{}{
			"Score": "string",
		}},
	}
	for _, test := range negativetests {
		if out, err := MapToStruct(test.in, test.data); err == nil {
			t.Errorf("MapToStruct invalid, test:%v, out:%v", test, out)
		}
	}
}
