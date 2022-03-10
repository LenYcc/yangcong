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

func NewMySqlConn(config config.MysqlConfig) *sql.DB{
	//数据库连接
	db,_:=sql.Open("mysql","root:root@(106.55.149.13:3306)/golang")
	err :=db.Ping()
	if err != nil{
		fmt.Println("数据库链接失败")
	}
	defer db.Close()

	//多行查询
	rows,_:=db.Query("select * from stu")
	var id,name string
	for rows.Next(){
		rows.Scan(&id,&name)
		fmt.Println(id,name)
	}
	return db
}