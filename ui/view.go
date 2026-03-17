package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/edwinboon/tui-portfolio/ui/pages"
)

const minWidth = 70
const minHeight = 20

func (m Model) View() tea.View {
	if m.width > 0 && (m.width < minWidth || m.height < minHeight) {
		return tea.NewView("  Terminal too small.\n  Please resize to at least 70×20.")
	}

	switch m.currentPage {
	case PageAbout:
		return tea.NewView(pages.AboutView())
	case PageContact:
		return tea.NewView(pages.ContactView())
	}

	labels := make([]string, len(m.menuItems))
	for i, item := range m.menuItems {
		labels[i] = item.label
	}

	return tea.NewView(pages.HomeView(m.cursor, labels))
}
