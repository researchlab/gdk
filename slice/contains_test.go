package slice

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		sl []interface{}
		v  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sl contains v", args: args{sl: []interface{}{"mike", "jack"}, v: "mike"}, want: true},
		{name: "sl not contains v", args: args{sl: []interface{}{"mike", "jack"}, v: "tom"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.sl, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	type args struct {
		sl []int
		v  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sl contains the int v", args: args{sl: []int{1, 2}, v: 1}, want: true},
		{name: "sl doesn't contains the int v", args: args{sl: []int{1, 2}, v: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt(tt.args.sl, tt.args.v); got != tt.want {
				t.Errorf("ContainsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt64(t *testing.T) {
	type args struct {
		sl []int64
		v  int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sl contains the int64 v", args: args{sl: []int64{1, 2}, v: 1}, want: true},
		{name: "sl doesn't contains the int64 v", args: args{sl: []int64{1, 2}, v: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt64(tt.args.sl, tt.args.v); got != tt.want {
				t.Errorf("ContainsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	type args struct {
		sl []string
		v  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sl contains the string v", args: args{sl: []string{"mike", "jack"}, v: "mike"}, want: true},
		{name: "sl doesn't contains the string v", args: args{sl: []string{"mike", "jack"}, v: "tom"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.sl, tt.args.v); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
