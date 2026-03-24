package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/edwinboon/tui-portfolio/internal/models"
)

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
	projects    []models.Project
	width       int
	height      int
}

func InitialModel(p []models.Project) Model {
	return Model{
		currentPage: PageHome,
		cursor:      0,
		projects:    p,
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
