package slice

import (
	"fmt"
	"testing"
)

func TestUniqueInt64(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "{1,2,3,3} unique to {1,2,3}", args: args{s: []int64{1, 2, 3, 3}}, want: []int64{1, 2, 3}},
		{name: "{1,2,3} unique to {1,2,3}", args: args{s: []int64{1, 2, 3}}, want: []int64{1, 2, 3}},
		{name: "{1,1,1} unique to {1}", args: args{s: []int64{1, 1, 1}}, want: []int64{1}},
		{name: "{} unique to {}", args: args{s: []int64{}}, want: []int64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt64(tt.args.s); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("UniqueInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "{1,2,3,3} unique to {1,2,3}", args: args{s: []int{1, 2, 3, 3}}, want: []int{1, 2, 3}},
		{name: "{1,2,3} unique to {1,2,3}", args: args{s: []int{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "{1,1,1} unique to {1}", args: args{s: []int{1, 1, 1}}, want: []int{1}},
		{name: "{} unique to {}", args: args{s: []int{}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt(tt.args.s); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("UniqueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "{\"mike\",\"jack\",\"mike\"} unique to {\"mike\",\"jack\"}", args: args{s: []string{"mike", "jack", "mike"}}, want: []string{"mike", "jack"}},
		{name: "{\"mike\",\"mike\"} unique to {\"mike\"}", args: args{s: []string{"mike", "mike"}}, want: []string{"mike"}},
		{name: "{} unique to {}", args: args{s: []string{}}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueString(tt.args.s); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("UniqueString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		expected interface{}
	}{
		{name: "{\"mike\",\"jack\",\"mike\"} unique to {\"mike\",\"jack\"}", args: args{data: []string{"mike", "jack", "mike"}}, want: true, expected: []interface{}{"mike", "jack"}},
		{name: "{\"mike\",\"mike\"} unique to {\"mike\"}", args: args{data: []string{"mike", "mike"}}, want: true, expected: []interface{}{"mike"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(&tt.args.data); got != tt.want {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
			if fmt.Sprintf("%v", tt.args.data) != fmt.Sprintf("%v", tt.expected) {
				t.Errorf("Unique() = %v, want:%v", tt.args.data, tt.expected)
			}
		})
	}

	negativetests := []struct {
		name     string
		args     args
		want     bool
		expected interface{}
	}{
		{name: "data is not ptr", args: args{data: []string{"mike", "jack", "mike"}}, want: false, expected: nil},
		{name: "data type is not slice or array", args: args{data: &map[string]interface{}{"name": "mike"}}, want: true, expected: nil},
	}
	for _, tt := range negativetests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.data); got != tt.want {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}

		})
	}

}
