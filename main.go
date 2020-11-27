package main

import (
	"fmt"
	"log"
	"context"
	"net"
	"google.golang.org/grpc"
	"gocool/proto"
)

type routeServer struct {
	// proto.mustEmbedUnimplementedObserveServer
}

var (
	RouteHost = "localhost"
	RoutePort = 7777
)

// rpc SayRoute(RouteReq) returns (RouteResp) {}
func (r *routeServer) SayRoute(ctx context.Context, req *proto.RouteReq) (*proto.RouteResp, error) {
	fmt.Println(req.GetName())
	return &proto.RouteResp{
		Ip:   RouteHost,
		Port: fmt.Sprintf("%d", RoutePort),
	}, nil
}

func main() {
	fmt.Println("hello world")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", RouteHost, RoutePort))
	if err != nil {
		log.Fatalf("failed to listen: +%v", err)
	}
	r := routeServer{}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterObserveServer(grpcServer, &r)
	log.Fatal(grpcServer.Serve(lis))

	/*
	req := proto.RouteReq{
		Name: "RouteName",
	}
	resp, err := r.SayRoute(context.Background(), &req)
	if err == nil {
		fmt.Printf("ip[%s] port[%s]\n", resp.GetIp(), resp.GetPort())
	}*/
}
