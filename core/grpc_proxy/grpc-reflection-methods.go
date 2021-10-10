package grpc_proxy

import (
	"encoding/json"
	"fmt"
	"net/http"

	grpc2 "github.com/wmdev4/shipswift-gateway/grpc_invoke"
)

var infoProvider = grpc2.ReflectionServiceProvider{}

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
