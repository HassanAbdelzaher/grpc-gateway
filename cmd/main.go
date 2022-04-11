package main

import (
	"flag"
	"log"

	"github.com/wmdev4/shipswift-gateway/core"
	"github.com/wmdev4/shipswift-gateway/version"
)

var WithQue = flag.String("que", "false", "debug")

func main() {
	flag.Parse()
	useQue := false
	if WithQue != nil && (*WithQue == "true" || *WithQue == "1") {
		useQue = true
	}
	log.Println("staring gateway version " + version.VERSION)
	core.Start(useQue)
}
