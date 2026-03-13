# edwinboon.dev

An interactive terminal portfolio built with Go and Bubble Tea. Connect via SSH — no installation required.

```
ssh edwinboon.dev
```

## Stack

- [Bubble Tea v2](https://charm.land/bubbletea/v2) — TUI framework
- [Lip Gloss v2](https://charm.land/lipgloss/v2) — terminal styling
- [Wish](https://github.com/charmbracelet/wish) — SSH server

## Run locally

```bash
go run .
```

Listens on `localhost:2222` by default. Override with env vars:

```bash
HOST=0.0.0.0 PORT=2222 go run .
```

| Variable | Default     | Description        |
|----------|-------------|--------------------|
| `HOST`   | `localhost` | Bind address       |
| `PORT`   | `2222`      | Bind port (1–65535)|

> **Note:** binding to port 22 requires root or elevated privileges (e.g. `sudo` or `systemd` with `AmbientCapabilities`). Use a port above 1024 for local development.

## Navigate

| Key     | Action          |
|---------|-----------------|
| `j/k`   | Move up / down  |
| `enter` | Select page     |
| `q`     | Back / quit     |
