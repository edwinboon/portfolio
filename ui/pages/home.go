package pages

import (
	"fmt"
)

func HomeView(cursor int, menuItems []string) string {
	s := "use the j and k keys to navigate, and enter to select a page\n\n"

	for i, item := range menuItems {
		c := " " // no cursor
		if cursor == i {
			c = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", c, item)
	}

	return s
}
