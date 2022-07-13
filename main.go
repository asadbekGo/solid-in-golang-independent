package main

import (
	"github.com/asadbekGo/solid-in-golang-independent/dip"
	"github.com/asadbekGo/solid-in-golang-independent/isp"
	"github.com/asadbekGo/solid-in-golang-independent/lsp"
	"github.com/asadbekGo/solid-in-golang-independent/ocp"
	"github.com/asadbekGo/solid-in-golang-independent/srp"
)

func main() {

	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
	dip.Run()
}
