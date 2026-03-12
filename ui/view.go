package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/edwinboon/tui-portfolio/ui/pages"
)

func (m Model) View() tea.View {
	if m.currentPage != PageHome {
		return tea.NewView("Page: " + string(m.currentPage))
	}

	labels := make([]string, len(m.menuItems))
	for i, item := range m.menuItems {
		labels[i] = item.label
	}

	return tea.NewView(pages.HomeView(m.cursor, labels))
}
