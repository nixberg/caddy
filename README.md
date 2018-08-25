# caddy
Snap of the Caddy web server.

Telemetry is disabled. Includes the Cloudflare DNS provider plugin.

## Usage:

```bash
sudo snap set caddy caddyfile="$(cat Caddyfile)"
sudo snap set caddy email="user@example.test"
```

```bash
sudo snap set caddy environment="CLOUDFLARE_EMAIL=user@example.test CLOUDFLARE_API_KEY=abcdefgh"
```

```bash
cat Caddyfile | caddy.check
```

Optional:
```bash
sudo snap connect caddy:home
```

```
example.test {
    root /home/user/www/
}
```