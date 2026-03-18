package pages

import "strings"

func ContactView() string {
	contactTitle := `# Contact Me`
	markdown := RenderMarkdown(contactTitle)

	email := linkView("Email:", "hello@edwinboon.dev", "mailto:hello@edwinboon.dev")
	github := linkView("GitHub:", "github.com/edwinboon", "https://github.com/edwinboon")
	website := linkView("Website:", "edwinboon.dev", "https://edwinboon.dev")

	hint := Hint("q", "back to menu")

	contact := strings.Join([]string{
		strings.TrimRight(markdown, "\n"), email, github, website, "",
		hint,
	}, "\n")

	return StylePage.Render(contact)
}
