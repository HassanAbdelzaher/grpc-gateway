package main

import (
	"encoding/json"
	"fmt"
	"github.com/wmdev4/shipswift-gateway/grpc"
	"github.com/wmdev4/shipswift-gateway/core"
	helloworld "github.com/wmdev4/shipswift-gateway/test/service"
	"net/http"
)
var infoProvider=grpc.ReflectionServiceProvider{}

func main() {
	addr := ":8091"
	//starting testing service
	go func(_addr string) {
		helloworld.Start(_addr,"dev")
	}(addr)
	core.Start()
	//create http client
}

func HandleListMethods(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in HandleListMethods", r)
		}
	}()

	var ret = make(map[string]interface{})
	//fullName := fmt.Sprintf("%v.%v.%v", pckg, service, method)
	addr := r.URL.Query().Get("address")
	if addr == "" {
		addr = r.URL.Query().Get("url")
	}
	services, err := infoProvider.ListMethods(r.Context(), addr)
	if err != nil {
		ret["error"] = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}
