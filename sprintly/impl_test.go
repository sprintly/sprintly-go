package sprintly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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

type values map[string]string

// blatantly copy + pasted from go-github
func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	err := r.ParseForm()
	if err != nil {
		t.Errorf("Error parsing form: %s\n", err)
	}
	if !reflect.DeepEqual(want, r.Form) {
		t.Errorf("Request parameters = %v, want %v", r.Form, want)
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
		err := r.ParseForm()
		if err != nil {
			t.Errorf("Error parsing form results: %s\n", err)
			return
		}
		testFormValues(t, r, values{
			"description": "Sickness is a defect in humans.",
			"tags":        "uservoice",
			"title":       "Sickness.",
			"type":        "defect",
		})
		fmt.Fprint(w, `{"number":1492}`)
	})

	url, err := client.CreateDefect("Sickness.", "Sickness is a defect in humans.")
	if err != nil {
		t.Error(err)
	}

	if url != client.ItemLink(1492) {
		t.Errorf("Incorrect defect url. %v (expected) != %v", client.ItemLink(1492), url)
	}
}

func TestAddAnnotation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/products/1/items/1234/annotations.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		err := r.ParseForm()
		if err != nil {
			t.Errorf("Error parsing form: %s\n", err)
		}
		testFormValues(t, r, values{
			"label":  "lab",
			"action": "act",
			"body":   "bdy",
		})
	})

	err := client.AddAnnotation(1234, "lab", "act", "bdy")
	if err != nil {
		t.Errorf("Error adding annotation: %s\n", err)
	}
}

func TestMockDefectUrl(t *testing.T) {
	// I really just want higher coverage.
	a := NewMockSprintlyApi()
	url, err := a.CreateDefect("a", "b")
	if err != nil {
		t.Error(err)
	}
	if url != "path/to/item" {
		t.Errorf("Expected the url to be `path/to/item`, not %s\n", url)
	}
}
