package pages

import "strings"

func ContactView() string {
	contactTitle := `# Contact Me`
	markdown := RenderMarkdown(contactTitle)

	email := LinkView("Email:", "hello@edwinboon.dev", "mailto:hello@edwinboon.dev")
	github := LinkView("Github:", "github.com/edwinboon", "https://github.com/edwinboon")
	website := LinkView("Website:", "edwinboon.dev", "https://edwinboon.dev")

	hint := Hint("q", "back to menu")

	contact := strings.Join([]string{
		strings.TrimRight(markdown, "\n"), email, github, website, "",
		hint,
	}, "\n")

	return StylePage.Render(contact)
}
