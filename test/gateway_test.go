package test

import (
	"github.com/wmdev4/shipswift-gateway/core"
	helloworld "github.com/wmdev4/shipswift-gateway/test/service"
	"testing"
	"time"
)

func TestGatway(t *testing.T){
	addr := "localhost:8067"
	//starting testing service
	go func(_addr string) {
		helloworld.Start(_addr,"dev")
	}(addr)
	<-time.After(5*time.Second)
	core.Start()
	<-time.After(1*time.Second)
	//create http client
	<-time.After(1*time.Minute)
}
