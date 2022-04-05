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
	"math/rand"
	"net/http"
	"yangcong/config"
	"yangcong/modules/mysql"
	"yangcong/service/node"
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
		Host:     "106.55.149.13",
		Port:     "3306",
		DBName:   "test", //todo 修改数据库名字
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

	//初始化 GeoHashMap
	nodeServiceHandler.GeoHash = node.NewPrefixTree()

	//初始化 UserTags
	nodeServiceHandler.UserTags = node.NewUserTags("User", nodeServiceHandler.MySqlDB)
	//读取 mysql 数据构建索引
	nodeServiceHandler.UserTags.RunSchedulerJob()



	return nodeServiceHandler
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "测试路径,表示程序开始")
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