package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	s := "use the j and k keys to navigate, and enter to select a page\n\n"

	for i, item := range m.menuItems {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, item.label)
	}

	return tea.NewView(s)
}
