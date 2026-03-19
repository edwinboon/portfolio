package ui

import tea "charm.land/bubbletea/v2"

type Page string

type MenuItem struct {
	label string
	page  Page
}

const (
	PageHome     Page = "home"
	PageAbout    Page = "about"
	PageProjects Page = "projects"
	PageContact  Page = "contact"
)

type Model struct {
	currentPage Page
	cursor      int
	menuItems   []MenuItem
	width       int
	height      int
}

func InitialModel() Model {
	return Model{
		currentPage: PageHome,
		cursor:      0,
		menuItems: []MenuItem{
			{label: "About", page: PageAbout},
			{label: "Projects", page: PageProjects},
			{label: "Contact", page: PageContact},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
