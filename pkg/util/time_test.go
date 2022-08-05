package util

import (
	"testing"
	"time"
)

func TestTimeFormatNow(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				"YYYY-MM-DD",
			},
			want: time.Now().Format("2006-01-02"),
		},
		{
			name: "ok-hour",
			args: args{
				"YYYY-MM-DD:hh",
			},
			want: time.Now().Format("2006-01-02:15"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeFormatNow(tt.args.pattern)
			if got != tt.want {
				t.Errorf("TimeFormatNow() = %v, want %v", got, tt.want)
			}
		})
	}
}
