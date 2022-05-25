/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/03/15 10:45 下午
 */

package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"yangcong/grpc/pb_file"
	"yangcong/service/node/service"
)

func main() {

	server := service.NewNodeServer()
	go func() {
		http.HandleFunc("/search", server.SearchHttp)
		http.ListenAndServe("0.0.0.0:8000", nil)
	}()


	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("端口监听成功：9000")
	grpcServer := grpc.NewServer()
	pb_file.RegisterNodeServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
