package sprintly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client SprintlyClient
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewSprintlyClient(
		"user@example.org",
		"test_api_key",
		1, // ProductId
	).(SprintlyClient)

	url, _ := url.Parse(server.URL)
	client.BaseUrl = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, method string) {
	if r.Method != method {
		t.Errorf("Method mismatch. %v (expected) != %v", method, r.Method)
	}
}

func TestNewClient(t *testing.T) {
	// Maybe too many implementation details?
	c := NewSprintlyClient("user@example.org", "key", 1234).(SprintlyClient)

	expected := "https://sprint.ly"
	if c.BaseUrl.String() != expected {
		t.Errorf("NewClient BaseUrl was %v not %v", c.BaseUrl, expected)
	}
}

func TestNewRequest(t *testing.T) {
	req, err := client.newRequest("POST", "foo", nil)
	if err != nil {
		t.Errorf("Error fetching new request: %s\n", err)
	}

	if req.Header["Authorization"] == nil {
		t.Error("Expected Authorization header to be set for all API requests.")
	}
}

func TestDefectCreation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/products/1/items.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		// TODO: test correct form values
		fmt.Fprint(w, `{"number":1492}`)
	})

	url, err := client.CreateDefect("Sickeness.", "Sickness is a defect in humans.")
	if err != nil {
		t.Error(err)
	}

	if url != client.ItemLink(1492) {
		t.Errorf("Incorrect defect url. %v (expected) != %v", client.ItemLink(1492), url)
	}
}
