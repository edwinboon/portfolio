package pages

import (
	"strings"

	"github.com/edwinboon/tui-portfolio/internal/models"
)

var styleLang = StyleMuted.Foreground(colorAccent2)

func projectCard(p models.Project) string {
	name := osc8Link(p.Name, p.URL)
	desc := StyleMuted.Render(p.Description)
	lang := styleLang.Render(p.Language)

	return name + "  " + lang + "\n" + desc
}

func ProjectsView(p []models.Project) string {
	title := RenderMarkdown("# Projects")

	cards := []string{}
	for _, proj := range p {
		cards = append(cards, projectCard(proj))
	}

	hint := Hint("q", "back to menu")

	content := strings.TrimRight(title, "\n") + "\n" + strings.Join(cards, "\n\n") + "\n\n" + hint
	return StylePage.Render(content)
}
