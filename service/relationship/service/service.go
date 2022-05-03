/**
 * @Author: dengmingcong
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/02/07 5:05 下午
 */

package service

import (
	"yangcong/config"
	"yangcong/modules/redis"
)

type RelationshipServer struct {

}

func NewRelationShipService()(*RelationshipServer)  {
	redisConfig := config.RedisConfig{
		Type: "tcp",
		Host: "127.0.0.1",
		Port: "3306",
	}
	conn := redis.NewRedisClient(redisConfig)
	return nil
}
