package slice

import "testing"

func TestSumInt64(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name    string
		args    args
		wantSum int64
	}{
		{name: "{1,2} sum 3", args: args{s: []int64{1, 2}}, wantSum: 3},
		{name: "{1,2,3} sum 6", args: args{s: []int64{1, 2, 3}}, wantSum: 6},
		{name: "{1} sum 1", args: args{s: []int64{1}}, wantSum: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumInt64(tt.args.s); gotSum != tt.wantSum {
				t.Errorf("SumInt64() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSumInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{name: "{1,2} sum 3", args: args{s: []int{1, 2}}, wantSum: 3},
		{name: "{1,2,3} sum 6", args: args{s: []int{1, 2, 3}}, wantSum: 6},
		{name: "{1} sum 1", args: args{s: []int{1}}, wantSum: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumInt(tt.args.s); gotSum != tt.wantSum {
				t.Errorf("SumInt() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSumFloat64(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name    string
		args    args
		wantSum float64
	}{
		{name: "{1.0,2.1} sum 3.1", args: args{s: []float64{1.0, 2.1}}, wantSum: 3.1},
		{name: "{1.1,2.2,3.3} sum 6.6", args: args{s: []float64{1.1, 2.2, 3.3}}, wantSum: 6.6},
		{name: "{1.1} sum 1.1", args: args{s: []float64{1.1}}, wantSum: 1.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumFloat64(tt.args.s); gotSum != tt.wantSum {
				t.Errorf("SumFloat64() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
