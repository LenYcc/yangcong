/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/05/03 4:42 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	"yangcong/service/recall"
)

func main()  {
	server := recall.NewRecallServiceHandler()
	r := gin.Default()
	r.POST("/search", server.Search)
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
}
