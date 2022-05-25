/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/05/04 12:38 上午
 */

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"yangcong/grpc/pb_file"
	"yangcong/service/relationship/service"
)

func main()  {
	server := service.NewRelationShipService()
	//Default返回一个默认的路由引擎
	go func() {
		r := gin.Default()
		r.POST("/like", server.Like)
		r.POST("/dislike", server.DisLike)
		r.Run("0.0.0.0:8082") // listen and serve on 0.0.0.0:8080
	}()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("端口监听成功：9001")
	grpcServer := grpc.NewServer()
	pb_file.RegisterRelationshipServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}