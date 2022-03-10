/**
 * @Author: dengmingcong
 * @Description:
 * @File:  pg_driver
 * @Version: 1.0.0
 * @Date: 2022/01/24 10:00 上午
 */

package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"yangcong/config"
)

func NewDBConn(config config.PGConfig) *sql.DB{
	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=%v",config.Host, config.Port, config.DBName, "verify-full host")
	db, _ := sql.Open("postgres", dataSourceName)
	return db
	// ...
}
