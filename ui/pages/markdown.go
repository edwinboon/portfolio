package pages

import (
	"fmt"

	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
)

func stringPtr(s string) *string {
	return &s
}

func uintPtr(i uint) *uint {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func RenderMarkdown(md string) string {
	cfg := styles.DarkStyleConfig
	// General document styling
	cfg.Document.Margin = uintPtr(0)

	// Header styling
	cfg.H1.BackgroundColor = stringPtr("#0d9488")
	cfg.H1.Color = stringPtr("#ffffff")
	cfg.H4.Prefix = ""
	cfg.H4.Color = stringPtr("#0d9488")
	cfg.H4.Bold = boolPtr(false)

	// Link styling
	cfg.Link.Color = stringPtr("#0d9488")
	cfg.LinkText.Color = stringPtr("#0d9488")

	r, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(0),
		glamour.WithStyles(cfg),
	)
	if err != nil {
		return fmt.Sprintf("Failed to create renderer: %v", err)
	}

	out, err := r.Render(md)
	if err != nil {
		return fmt.Sprintf("Failed to render the markdown content: %v", err)
	}

	return out
}
