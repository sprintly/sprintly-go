package sprintly

import (
	"fmt"
)

type MockSprintlyApi struct{}

func (a MockSprintlyApi) CreateDefect(title, description string) (string, error) {
	fmt.Println("Creating item.")
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Description: %s\n", description)
	return "path/to/item", nil
}

func (a MockSprintlyApi) ItemLink(number int) string {
	return ""
}
func (a MockSprintlyApi) AddAnnotation(number int, label, action, body string) error {
	return nil
}

func NewMockSprintlyApi() SprintlyApi {
	return MockSprintlyApi{}
}
