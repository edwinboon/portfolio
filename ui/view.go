package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/edwinboon/tui-portfolio/ui/pages"
)

const (
	minWidth  = 70
	minHeight = 20
)

func (m Model) View() tea.View {
	if m.width > 0 && (m.width < minWidth || m.height < minHeight) {
		return tea.NewView(fmt.Sprintf("  Terminal too small.\n  Please resize to at least %d×%d.", minWidth, minHeight))
	}

	switch m.currentPage {
	case PageAbout:
		return tea.NewView(pages.AboutView())
	case PageProjects:
		return tea.NewView(pages.ProjectsView())
	case PageContact:
		return tea.NewView(pages.ContactView())
	}

	labels := make([]string, len(m.menuItems))
	for i, item := range m.menuItems {
		labels[i] = item.label
	}

	return tea.NewView(pages.HomeView(m.cursor, labels))
}
