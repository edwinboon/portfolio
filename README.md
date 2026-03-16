# edwinboon.dev

An interactive terminal portfolio built with Go and Bubble Tea. Connect via SSH — no installation required.

```
ssh edwinboon.dev
```

## Structure

```
tui-portfolio/
├── main.go        # SSH server entrypoint
├── ui/            # TUI (Bubble Tea)
│   ├── model.go
│   ├── update.go
│   ├── view.go
│   └── pages/
└── web/           # Landing page (Astro)
    └── src/
```

## Stack

**TUI**
- [Bubble Tea v2](https://charm.land/bubbletea/v2) — TUI framework
- [Lip Gloss v2](https://charm.land/lipgloss/v2) — terminal styling
- [Wish](https://github.com/charmbracelet/wish) — SSH server

**Web**
- [Astro](https://astro.build) — static site framework
- [Tailwind CSS v4](https://tailwindcss.com) — styling

## Run locally

**TUI**

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

**Web**

```bash
cd web
npm install
npm run dev
```

## Deploy

Deployments are triggered by pushing a version tag:

```bash
git tag v1.0.0
git push --tags
```

This will:
1. Build the Go binary for Linux ARM64
2. Upload and deploy it to the VPS
3. Restart the systemd service
4. Create a GitHub Release with release notes

## Navigate

| Key     | Action          |
|---------|-----------------|
| `j/k`   | Move up / down  |
| `enter` | Select page     |
| `q`     | Back / quit     |
