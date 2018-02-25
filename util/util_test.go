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

func TestEncryptPassword(t *testing.T) {
	type args struct {
		password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Hashed Password creation",
			args: args{
				password: "T0mmyP@$$",
			},
			wantErr: false,
		},
		{
			name: "Test Empty Hashed Password creation",
			args: args{
				password: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptPassword(tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == "" && err != nil {
				t.Errorf("EncryptPassword() = %v, want hashed password", got)
			}
		})
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}

	password := "T0mmyP@$$"
	hashedPassword, _ := EncryptPassword(password)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test hashed password and actual password",
			args: args{
				hashedPassword: hashedPassword,
				password:       password,
			},
			wantErr: false,
		},
		{
			name: "Test failed hashed password and actual password",
			args: args{
				hashedPassword: hashedPassword,
				password:       "dev",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompareHashAndPassword(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("CompareHashAndPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
