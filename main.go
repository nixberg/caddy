package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
