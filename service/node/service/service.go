/**
 * @Author: dengmingcong
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/03/06 12:55 上午
 */

package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"yangcong/config"
	"yangcong/grpc/pb_file"
	"yangcong/modules/mysql"
	"yangcong/service/node"
	"yangcong/service/node/dto"
)

type NodeService struct {
	MySqlDB  *sql.DB
	UserTags *node.UserTags
	GeoHash  *node.PrefixTree
}

func NewNodeServer() *NodeService {

	nodeService := &NodeService{}
	//初始化 Mysql
	mysqlConfig := config.MysqlConfig{
		//Host:     "81.68.70.6",
		Host:     "106.55.149.13",
		Port:     "3306",
		DBName:   "yangcong", //todo 修改数据库名字
		DBUser:   "root",
		Password: "Dmc362@@",
		Timeout:  0,
		SSLMode:  "",
	}
	var err error
	nodeService.MySqlDB, err = mysql.NewMySqlConn(mysqlConfig)
	if err != nil {
		fmt.Errorf("初始化 Mysql 失败:ERROR:%v", err)
		return nil
	}
	fmt.Println(nodeService.MySqlDB , "------------")
	//初始化 GeoHashMap
	nodeService.GeoHash = node.NewPrefixTree()
	fmt.Println(nodeService.GeoHash, "------------")
	//初始化 UserTags
	nodeService.UserTags = node.NewUserTags("User", nodeService.MySqlDB)
	fmt.Println(nodeService.UserTags, "------------")
	//读取 mysql 数据构建索引
	go func() {
		nodeService.UserTags.RunSchedulerJob(nodeService.GeoHash)
	}()
	fmt.Println("服务初始化完成，正在读取用户数据。。。。。")
	return nodeService
}


func (NodeService NodeService) SearchHttp(w http.ResponseWriter, r *http.Request) {
	reply := NodeService.Search(dto.SearchRequest{
		UserId: 1,
		Radius: 15,
		Gender: 1,
		X:      116.397,
		Y:      39.9165,
	})
	for _, user := range reply.Users {
		fmt.Fprintln(w, "ID:" , user.UserId)
		fmt.Fprintln(w, "TAGS:" ,user.Tags)
		fmt.Fprintln(w, "GENDER:" ,user.Gender)
		fmt.Fprintln(w, "POP:" ,user.Pop)
		fmt.Fprintln(w, "DISTANCE:" ,user.Distance)
	}

}

func (NodeService NodeService) Search(request dto.SearchRequest) dto.SearchReply {
	//解析 deep
	deep := ParserRadius(request.Radius)
	fmt.Println("deep",deep)
	ids := NodeService.GeoHash.Search(request.X, request.Y, deep)

	// 拿取用户信息
	reply := dto.SearchReply{}
	reply.Users = []*dto.User{}
	for _, id := range ids {
		user := NodeService.UserTags.Get(int64(id))
		if user != nil && user.UserId != request.UserId{
			reply.Users = append(reply.Users, user)
		}
	}
	reply.Me = NodeService.UserTags.Get(request.UserId)
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


func (n NodeService) Recall(context context.Context, request *pb_file.RecallRequest) (*pb_file.RecallReply, error) {
	reply := &pb_file.RecallReply{
		Users: nil,
		Meta:  nil,
	}
	fmt.Println(request.String())
	Users := []*pb_file.User{}
	searchReply := n.Search(dto.SearchRequest{
		UserId: int64(request.UserId),
		Radius: int64(request.Radius),
		Gender: request.Gender,
		X:      float64(request.X),
		Y:      float64(request.Y),
	})

	fmt.Println(len(searchReply.Users))
	if searchReply.Users != nil {
		for i, _ := range searchReply.Users {
			if searchReply.Users[i].Gender == request.Gender {
				Users = append(Users, &pb_file.User{
					UserId:   int32(searchReply.Users[i].UserId),
					Tags: 	  uint32(searchReply.Users[i].Tags),
					Gender:   searchReply.Users[i].Gender,
					Pop:      float32(searchReply.Users[i].Pop),
					Distance: float32(searchReply.Users[i].Distance),
				})
			}
		}
	}
	reply.Users = Users
	reply.Me = &pb_file.User{
		UserId:   int32(searchReply.Me.UserId),
		Tags:     uint32(searchReply.Me.Tags),
		Gender:   searchReply.Me.Gender,
		Pop:      float32(searchReply.Me.Pop),
		Distance: float32(searchReply.Me.Distance),
	}
	reply.Meta = &pb_file.Meta{
		Code:    "OK",
		Message: "",
	}

	return reply, nil
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