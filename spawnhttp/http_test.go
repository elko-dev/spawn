package spawnhttp

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

type TestRequest struct {
	id   string `json:"id"`
	name string `json:"name"`
}

func Test_isSuccessStatusCode(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		struct {
			name string
			args args
			want bool
		}{
			name: "200 is success", args: args{200}, want: true,
		},
		{
			name: "201 is success", args: args{201}, want: true,
		},
		{
			name: "400 is failure", args: args{400}, want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSuccessStatusCode(tt.args.statusCode); got != tt.want {
				t.Errorf("isSuccessStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createRequest_containsapplicationjsonhead(t *testing.T) {
	url := "https://url"
	testRequest := TestRequest{
		id:   "ID",
		name: "NAME",
	}
	req, err := CreateRequest(url, testRequest)

	if err != nil {
		t.Log("Error received when none expected ", err)
		t.Fail()
		return
	}

	contentType := req.Header.Get("Content-Type")

	if contentType != "application/json" {
		t.Log("Wrong application type, got ", contentType)
		t.Fail()
		return
	}
}

func Test_createRequest_containsbody(t *testing.T) {
	url := "https://url"
	testRequest := TestRequest{
		id:   "ID",
		name: "NAME",
	}
	req, err := CreateRequest(url, testRequest)

	if err != nil {
		t.Log("Error received when none expected ", err)
		t.Fail()
		return
	}

	requestTestRequest := TestRequest{}
	body, err := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &requestTestRequest)

	if reflect.DeepEqual(requestTestRequest, testRequest) {
		t.Log("expected ", testRequest, " got ", requestTestRequest)
		t.Fail()
		return
	}
}

func Test_createRequest_handlesnilbody(t *testing.T) {
	url := "https://url"
	_, err := CreateRequest(url, nil)

	if err != nil {
		t.Log("Error received when none expected ", err)
		t.Fail()
		return
	}
}
