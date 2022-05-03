/**
 * @Author: dengmingcong
 * @Description:
 * @File:  redis_driver
 * @Version: 1.0.0
 * @Date: 2022/05/04 12:30 上午
 */

package redis

import (
	"fmt"
	"yangcong/config"
)
import "github.com/garyburd/redigo/redis"

func NewRedisClient(config config.RedisConfig)(*redis.Conn) {
	conn, err := redis.Dial(config.Type, config.Host + ":" + config.Port)
	if err != nil {
		fmt.Println("连接出错,", err)
		return nil
	}
	fmt.Println("连接成功....")

	////向redis写入数据
	//_, err = conn.Do("set", "name", "小明") //等价于在redis服务端执行 set name 小明
	//if err != nil {
	//	fmt.Println("写入错误", err)
	//	return
	//}
	return &conn
}
