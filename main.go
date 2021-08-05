package main

import (
	"github.com/wmdev4/shipswift-gateway/core"
	helloworld "github.com/wmdev4/shipswift-gateway/test/service"
)

func main() {
	addr := ":8091"
	//starting testing service
	go func(_addr string) {
		helloworld.Start(_addr,"dev")
	}(addr)
	core.Start()
	//create http client
}
