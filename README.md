# caddy
Snap of the Caddy web server.

## Usage:

```bash
sudo snap set caddy caddyfile="$(cat Caddyfile)"
sudo snap set caddy email="user@example.test"
```

```bash
sudo snap set caddy environment="A=1 B=2"
```

```bash
cat Caddyfile | caddy.check
```

Optional:
```bash
sudo snap connect caddy:home
sudo chown roor.root /home/user/www
```

```
example.test {
    root /home/user/www/
}
```