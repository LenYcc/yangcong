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
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"yangcong/config"
)

type DBMysql struct {
	DB *sql.DB
	Name string
}

func NewMySqlConn(config config.MysqlConfig) *sql.DB{
	//数据库连接
	dataSourceName := config.DBUser + ":" + config.Password + "@tcp" + "(" + config.Host + ":" + config.Port + ")/" + config.DBName
	fmt.Println(dataSourceName)
	db1 := orm.RegisterDataBase("default", "mysql", dataSourceName)
	db2,_:=sql.Open("mysql",dataSourceName)
	err :=db2.Ping()
	if err != nil{
		fmt.Println("数据库链接失败")
	}
	defer db2.Close()
	return db
}