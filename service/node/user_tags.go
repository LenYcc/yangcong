/**
 * @Author: dengmingcong
 * @Description:
 * @File:  user_tags
 * @Version: 1.0.0
 * @Date: 2022/03/21 7:23 下午
 */

package node

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type TagTypeEnum int

const (
	Movie = TagTypeEnum(iota)
	Music
	Book
	Sport
	Meet
)

type User struct {
	UserId int
	Gender int
	Tel       int64
	Name      string
	Birthdate string
	X   float64
	Y       float64
	Pop     float64
	HeadImg string
	Country int
	City    int
	Tags    string
	CreateTime string
	UpdateTime string
}

type UserTags struct {
	Name string
	Mysql *sql.DB
	User map[int]User
	Tags map[TagTypeEnum]*BitMap
	Done bool
}

func (cache *UserTags) refresh(ctx context.Context) {
	rows, _ := cache.Mysql.Query("select * from user")
	user := User{}
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.Gender, &user.Tel, &user.Name, &user.Birthdate, &user.X, &user.Y, &user.Pop, &user.HeadImg, &user.Country,
			&user.City, &user.Tags, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)
	}
}

//func (cache *UserTags) RunSchedulerJob() {
//	ctx := context.Background()
//	timer := time.NewTimer(time.Second * 10)
//	for {
//		select {
//		case <-timer.C:
//			go func() {
//				defer func() {
//					if err := recover(); err != nil {
//						fmt.Errorf("UserTagCache[%s] - RunSchedulerJob Error %v", cache.Name, err)
//					}
//				}()
//				cache.refresh(ctx)
//				timer.Reset(cache.getSchedulerInterval())
//			}()
//		}
//	}
//}
//func (cache *UserTags) getSchedulerInterval() time.Duration {
//	nowHour := (time.Now().UTC().Hour() + 8) % 24
//	var retryInterval int64 = 0
//	//todo 遍历 mysql 拿取用户数据
//	for i, s := range cache.status {
//		if s != DONE {
//			slf.Errorf("UserTagCache[%s] get user tag[%s] status[%v], need retry", cache.name, cache.allTags[i], s)
//			retryInterval = FailedReqInterval
//			return time.Second * time.Duration(retryInterval)
//		}
//	}
//
//	if cache.name == "match_purchase" {
//		retryInterval = GetResetTime(11, 50, 0)  // 每天早上大约11:30点产生，11:50开始查询数据
//	}else if cache.name == "heavy_users" {
//		retryInterval = GetResetTime(11, 50, 0)  // 每天早上大约11:30点产生，11:50开始查询数据
//	} else {
//		if nowHour > cache.frequentRefreshHourRange.End || nowHour < cache.frequentRefreshHourRange.Begin {
//			retryInterval = HourlyRefreshInterval
//		} else {
//			retryInterval = FrequentRefreshInterval
//		}
//	}
//	for i := range cache.status {
//		cache.status[i] = UNSTART
//	}
//	return time.Second * time.Duration(retryInterval)
//}
func GetResetTime(hour int, min int, sec int) int64 {
	currentTime := time.Now()
	resetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hour, min, sec, 0, time.Local)
	if currentTime.Before(resetTime) {
		return int64(resetTime.Sub(currentTime) / time.Second)
	}
	return int64(resetTime.Add(24*time.Hour).Sub(currentTime) / time.Second)
}

func NewUserTags(tagsName string , db *sql.DB) (userTags *UserTags) {
	userTags.Name = tagsName
	userTags.Mysql = db
	return userTags
}