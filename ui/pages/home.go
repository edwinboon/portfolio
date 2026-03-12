package pages

func HomeView(cursor int, menuItems []string) string {
	name := StyleTitle.Render("Edwin Boon")
	subtitle := StyleSubtitle.Render("Software Engineer · edwinboon.dev")

	menu := ""
	for i, item := range menuItems {
		if cursor == i {
			menu += StyleMenuCursor.Render("  ❯ "+item) + "\n"
		} else {
			menu += StyleMenuItem.Render(item) + "\n"
		}
	}

	hint := StyleMuted.Render("  j/k  navigate   enter  select   q  quit")

	content := name + "\n" + subtitle + "\n\n" + menu + "\n" + hint
	return StyleBorder.Render(content)
}
