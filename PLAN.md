# TUI Portfolio — Build Plan

## Concept

An interactive terminal portfolio accessible via `ssh edwinboon.dev`.
Built with Go, Bubble Tea and Wish.

---

## Stack

- **Go** — language
- **Bubble Tea** — TUI framework (Elm architecture: Model / Update / View)
- **Lip Gloss** — styling (colors, borders, layout)
- **Bubbles** — ready-made components (list, spinner, etc.)
- **Wish** — SSH server so visitors can connect without installing anything

---

## Elm Architecture

```
Model   → app state (current page, cursor position, data)
Init    → initial state + optional startup commands
Update  → reacts to events (keypresses) and updates the Model
View    → renders the UI based on the current Model
```

---

## Project Structure

```
tui-portfolio/
├── main.go              # entrypoint, starts the app
├── go.mod
├── go.sum
└── ui/
    ├── model.go         # Model struct and Init()
    ├── update.go        # Update() function and key handling
    ├── view.go          # View() function and page routing
    └── pages/
        ├── styles.go    # Lip Gloss styles and colors
        ├── home.go      # welcome screen + menu
        ├── about.go     # about me
        ├── projects.go  # projects list
        ├── blog.go      # blog posts
        └── contact.go   # links / contact
```

---

## Steps

### Step 1 — Initialize Go project ✅

```bash
go mod init github.com/edwinboon/tui-portfolio
go get charm.land/bubbletea/v2
go get charm.land/lipgloss/v2
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/wish
```

### Step 2 — Basic TUI without SSH ✅

- `main.go` starts the app directly (no Wish, just `tea.NewProgram`)
- `model.go` — Model with a `currentPage` field and menu items
- `update.go` — navigate with `j/k`, enter to open a page, `q` to quit/go back
- `view.go` — renders pages via a switch on `currentPage`

### Step 3 — Add pages ✅ (partial)

- Home (menu) ✅
- About (static text) ✅
- Contact ✅
- Projects (list with Bubbles `list` component) — todo
- Blog — todo

### Step 4 — Styling with Lip Gloss ✅

- Color scheme based on edwinboon.dev
- Rounded borders, padding
- Active menu item highlight
- Consistent page style via `StyleBorder` and `StylePage`

### Step 5 — SSH server with Wish

Wrap the Bubble Tea app in a Wish SSH server:

- Visitors connect via `ssh edwinboon.dev`
- Each connection gets its own session
- Configure port number (default 2222 or 22)

### Step 6 — Deploy

- VPS (Hetzner/DigitalOcean) or Railway
- Systemd service or Docker container
- Open port 22 or 2222

---

## Resources

- https://github.com/charmbracelet/bubbletea/tree/main/examples
- https://github.com/charmbracelet/wish/tree/main/examples
- https://charm.sh
- Tour of Go: https://go.dev/tour

---

## Inspiration

- https://www.terminal.shop — clean SSH portfolio example
- `ssh git.charm.sh` — try this yourself in your terminal
