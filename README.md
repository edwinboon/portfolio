# edwinboon.dev

An interactive terminal portfolio built with Go and Bubble Tea. Connect via SSH — no installation required.

```
ssh edwinboon.dev
```

## Structure

```
tui-portfolio/
├── main.go              # SSH server entrypoint
├── internal/
│   ├── github/          # GitHub API client
│   └── models/          # Shared data models
├── ui/                  # TUI (Bubble Tea)
│   ├── model.go
│   ├── update.go
│   ├── view.go
│   └── pages/
└── web/                 # Landing page (Astro)
    ├── src/
    │   ├── pages/
    │   └── styles/
    └── public/
```

## Stack

**TUI**
- [Bubble Tea v2](https://charm.land/bubbletea/v2) — TUI framework
- [Lip Gloss v2](https://charm.land/lipgloss/v2) — terminal styling
- [Wish v2](https://charm.land/wish/v2) — SSH server

**Web**
- [Astro](https://astro.build) — static site framework
- [Tailwind CSS v4](https://tailwindcss.com) — styling
- TypeScript

**Infrastructure**
- Hetzner CAX11 VPS (ARM64)
- Caddy — web server with automatic HTTPS
- systemd — process management
- Cloudflare — DNS

## Run locally

**TUI**

```bash
go run .
```

Listens on `localhost:2222` by default:

```bash
ssh localhost -p 2222
```

Override with env vars:

```bash
HOST=0.0.0.0 PORT=2222 go run .
```

| Variable | Default     | Description         |
| -------- | ----------- | ------------------- |
| `HOST`   | `localhost` | Bind address        |
| `PORT`   | `2222`      | Bind port (1–65535) |

> **Note:** binding to port 22 requires root or elevated privileges. Use a port above 1024 for local development.

**Web**

```bash
cd web
npm install
npm run dev
```

## Deploy

Both the TUI and web are deployed automatically by pushing a version tag:

```bash
git tag v1.0.0
git push --tags
```

**TUI deploy** (`deploy.yml`):
1. Builds the Go binary for Linux ARM64
2. Uploads to VPS and replaces the binary atomically
3. Restarts the systemd service
4. Creates a GitHub Release with release notes

**Web deploy** (`deploy-web.yml`):
1. Builds the Astro site (`npm run build`)
2. Uploads `dist/` to `/var/www/edwinboon.dev/` on the VPS
3. Caddy serves the updated files immediately — no restart needed

## Infrastructure

The VPS runs the following services:

| Service | Port | Description |
|---------|------|-------------|
| Wish (Go) | 22 | SSH portfolio |
| Caddy | 80, 443 | Web landing page |
| sshd | 2222 | Server management |

## Navigate

| Key     | Action         |
| ------- | -------------- |
| `j/k`   | Move up / down |
| `enter` | Select page    |
| `q`     | Back / quit    |
