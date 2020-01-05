package api

import (
	"testing"
)

func Test_isSuccessStatusCode(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				statusCode: 200,
			},
			want: true,
		},
		{
			args: args{
				statusCode: 201,
			},
			want: true,
		},
		{
			args: args{
				statusCode: 400,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSuccessStatusCode(tt.args.statusCode); got != tt.want {
				t.Errorf("isSuccessStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUnauthorized(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				statusCode: 401,
			},
			want: true,
		},
		{
			args: args{
				statusCode: 200,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnauthorized(tt.args.statusCode); got != tt.want {
				t.Errorf("isUnauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}
