package pages

func ContactView() string {
	title := StyleTitle.Render("Contact")

	links := StyleMuted.Render("  Email    ") + StyleLink.Render("hello@edwinboon.dev") + "\n" +
		StyleMuted.Render("  GitHub   ") + StyleLink.Render("github.com/edwinboon") + "\n" +
		StyleMuted.Render("  Website  ") + StyleLink.Render("edwinboon.dev")

	hint := StyleMuted.Render("  q  back to menu")

	content := title + "\n\n" + links + "\n\n" + hint
	return StylePage.Render(content)
}
