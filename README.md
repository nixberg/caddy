# caddy
Snap of the [Caddy](https://caddyserver.com) web server.

## Usage

Snaps are sandboxed, so use the [API](https://caddyserver.com/docs/api) to set the [`Caddyfile`](https://caddyserver.com/docs/caddyfile):

```shell
curl localhost:2019/load --request POST --header 'Content-Type: text/caddyfile' --data-binary @Caddyfile
```
