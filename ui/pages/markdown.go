package pages

import (
	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
)

func stringPtr(s string) *string {
	return &s
}

func uintPtr(i uint) *uint {
	return &i
}

func RenderMarkdown(md string) string {
	styles := styles.DarkStyleConfig
	styles.H1.BackgroundColor = stringPtr("#0d9488")
	styles.H1.Color = stringPtr("#ffffff")
	styles.Document.Margin = uintPtr(0)

	r, _ := glamour.NewTermRenderer(
		glamour.WithWordWrap(0),
		glamour.WithStyles(styles),
	)

	out, err := r.Render(md)
	if err != nil {
		return " Failed to render the markdown content."
	}

	return out
}
