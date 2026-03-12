package pages

import lip "charm.land/lipgloss/v2"

var (
	colorText    = lip.Color("#e2e8f0")
	colorMuted   = lip.Color("#94a3b8")
	colorAccent  = lip.Color("#2dd4bf")
	colorAccent2 = lip.Color("#14b8a6")
	colorBorder  = lip.Color("#0f766e")

	StyleTitle = lip.NewStyle().
			Foreground(colorAccent).
			Bold(true).
			MarginBottom(0)

	StyleSubtitle = lip.NewStyle().
			Foreground(colorMuted)

	StyleMenuItem = lip.NewStyle().
			Foreground(colorText).
			PaddingLeft(2)

	StyleMenuCursor = lip.NewStyle().
			Foreground(colorAccent2).
			Bold(true)

	StyleBody = lip.NewStyle().
			Foreground(colorText)

	StyleMuted = lip.NewStyle().
			Foreground(colorMuted)

	StyleLink = lip.NewStyle().
			Foreground(colorAccent2).
			Underline(true)

	StyleBorder = lip.NewStyle().
			Border(lip.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1, 4).
			Width(54)

	StylePage = lip.NewStyle().
			Border(lip.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1, 4).
			Width(64)
)
