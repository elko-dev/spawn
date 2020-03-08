package spawnhttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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

	//TODO: why does this fail here but pass when executing?
	if reflect.DeepEqual(requestTestRequest, testRequest) {
		t.Log(formatRequest(req))
		t.Log("expected ", testRequest, " got ", requestTestRequest)
		t.Fail()
		return
	}
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
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

func TestMarshalResponse(t *testing.T) {
	response := http.Response{}
	requestBody := strings.NewReader(`{"id":"1"}`)
	requestReaderClosed := ioutil.NopCloser(requestBody)
	response.Body = requestReaderClosed
	res := Response{}
	expected := Response{ID: "1"}
	MarshalResponse(&response, &res)

	if !reflect.DeepEqual(res, expected) {
		t.Log("expected", expected, "got ", res)
		t.Fail()
		return
	}
}

type Response struct {
	ID string `json:"id"`
}
