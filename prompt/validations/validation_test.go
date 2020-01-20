package validations

import (
	"testing"
)

func TestYOrNValidation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not y or n",
			args: args{
				input: "hello",
			},
			wantErr: true,
		},
		{
			name: "Upper case Y",
			args: args{
				input: "Y",
			},
			wantErr: false,
		},
		{
			name: "Lower case y",
			args: args{
				input: "y",
			},
			wantErr: false,
		},
		{
			name: "Lower case n",
			args: args{
				input: "n",
			},
			wantErr: false,
		},
		{
			name: "Upper case N",
			args: args{
				input: "N",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := YOrNValidation(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("YOrNValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitValidation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Ends in .git",
			args: args{
				input: "https://github.com/spawn.git",
			},
			wantErr: false,
		},
		{
			name: "Does not End in .git",
			args: args{
				input: "https://github.com/spawn",
			},
			wantErr: true,
		},
		{
			name: "Empty string returns error",
			args: args{
				input: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GitValidation(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("GitValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
