package appcenter

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
func TestGetClientByAPIURLReturnsClientWithFullURL(t *testing.T) {
	accessToken := "TOKEN"
	orgURL := "org"
	fullURL := "https://api.appcenter.ms/v0.1/" + orgURL
	connection := NewConnection(accessToken)
	actual := connection.GetClientByAPIURL(orgURL).baseURL
	if actual != fullURL {
		t.Log("Got baseUrl "+actual, " expected "+fullURL)
		t.Fail()
	}
}
