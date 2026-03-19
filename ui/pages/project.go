package pages

type Project struct {
	Name        string
	Description string
	URL         string
	Languages   []string
}

func ProjectView(p Project) string {
	return "Welcome to my projects page"
}
