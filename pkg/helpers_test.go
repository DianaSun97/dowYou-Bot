package pkg

import "testing"

func TestGetProgressBar(t *testing.T) {
	type args struct {
		currentValue int
		finalValue   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProgressBar(tt.args.currentValue, tt.args.finalValue); got != tt.want {
				t.Errorf("GetProgressBar() = %v, want %v", got, tt.want)
			}
		})
	}
}
