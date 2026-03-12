package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/edwinboon/tui-portfolio/ui"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
