package util

import (
	"testing"
)

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
		{
			name: "Test endpoint construction",
			args: args{
				basePath: "/api/v1",
				pathName: "/user/{id}",
			},
			want: "/api/v1/user/{id}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructEndpoint(tt.args.basePath, tt.args.pathName); got != tt.want {
				t.Errorf("ConstructEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateOTP(t *testing.T) {
	t.Run("Test OTP generation", func(t *testing.T) {
		if got := GenerateOTP(); got == 0 {
			t.Errorf("GenerateOTP() = %v, want 4 digit random number", got)
		}
	})
}
