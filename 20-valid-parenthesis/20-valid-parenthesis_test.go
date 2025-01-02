package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Ex 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "Ex 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "Ex 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "Ex 4",
			args: args{
				s: "([])",
			},
			want: true,
		},
		{
			name: "Ex 5",
			args: args{
				s: "([)]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
