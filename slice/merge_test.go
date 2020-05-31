package slice

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		slice1 []interface{}
		slice2 []interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantC []interface{}
	}{
		{name: "{1,2} merge {3,5} to {1,2,3,5}", args: args{slice1: []interface{}{1, 2}, slice2: []interface{}{3, 5}}, wantC: []interface{}{1, 2, 3, 5}},
		{name: "{1,2,3} merge {3,5} to {1,2,3,3,5}", args: args{slice1: []interface{}{1, 2, 3}, slice2: []interface{}{3, 5}}, wantC: []interface{}{1, 2, 3, 3, 5}},
		{name: "{1} merge {2,3,5} to {1,2,3,5}", args: args{slice1: []interface{}{1}, slice2: []interface{}{2, 3, 5}}, wantC: []interface{}{1, 2, 3, 5}},
		{name: "{1} merge {3,5} to {1,3,5}", args: args{slice1: []interface{}{1}, slice2: []interface{}{3, 5}}, wantC: []interface{}{1, 3, 5}},
		{name: "{1} merge {3} to {1,3}", args: args{slice1: []interface{}{1}, slice2: []interface{}{3}}, wantC: []interface{}{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Merge(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Merge() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestMergeInt(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name  string
		args  args
		wantC []int
	}{
		{name: "{1,2} merge {3,5} to {1,2,3,5}", args: args{slice1: []int{1, 2}, slice2: []int{3, 5}}, wantC: []int{1, 2, 3, 5}},
		{name: "{1,2,3} merge {3,5} to {1,2,3,3,5}", args: args{slice1: []int{1, 2, 3}, slice2: []int{3, 5}}, wantC: []int{1, 2, 3, 3, 5}},
		{name: "{1} merge {2,3,5} to {1,2,3,5}", args: args{slice1: []int{1}, slice2: []int{2, 3, 5}}, wantC: []int{1, 2, 3, 5}},
		{name: "{1} merge {3,5} to {1,3,5}", args: args{slice1: []int{1}, slice2: []int{3, 5}}, wantC: []int{1, 3, 5}},
		{name: "{1} merge {3} to {1,3}", args: args{slice1: []int{1}, slice2: []int{3}}, wantC: []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := MergeInt(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("MergeInt() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestMergeInt64(t *testing.T) {
	type args struct {
		slice1 []int64
		slice2 []int64
	}
	tests := []struct {
		name  string
		args  args
		wantC []int64
	}{
		{name: "{1,2} merge {3,5} to {1,2,3,5}", args: args{slice1: []int64{1, 2}, slice2: []int64{3, 5}}, wantC: []int64{1, 2, 3, 5}},
		{name: "{1,2,3} merge {3,5} to {1,2,3,3,5}", args: args{slice1: []int64{1, 2, 3}, slice2: []int64{3, 5}}, wantC: []int64{1, 2, 3, 3, 5}},
		{name: "{1} merge {2,3,5} to {1,2,3,5}", args: args{slice1: []int64{1}, slice2: []int64{2, 3, 5}}, wantC: []int64{1, 2, 3, 5}},
		{name: "{1} merge {3,5} to {1,3,5}", args: args{slice1: []int64{1}, slice2: []int64{3, 5}}, wantC: []int64{1, 3, 5}},
		{name: "{1} merge {3} to {1,3}", args: args{slice1: []int64{1}, slice2: []int64{3}}, wantC: []int64{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := MergeInt64(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("MergeInt64() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestMergeString(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name  string
		args  args
		wantC []string
	}{
		{name: "{\"mike\",\"jack\"} merge {\"tom\"} to {\"mike\",\"jack\",\"tom\"}", args: args{slice1: []string{"mike", "jack"}, slice2: []string{"tom"}}, wantC: []string{"mike", "jack", "tom"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := MergeString(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("MergeString() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
