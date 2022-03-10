/**
 * @Author: dengmingcong
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/10/12 9:42 上午
 */

package server

type HttpServer struct {
	ServerName string
	ServerRoute string
	ServerOption *interface{}
}

func (*HttpServer) NewHttpServer() error {
	return nil
}
