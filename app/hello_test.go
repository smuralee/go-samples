package app

import "testing"

func Test_FindMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find_min_1", args{2, 3}, 2},
		{"find_min_2", args{3, 2}, 2},
		{"find_min_3", args{3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMin(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
