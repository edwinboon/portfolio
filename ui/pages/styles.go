package pages

import lip "charm.land/lipgloss/v2"

var (
	colorText    = lip.Color("#e2e8f0")
	colorMuted   = lip.Color("#9ca3af")
	colorAccent  = lip.Color("#0d9488")
	colorAccent2 = lip.Color("#14b8a6")
	colorBorder  = lip.Color("#0d9488")

	StyleTitle = lip.NewStyle().
			Foreground(colorText).
			Background(colorAccent).
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
			Foreground(colorAccent).
			Underline(true)

	StyleKey = lip.NewStyle().
			Foreground(colorAccent2)

	StyleBorder = lip.NewStyle().
			Border(lip.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1, 4).
			Width(64)

	StylePage = lip.NewStyle().
			Border(lip.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1, 4).
			Width(64)
)

func Hint(key, label string) string {
	return StyleKey.Render(key) + StyleMuted.Render(" "+label)
}
