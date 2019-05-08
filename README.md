# caddy
Snap of the Caddy web server.

Telemetry is disabled.

## Usage:

```bash
sudo snap set caddy caddyfile="$(cat Caddyfile)"
sudo snap set caddy email="user@example.test"
```

```bash
cat Caddyfile | caddy.validate
```
