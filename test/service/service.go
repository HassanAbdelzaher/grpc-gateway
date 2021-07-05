/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package helloworld

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

//Server ... is used to implement helloworld.GreeterServer.
type Server struct {
	TAG string
}
//	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
func (s *Server) Check (context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse,error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	},nil
}

func (s *Server) Watch(req *grpc_health_v1.HealthCheckRequest,srv grpc_health_v1.Health_WatchServer) error{

	return nil
}


// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md)
	}
	return &HelloReply{Message: s.TAG + ":Welcom and Hello " + in.Name}, nil
}
func (s *Server) SayHelloStream(in *HelloRequest, srv Greeter_SayHelloStreamServer) error {
	//ctx := srv.Context()
	for i := 0; i < 5; i++ {
		srv.Send(&HelloReply{Message: s.TAG + ":WELCOM STREAM"})
		time.Sleep(1 * time.Second)
	}
	return nil
}
func Start(addr string, tag string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := Server{TAG: tag}
	service:=&server
	RegisterGreeterServer(s, service)
	reflection.Register(s)
	grpc_health_v1.RegisterHealthServer(s, service)
	fmt.Print("starting service ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	Start(":8087", "dev")
}
