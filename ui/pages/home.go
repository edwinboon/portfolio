package pages

var (
	title    = RenderMarkdown("# Edwin Boon")
	subtitle = StyleSubtitle.Render("Software Engineer · edwinboon.dev")
)

func HomeView(cursor int, menuItems []string) string {
	menu := ""
	for i, item := range menuItems {
		if cursor == i {
			menu += StyleMenuCursor.Render("  ❯ "+item) + "\n"
		} else {
			menu += StyleMenuItem.Render(item) + "\n"
		}
	}

	hint := "  " + Hint("j/k", "navigate") + "   " + Hint("enter", "select") + "   " + Hint("q", "quit")

	content := title + subtitle + "\n\n" + menu + "\n" + hint
	return StyleBorder.Render(content)
}
