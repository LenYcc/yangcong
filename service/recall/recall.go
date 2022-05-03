/**
 * @Author: dengmingcong
 * @Description:
 * @File:  recall
 * @Version: 1.0.0
 * @Date: 2022/03/15 11:36 下午
 */

package recall

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"sort"
	"strconv"
	"yangcong/grpc/pb_file"
)

type RecallServiceHandler struct {
	NodeConn *grpc.ClientConn
	NodeClient pb_file.NodeServiceClient
}

func NewRecallServiceHandler () (*RecallServiceHandler) {
	recallServiceHandler := &RecallServiceHandler{}
	conn, err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	recallServiceHandler.NodeConn = conn
	client := pb_file.NewNodeServiceClient(conn)
	recallServiceHandler.NodeClient = client
	return recallServiceHandler
}

func (server RecallServiceHandler)Search(c *gin.Context) {
	request := &pb_file.RecallRequest{}

	if id := c.PostForm("userId");id != "" {
		idInt,err := strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "message" + ":" + "id解析错误。")
			return
		}else {
			request.UserId  = int32(idInt)
		}
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请附带 id。")
		return
	}

	if XStr := c.PostForm("X");XStr != "" {
		X , err := strconv.ParseFloat(XStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "message" + ":" + "X解析错误。")
			return
		}else {
			request.X  = float32(X)
		}
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请附带 X。")
		return
	}

	if YStr := c.PostForm("Y");YStr != "" {
		Y , err := strconv.ParseFloat(YStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "message" + ":" + "Y解析错误。")
			return
		}else {
			request.Y  = float32(Y)
		}
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请附带 Y。")
		return
	}

	if gender := c.PostForm("gender");gender != "" {
		genderInt,err := strconv.Atoi(gender)
		if err != nil {
			c.String(http.StatusBadRequest, "message" + ":" + "gender解析错误。")
			return
		}else {
			request.Gender  = int32(genderInt)
		}
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请附带性别。")
		return
	}

	if radius := c.PostForm("radius");radius != "" {
		radiusInt,err := strconv.Atoi(radius)
		if err != nil {
			c.String(http.StatusBadRequest, "message" + ":" + "radius解析错误。")
			return
		}else {
			request.Radius = int32(radiusInt)
		}
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请附带搜索距离。")
		return
	}

	reply, err := server.NodeClient.Recall(context.Background(), request)
	if  err != nil {
		c.String(http.StatusBadRequest, "message" + ":" + "搜索失败。")
		return
	}
	fmt.Println(len(reply.Users))
	scoreMap := make(map[int32]float64)
	//开始打分
	for _, user := range reply.Users {
		if user.Tags == reply.Me.Tags {
			scoreMap[user.UserId] = float64(1) + float64(0.5 * user.Pop)
		}else{
			countA := uint32(0)
			countB := uint32(0)
			a := reply.Me.Tags & user.Tags
			b := reply.Me.Tags | user.Tags
			for a > 0 {
				// 位移后和1与运算，如果为1那最低一位就是1
				countA += a & 1
				// 向右移一位，最低位舍弃
				a >>= 1
			}
			for b > 0 {
				// 位移后和1与运算，如果为1那最低一位就是1
				countB += b & 1
				// 向右移一位，最低位舍弃
				b >>= 1
			}
			scoreMap[user.UserId] = float64(countA / countB) + float64(0.5 * user.Pop)
		}
	}
	sortByScore := SortByScore{
		Score: scoreMap,
		Users: reply.Users,
	}
	sort.Sort(sortByScore)
	c.String(http.StatusOK, "Me" + ":" + reply.Me.String() + "\n")

	for _, user := range reply.Users {
		c.String(http.StatusOK, "User" + ":" + user.String()  + "\n")
	}
}

type SortByScore struct {
	Score map[int32]float64
	Users []*pb_file.User
}

func (p SortByScore) Len() int {
	return len(p.Users)
}
func (p SortByScore) Less(i, j int) bool {
	return p.Score[p.Users[i].UserId] > p.Score[p.Users[j].UserId]
}
func (p SortByScore) Swap(i, j int) {
	p.Users[i], p.Users[j] = p.Users[j], p.Users[i]
}
