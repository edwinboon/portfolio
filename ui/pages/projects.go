package pages

type Project struct {
	Name        string
	Description string
	URL         string
	Language    string
}

func ProjectsView() string {
	return "Welcome to my projects page"
}
