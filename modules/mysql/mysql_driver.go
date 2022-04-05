/**
 * @Author: dengmingcong
 * @Description:
 * @File:  mysql_driver
 * @Version: 1.0.0
 * @Date: 2022/03/06 12:38 上午
 */

package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"yangcong/config"
)

type DBMysql struct {
	DB *sql.DB
	Name string
}

func NewMySqlConn(config config.MysqlConfig) (*sql.DB, error){
	//数据库连接
	dataSourceName := config.DBUser + ":" + config.Password + "@tcp" + "(" + config.Host + ":" + config.Port + ")/" + config.DBName
	fmt.Println(dataSourceName)
	//db1 := orm.RegisterDataBase("default", "mysql", dataSourceName)
	db,_:=sql.Open("mysql",dataSourceName)
	err :=db.Ping()
	return db, err
}