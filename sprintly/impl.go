package sprintly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SprintlyClient struct {
	Email     string
	ApiKey    string
	ProductId int
	// Should I be taking an HTTP client here?
	// baseUrl?
}

type CreateItemResult struct {
	Number int `json:"number"`
}

func (a SprintlyClient) ItemLink(number int) string {
	return fmt.Sprintf("https://sprint.ly/product/%d/#!/item/%d", a.ProductId, number)
}

func (a SprintlyClient) AddAnnotation(number int, label, action, body string) error {
	v := url.Values{}
	v.Set("label", label)
	v.Set("action", action)
	v.Set("body", body)

	client := new(http.Client)
	url := fmt.Sprintf("https://sprint.ly/api/products/%d/items/%d/annotations.json", a.ProductId, number)
	req, err := http.NewRequest("POST", url, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(a.Email, a.ApiKey)

	_, err = client.Do(req)
	return err
}

func (a SprintlyClient) CreateDefect(title, description string) (string, error) {
	v := url.Values{}
	v.Set("type", "defect")
	v.Set("title", title)
	v.Set("description", description)
	v.Set("tags", "uservoice")

	client := new(http.Client)
	url := fmt.Sprintf("https://sprint.ly/api/products/%d/items.json", a.ProductId)

	req, err := http.NewRequest("POST", url, strings.NewReader(v.Encode()))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(a.Email, a.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result := CreateItemResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return a.ItemLink(result.Number), nil
}

func NewSprintlyClient(email, api_key string, product_id int) SprintlyApi {
	return SprintlyClient{email, api_key, product_id}
}