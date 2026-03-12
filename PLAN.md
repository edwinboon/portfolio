# TUI Portfolio — Bouwplan

## Concept

Een interactief terminal portfolio bereikbaar via `ssh edwinboon.dev`.
Gebouwd met Go, Bubble Tea en Wish.

---

## Stack

- **Go** — taal
- **Bubble Tea** — TUI framework (Elm-architectuur: Model / Update / View)
- **Lip Gloss** — styling (kleuren, borders, layout)
- **Bubbles** — kant-en-klare componenten (lijst, spinner, etc.)
- **Wish** — SSH server zodat mensen kunnen verbinden zonder iets te installeren

---

## Elm-architectuur (belangrijk om te begrijpen)

```
Model   → de state van je app (welke pagina, cursor positie, data)
Init    → beginstaat + eventuele startup commands
Update  → reageert op events (toetsaanslagen) en past Model aan
View    → rendert de UI op basis van het huidige Model
```

---

## Mapstructuur

```
tui-portfolio/
├── main.go              # entrypoint, start SSH server (Wish)
├── go.mod
├── go.sum
└── ui/
    ├── model.go         # Model struct en Init()
    ├── update.go        # Update() functie en key handling
    ├── view.go          # View() functie en layout
    ├── styles.go        # Lip Gloss stijlen
    └── pages/
        ├── home.go      # welkomstscherm + menu
        ├── about.go     # over mij
        ├── projects.go  # projecten lijst
        ├── blog.go      # blog posts
        └── contact.go   # links / contact
```

---

## Stap-voor-stap bouwen

### Stap 1 — Go project initialiseren

```bash
go mod init github.com/edwinboon/tui-portfolio
go get charm.land/bubbletea/v2
go get charm.land/lipgloss/v2
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/wish
```

### Stap 2 — Basis TUI zonder SSH

Maak een werkende Bubble Tea app lokaal:

- `main.go` start de app direct (zonder Wish, gewoon `tea.NewProgram`)
- `model.go` — Model met een `currentPage` veld en een menu met opties
- `update.go` — navigeer met `j/k` of pijltjes, enter om pagina te openen, `q` om te quiten
- `view.go` — render het menu

### Stap 3 — Pagina's toevoegen

Voeg één voor één pagina's toe:

- Home (menu)
- About (statische tekst)
- Projects (lijst met Bubbles `list` component)
- Blog
- Contact

### Stap 4 — Styling met Lip Gloss

- Kleurenschema kiezen
- Borders, padding, margins
- Highlight actief menu-item
- Header en footer

### Stap 5 — SSH server met Wish

Wrap de Bubble Tea app in een Wish SSH server:

- Mensen verbinden via `ssh edwinboon.dev`
- Elke connectie krijgt een eigen sessie
- Configureer poortnummer (standaard 2222 of 22)

### Stap 6 — Deployen

- VPS (Hetzner/DigitalOcean) of Railway
- Systemd service of Docker container
- Poort 22 of 2222 openzetten

---

## Leerzame bronnen

- https://github.com/charmbracelet/bubbletea/tree/main/examples
- https://github.com/charmbracelet/wish/tree/main/examples
- https://charm.sh
- Tour of Go: https://go.dev/tour

---

## Inspiratie

- https://www.terminal.shop — clean SSH portfolio voorbeeld
- `ssh git.charm.sh` — probeer dit zelf eens in je terminal
