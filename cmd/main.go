package main

import (
	"log"

	"github.com/wmdev4/shipswift-gateway/core"
)

var VERSION = "1.1.1"

func main() {
	log.Println("staring gateway version " + VERSION)
	core.Start()
}
