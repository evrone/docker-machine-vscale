package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/evrone/docker-machine-vscale"
)

func main() {
	plugin.RegisterDriver(vscale.NewDriver("", ""))
}
