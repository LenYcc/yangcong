/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/04/17 7:37 下午
 */

package main

import (
	"net/http"
	"yangcong/service/user/service"
)

func main() {
	server := service.NewUserServer()
	http.HandleFunc("/register", server.Register)
	http.HandleFunc("/login", server.Login)
	http.HandleFunc("/update", server.Update)
	http.ListenAndServe("127.0.0.1:8002", nil)
}
