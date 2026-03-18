package pages

import "fmt"

// linkView builds an OSC-8 hyperlink. Only call with trusted, hardcoded strings.
func linkView(label, text, url string) string {
	paddedLabel := fmt.Sprintf("%-10s", label)
	return StyleBody.Render(paddedLabel) + "\033]8;;" + url + "\033\\" + StyleLink.Render(text) +
		"\033]8;;\033\\"
}
