package test

import (
	"context"
	"github.com/wmdev4/shipswift-gateway/grpc"
	helloworld "github.com/wmdev4/shipswift-gateway/test/service"
	"testing"
	"time"
)
var infoProvider=grpc.ReflectionServiceProvider{}

func TestServiceMethodUsingReflection(t *testing.T){
	//fullName := fmt.Sprintf("%v.%v.%v", pckg, service, method)
	addr := "localhost:8066"
	//starting testing service
	go func(_addr string) {
		helloworld.Start(_addr,"dev")
	}(addr)
	//waiting service startup
	<-time.After(5*time.Second)
	//retrive services
	services, err := infoProvider.ListMethods(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}
	if len(services)==0{
		t.Error("Expected one service while found nothing")
		return
	}
	//Search for helloworld.Greeter
	isServiceFound:=false
	var order int =-1
	for id,srv:=range services{
		if srv.SeriveName=="helloworld.Greeter"{
			isServiceFound=true
			order=id
		}
	}
	if !isServiceFound{
		t.Error("missing service:helloworld.Greeter")
		return
	}

	//testing methods
	srv:=services[order]
	methods:=srv.Methods
	if len(methods)==0{
		t.Error("missing methods in service")
	}
	// SayHello
	//SayHelloStream
	sayHello:=false
	sayHelloStream:=false
	for _,m:=range methods{
		t.Log(m.Name)
		if m.Name=="SayHello"{
			sayHello=true
		}
		if m.Name=="sayHelloStream"{
			sayHelloStream=true
		}
	}
	if !sayHello{
		t.Error("Missing SayHello Method")
		return
	}
	if !sayHelloStream{
		t.Error("Missing SayHelloStream Method")
		return
	}
}