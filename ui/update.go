package ui

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if m.currentPage != PageHome {
				m.currentPage = PageHome
				return m, nil
			}

			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
		case "enter":
			m.currentPage = m.menuItems[m.cursor].page
		}
	}

	return m, nil
}
