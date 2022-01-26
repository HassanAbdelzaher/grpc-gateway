package main

import (
	"log"

	"github.com/wmdev4/shipswift-gateway/core"
	"github.com/wmdev4/shipswift-gateway/version"
)

func main() {
	log.Println("staring gateway version " + version.VERSION)
	core.Start()
}
