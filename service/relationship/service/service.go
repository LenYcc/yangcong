/**
 * @Author: dengmingcong
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/02/07 5:05 下午
 */

package service

import (
	"context"
	"fmt"
	redis2 "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yangcong/config"
	"yangcong/grpc/pb_file"
	"yangcong/modules/redis"
)

type RelationshipService struct {
	redisClient redis2.Conn
}


func NewRelationShipService() *RelationshipService {
	relationshipServer := &RelationshipService{}
	redisConfig := config.RedisConfig{
		Type: "tcp",
		Host: "127.0.0.1",
		Port: "6379",
	}
	conn := redis.NewRedisClient(redisConfig)

	relationshipServer.redisClient = conn
	return relationshipServer
}

func (relation *RelationshipService) Like(c *gin.Context) {

	userId := 0
	otherUserId := 0

	if id := c.PostForm("userId"); id != "" {
		var err error
		userId, err = strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "message"+":"+"请输入正确的用户 id。")
			return
		}
	} else {
		c.String(http.StatusBadRequest, "message"+":"+"请输入用户 id。")
	}

	if id := c.PostForm("otherUserId"); id != "" {
		var err error
		otherUserId, err = strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "message"+":"+"请输入正确的用户 oid。")
			return
		}
	} else {
		c.String(http.StatusBadRequest, "message"+":"+"请输入用户 oid。")
	}

	//向redis写入数据
	_, err := relation.redisClient.Do("lpush", "like:"+fmt.Sprint(userId), otherUserId) //等价于在redis服务端执行 set name 小明
	if err == nil {
		fmt.Println("写入成功", err)
		return
	}
}

func (relation *RelationshipService) DisLike(c *gin.Context) {

	userId := 0
	otherUserId := 0

	if id := c.PostForm("userId"); id != "" {
		var err error
		userId, err = strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "message"+":"+"请输入正确的用户 id。")
			return
		}
	} else {
		c.String(http.StatusBadRequest, "message"+":"+"请输入用户 id。")
	}

	if id := c.PostForm("otherUserId"); id != "" {
		var err error
		otherUserId, err = strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "message"+":"+"请输入正确的用户 oid。")
			return
		}
	} else {
		c.String(http.StatusBadRequest, "message"+":"+"请输入用户 oid。")
	}
	//向redis写入数据
	_, err := relation.redisClient.Do("lpush", "dislike:"+fmt.Sprint(userId), otherUserId) //等价于在redis服务端执行 set name 小明
	if err == nil {
		fmt.Println("写入成功", err)
		return
	}
}

func (relation *RelationshipService) GetLike(context context.Context, request *pb_file.LikeRequest) (*pb_file.LikeReply, error) {
	userId := request.UserId

	//向redis写入数据
	result, err := redis2.Ints(relation.redisClient.Do("lrange", "like:"+fmt.Sprint(userId), "0", "-1")) //等价于在redis服务端执行 set name 小明
	if err != nil {
		fmt.Println("读取失败", err)
		return nil,err
	}
	reply := &pb_file.LikeReply{
		Ids:  make(map[int32]string),
		Meta: nil,
	}
	for _, id := range result {
		reply.Ids[int32(id)] = ""
	}
	reply.Meta = &pb_file.Meta{
		Code:    "OK",
		Message: "",
	}
	return reply, nil
}

func (relation *RelationshipService) GetDislike(context context.Context, request *pb_file.DislikeRequest) (*pb_file.DislikeReply, error) {
	userId := request.UserId

	//向redis写入数据
	result, err := redis2.Ints(relation.redisClient.Do("lrange", "dislike:"+fmt.Sprint(userId), "0", "-1")) //等价于在redis服务端执行 set name 小明
	if err != nil {
		fmt.Println("读取失败", err)
		return nil,err
	}
	reply := &pb_file.DislikeReply{
		Ids:  make(map[int32]string),
		Meta: nil,
	}
	for _, id := range result {
		reply.Ids[int32(id)] = ""
	}
	reply.Meta = &pb_file.Meta{
		Code:    "OK",
		Message: "",
	}
	return reply, nil
}

func (relation *RelationshipService) GetAll(context context.Context, request *pb_file.AllRequest) (*pb_file.AllReply, error) {
	userId := request.UserId

	//向redis写入数据
	result, err := redis2.Ints(relation.redisClient.Do("lrange", "like:"+fmt.Sprint(userId), "0", "-1")) //等价于在redis服务端执行 set name 小明
	if err != nil {
		fmt.Println("读取失败", err)
		return nil ,err
	}

	reply := &pb_file.AllReply{
		Ids:  make(map[int32]string),
		Meta: nil,
	}

	for _, id := range result {
		reply.Ids[int32(id)] = ""
	}

	result, err = redis2.Ints(relation.redisClient.Do("lrange", "dislike:"+fmt.Sprint(userId), "0", "-1")) //等价于在redis服务端执行 set name 小明
	if err != nil {
		fmt.Println("读取失败", err)
		return nil, err
	}
	for _, id := range result {
		reply.Ids[int32(id)] = ""
	}
	reply.Meta = &pb_file.Meta{
		Code:    "OK",
		Message: "",
	}
	return reply, nil
}