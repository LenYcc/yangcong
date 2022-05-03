/**
 * @Author: dengmingcong
 * @Description:
 * @File:  config_struct
 * @Version: 1.0.0
 * @Date: 2022/01/26 10:25 上午
 */

package config

import (
	"time"
)

type PGConfig struct {
	Host string
	Port string
	DBName string
	DBUser string
	Password string
	Timeout time.Duration
	SSLMode string
}

type MysqlConfig struct {
	Host string
	Port string
	DBName string
	DBUser string
	Password string
	Timeout time.Duration
	SSLMode string
}

type RedisConfig struct {
	Type string
	Host string
	Port string
}

//type KafkaConfig struct {
//	Brokers    []string
//	Zookeepers []string
//	Zookeeper  KafkaZookeeperConfig
//	Producer   kafka.TopicConfig
//	Consumer   *KafkaConsumerConfig
//	Group      *KafkaGroupConfig
//}

type KafkaZookeeperConfig struct {
	Chroot string
}

