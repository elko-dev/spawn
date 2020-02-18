package api

import (
	"testing"
)

func TestBaseUrlIsSetWhenNewConnectionIsCreated(t *testing.T) {
	accessToken := "TOKEN"
	connection := NewConnection(accessToken)
	expectedBaseURL := "https://api.appcenter.ms/v0.1/"
	actual := connection.BaseUrl
	if actual != expectedBaseURL {
		t.Log("Got baseUrl "+actual, " expected "+expectedBaseURL)
		t.Fail()
	}
}

func TestAuthorizationIsSetWhenNewConnectionIsCreated(t *testing.T) {
	accessToken := "TOKEN"
	connection := NewConnection(accessToken)
	expectedAuthorization := accessToken
	actual := connection.AuthorizationToken
	if actual != expectedAuthorization {
		t.Log("Got authorization "+actual, " expected "+expectedAuthorization)
		t.Fail()
	}
}
func TestGetClientByReturnsClientWithBaseURL(t *testing.T) {
	accessToken := "TOKEN"
	baseURL := "https://api.appcenter.ms/v0.1/"
	connection := NewConnection(accessToken)
	actual := connection.GetClient().baseURL
	if actual != baseURL {
		t.Log("Got baseUrl "+actual, " expected "+baseURL)
		t.Fail()
	}
}
