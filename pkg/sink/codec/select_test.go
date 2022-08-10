package codec

import (
	"reflect"
	"testing"
)

func TestInitMatcher(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "const",
			args: args{
				"aa-bb-cc",
			},
			want: nil,
		},
		{
			name: "ok",
			args: args{
				"aa-${bb}-${cc.dd}",
			},
			want: [][]string{
				{"${bb}", "bb"},
				{"${cc.dd}", "cc.dd"},
			},
		},
		{
			name: "ok-with-time",
			args: args{
				"aa-${bb}-${cc.dd}-${+YYYY.MM.DD}",
			},
			want: [][]string{
				{"${bb}", "bb"},
				{"${cc.dd}", "cc.dd"},
				{"${+YYYY.MM.DD}", "+YYYY.MM.DD"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InitMatcher(tt.args.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
