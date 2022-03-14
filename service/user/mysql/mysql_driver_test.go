/**
 * @Author: dengmingcong
 * @Description:
 * @File:  mysql_driver_test
 * @Version: 1.0.0
 * @Date: 2022/03/13 11:15 下午
 */

package mysql

import (
	"testing"
	config2 "yangcong/config"
    "github.com/astaxie/beego/orm"
)

func TestNewMySqlConn(t *testing.T) {
	config := config2.MysqlConfig{
		Host:     "106.55.149.13",
		Port:     "3306",
		DBName:   "test",
		DBUser:   "root",
		Password: "Dmc362@@",
		Timeout:  0,
		SSLMode:  "",
	}
	db := NewMySqlConn(config)
	orm := orm.RegisterDataBase("default", )
}