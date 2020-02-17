package api

import (
	"testing"
)

func TestClient_GetUrlByPath(t *testing.T) {
	connection := NewConnection("")
	client := NewClient(connection, "https://api.appcenter.ms/v0.1/")
	expected := "https://api.appcenter.ms/v0.1/orgs"

	actual := client.GetUrlByPath("orgs")
	if actual != expected {
		t.Log("Got " + actual + " expected " + expected)
		t.Fail()
	}
}
