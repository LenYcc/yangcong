/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/04/17 7:37 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	"yangcong/service/user/service"
)

//func main() {
//	server := service.NewUserServer()
//	http.HandleFunc("/register", server.Register)
//	http.HandleFunc("/login", server.Login)
//	http.HandleFunc("/update", server.Update)
//	http.ListenAndServe("127.0.0.1:8002", nil)
//
//
//}


func main() {
	//Default返回一个默认的路由引擎
	server := service.NewUserServer()
	r := gin.Default()
	r.POST("/register", server.Register)
	r.POST("/login", server.Login)
	r.POST("/update", server.Update)
	r.Run("0.0.0.0:8083") // listen and serve on 0.0.0.0:8080
}