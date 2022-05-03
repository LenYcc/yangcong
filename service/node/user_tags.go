/**
 * @Author: dengmingcong
 * @Description:
 * @File:  user_tags
 * @Version: 1.0.0
 * @Date: 2022/03/21 7:23 下午
 */

package node

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	dto2 "yangcong/service/node/dto"
	"yangcong/service/user/dto"
)

type TagTypeEnum int

const (
	Movie = TagTypeEnum(iota)
	Music
	Book
	Sport
	Meet
	Food
	Travel
	Fashion
	Tea
	Draw
)

type UserTags struct {
	Name   string
	Mysql  *sql.DB
	User   map[int]dto.User
	Tags   map[TagTypeEnum]*BitMap
	Gender map[int]*BitMap
	pop    map[int64]float64
}

func (cache *UserTags) refresh(prefixTree *PrefixTree) error {
	fmt.Println(cache.Mysql.Stats())
	rows, _ := cache.Mysql.Query("select * from user")
	fmt.Println("Query success！！")
	user := dto.User{}
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.Gender, &user.Tel, &user.Name, &user.Birthdate, &user.X, &user.Y, &user.Pop, &user.HeadImg, &user.Country,
			&user.City, &user.Tags, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			fmt.Println("读取数据错误", err)
			return err
		} else {
			logTime := false
			if user.UserId % 1000 == 0 {
				fmt.Println("ID",user.UserId)
				logTime = true
			}
			Time := 0
			if logTime {
				fmt.Println("Begin",Time)
				Time = int(time.Now().UnixNano())
			}
			if user.Gender == 1 {
				if cache.Gender[1] == nil {
					cache.Gender[1] = NewBitMap()
				}
				cache.Gender[1].Set(int64(user.UserId))
			}else{
				if cache.Gender[0] == nil {
					cache.Gender[0] = NewBitMap()
				}
				cache.Gender[0].Set(int64(user.UserId))
			}
			if logTime {
				fmt.Println("GENDER",int(time.Now().UnixNano()) - Time)
				Time = int(time.Now().UnixNano())
			}
			cache.pop[int64(user.UserId)] = user.Pop
			if logTime {
				fmt.Println("POP",int(time.Now().UnixNano()) - Time)
				Time = int(time.Now().UnixNano())
			}
			prefixTree.Insert(user.X, user.Y, int64(user.UserId))
			if logTime {
				fmt.Println("GEOHASH",int(time.Now().UnixNano()) - Time)
				Time = int(time.Now().UnixNano())
			}
			tags := strings.Split(user.Tags, "/")
			for _, tag := range tags {
				tagInt, errAtoi := strconv.Atoi(tag)
				if errAtoi != nil {
					fmt.Errorf("Atoi err %v", errAtoi)
				} else {
					if cache.Tags[TagTypeEnum(tagInt)] == nil {
						cache.Tags[TagTypeEnum(tagInt)] = NewBitMap()
					}
					cache.Tags[TagTypeEnum(tagInt)].Set(int64(user.UserId))
				}
			}
			if logTime {
				fmt.Println("TAGS",int(time.Now().UnixNano()) - Time)
				Time = int(time.Now().UnixNano())
			}
		}
	}
	fmt.Println("用户初始化完毕")
	return nil
}

func (cache *UserTags) RunSchedulerJob(prefixTree *PrefixTree) {
	timer := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-timer.C:
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Errorf("UserTagCache[%s] - RunSchedulerJob Error %v", cache.Name, err)
					}
				}()
				fmt.Println("开始刷新数据")
				err := cache.refresh(prefixTree)
				if err != nil {
					timer.Reset(time.Second * time.Duration(120))
				}
				timer.Reset(time.Hour * time.Duration(24))
			}()
		}
	}
}

func NewUserTags(tagsName string, db *sql.DB) (userTags *UserTags) {
	userTags = &UserTags{}
	userTags.Name = tagsName
	userTags.Mysql = db
	userTags.Gender = make(map[int]*BitMap)
	userTags.Tags = make(map[TagTypeEnum]*BitMap)
	userTags.pop = make(map[int64]float64)
	return userTags
}

func (cache *UserTags) Get(id int64 ) *dto2.User {
	user  := &dto2.User{
		UserId:     id,
		Pop:        0,
		Tags:       0,
	}
	if cache.Gender[1].Has(id) {
		user.Gender = 1
	}else if cache.Gender[0].Has(id) {
		user.Gender = 0
	}else{
		return nil
	}
	user.Pop = cache.pop[id]
	if user.Pop == 0 {
		if user.Gender == 1 {
			user.Pop = 0.3
		}else{
			user.Pop = 0.5
		}
	}
	for enum, bitMap := range cache.Tags {
		if bitMap.Has(id) {
			user.Tags = user.Tags | 1 << enum
		}
	}
	return user
}