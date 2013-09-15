package sprintly

type MockSprintlyApi struct{}

func (a MockSprintlyApi) CreateDefect(title, description string) (string, error) {
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
