/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/03/15 10:45 下午
 */

package main

import (
	"net/http"
	"yangcong/service/node/service"
)

func main() {
	server := service.NewNodeServer()
	http.HandleFunc("/search", server.SearchHttp)
	http.ListenAndServe("127.0.0.1:8001", nil)
}
