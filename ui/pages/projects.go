package pages

type Project struct {
	Name        string
	Description string
	URL         string
	Languages   []string
}

func ProjectsView() string {
	return "Welcome to my projects page"
}
