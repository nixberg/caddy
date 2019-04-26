package main

import (
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
