package api

import (
	"reflect"
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

func Test_createPostRequest(t *testing.T) {
	type args struct {
		accessToken string
		url         string
		request     []byte
	}
	var token = "token"

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				accessToken: token,
				url:         "https://url",
				request:     []byte(`{"name":"test"}`),
			},
			want: token,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createPostRequest(tt.args.accessToken, tt.args.url, tt.args.request)
			actual := got.Header.Get("PRIVATE-TOKEN")
			if (err != nil) != tt.wantErr {
				t.Errorf("createPostRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("createPostRequest() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestCreateProjectRequest(t *testing.T) {
	expected := `{"path":"repo-name", "namespace_id": 6947500}`
	respositoryName := "repo-name"
	group := "6947500"
	actual := string(createProjectRequest(respositoryName, group))
	if actual != expected {
		t.Errorf("createProjectRequest() got %v expected %v", actual, expected)
		return
	}

}
