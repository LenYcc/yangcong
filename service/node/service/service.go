/**
 * @Author: dengmingcong
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/03/06 12:55 上午
 */

package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"yangcong/config"
	"yangcong/modules/mysql"
	"yangcong/service/node"
	"yangcong/service/node/dto"
)

type NodeServiceHandler struct {
	MySqlDB  *sql.DB
	UserTags *node.UserTags
	GeoHash  *node.PrefixTree
}

func NewNodeServer() *NodeServiceHandler {
	nodeServiceHandler := &NodeServiceHandler{}
	//初始化 Mysql
	mysqlConfig := config.MysqlConfig{
		Host:     "81.68.70.6",
		//Host:     "106.55.149.13",
		Port:     "3306",
		DBName:   "yangcong", //todo 修改数据库名字
		DBUser:   "root",
		Password: "Dmc362@@",
		Timeout:  0,
		SSLMode:  "",
	}
	var err error
	nodeServiceHandler.MySqlDB, err = mysql.NewMySqlConn(mysqlConfig)
	if err != nil {
		fmt.Errorf("初始化 Mysql 失败:ERROR:%v", err)
		return nil
	}
	fmt.Println(nodeServiceHandler.MySqlDB , "------------")
	//初始化 GeoHashMap
	nodeServiceHandler.GeoHash = node.NewPrefixTree()
	fmt.Println(nodeServiceHandler.GeoHash, "------------")
	//初始化 UserTags
	nodeServiceHandler.UserTags = node.NewUserTags("User", nodeServiceHandler.MySqlDB)
	fmt.Println(nodeServiceHandler.UserTags, "------------")
	//读取 mysql 数据构建索引
	go func() {
		nodeServiceHandler.UserTags.RunSchedulerJob(nodeServiceHandler.GeoHash)
	}()
	fmt.Println("服务初始化完成，正在读取用户数据。。。。。")
	return nodeServiceHandler
}


func (nodeServiceHandler NodeServiceHandler) SearchHttp(w http.ResponseWriter, r *http.Request) {
	reply := nodeServiceHandler.Search(dto.SearchRequest{
		UserId: 1,
		Radius: 78,
		Gender: 1,
		X:      116.397,
		Y:      39.9165,
	})
	for _, user := range reply.Users {
		fmt.Fprintln(w, user.UserId)
		fmt.Fprintln(w, user.Tags)
		fmt.Fprintln(w, user.Gender)
		fmt.Fprintln(w, user.Pop)
		fmt.Fprintln(w, user.Distance)
	}

}

func (nodeServiceHandler NodeServiceHandler) Search(request dto.SearchRequest) dto.SearchReply {
	//解析 deep
	deep := ParserRadius(request.Radius)
	fmt.Println("deep",deep)
	ids := nodeServiceHandler.GeoHash.Search(request.X, request.Y, deep)
	//TODO 获取过滤列表

	//TODO 拿取用户信息
	reply := dto.SearchReply{}
	reply.Users = []*dto.User{}
	for _, id := range ids {
		user := nodeServiceHandler.UserTags.Get(int64(id))
		if user != nil {
			reply.Users = append(reply.Users, user)
		}
	}
	return reply
}

func ParserRadius(radius int64) (deep int) {
	switch radius {
	case 1:
		deep = 5
	case 20:
		deep = 4
	case 78:
		deep = 3
	case 500:
		deep = 2
	default:
		deep = 3
	}
	return deep
}


//func GeoHashSetHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, r.Header)
//	position := node.Position{
//		Latitude:  39.90113632516148,
//		Longitude: 116.68230367445747,
//	}
//	position.Insert(int64(rand.Intn(200)))
//	position = node.Position{
//		Latitude:  40.90113632516148,
//		Longitude: 116.68230367445747,
//	}
//	position.Insert(int64(rand.Intn(200)))
//	position = node.Position{
//		Latitude:  31.90113632516148,
//		Longitude: 120.68230367445747,
//	}
//	position.Insert(int64(rand.Intn(200)))
//	fmt.Fprintln(w, position.Search(1))
//}