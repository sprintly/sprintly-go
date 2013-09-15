package sprintly

type SprintlyApi interface {
	// returns the URL of the created defect.
	CreateDefect(title, description string) (string, error)
	ItemLink(number int) string
	AddAnnotation(number int, label, action, body string) error
}
