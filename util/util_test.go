package util

import "testing"

func TestConstructEndpoint(t *testing.T) {
	type args struct {
		basePath string
		pathName string
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
			if got := ConstructEndpoint(tt.args.basePath, tt.args.pathName); got != tt.want {
				t.Errorf("ConstructEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
