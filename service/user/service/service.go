/**
 * @Author: dengmingcong
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/04/17 7:30 下午
 */

package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"yangcong/config"
	"yangcong/modules/mysql"
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

type UserServiceHandler struct {
	MySqlDB *sql.DB
	TagsMap map[string]int
}

func NewUserServer() *UserServiceHandler {
	userServiceHandler := &UserServiceHandler{}
	//初始化 Mysql
	mysqlConfig := config.MysqlConfig{
		Host: "81.68.70.6",
		//Host:     "106.55.149.13",
		Port:     "3306",
		DBName:   "yangcong", //todo 修改数据库名字
		DBUser:   "root",
		Password: "Dmc362@@",
		Timeout:  0,
		SSLMode:  "",
	}
	var err error
	userServiceHandler.MySqlDB, err = mysql.NewMySqlConn(mysqlConfig)
	if err != nil {
		fmt.Errorf("初始化 Mysql 失败:ERROR:%v", err)
		return nil
	}

	userServiceHandler.TagsMap = make(map[string]int)
	userServiceHandler.TagsMap["Movie"] = 0
	userServiceHandler.TagsMap["Music"] = 1
	userServiceHandler.TagsMap["Book"] = 2
	userServiceHandler.TagsMap["Sport"] = 3
	userServiceHandler.TagsMap["Meet"] = 4
	userServiceHandler.TagsMap["Food"] = 5
	userServiceHandler.TagsMap["Travel"] = 6
	userServiceHandler.TagsMap["Fashion"] = 7
	userServiceHandler.TagsMap["Tea"] = 8
	userServiceHandler.TagsMap["Draw"] = 9


	return userServiceHandler
}

func (userServiceHandler UserServiceHandler) Register(w http.ResponseWriter, r *http.Request) {

	success := true
	user := dto.User{}
	tx, _ := userServiceHandler.MySqlDB.Begin()
	
	gender := 0
	if r.PostForm.Get("gender") == "male" {
		gender = 1
	}else if r.PostForm.Get("gender") == "female" {
		gender = 0
	}else {
		w.Write([]byte("请输入正确的性别。"))
		success = false
	}
	user.Gender = gender

	telString := r.PostForm.Get("tel")
	tel , err := strconv.Atoi(telString)
	if err != nil || len(telString) != 11 {
		w.Write([]byte("请输入正确的电话号码。"))
		success = false
	}
	user.Tel = int64(tel)

	birthdateString := r.PostForm.Get("tel")
	birthdate , err := strconv.Atoi(birthdateString)
	if err != nil {
		w.Write([]byte("请输入正确生日。"))
		success = false
	}
	user.Birthdate = time.Unix(int64(birthdate),0).Format("2006-01-02 03:04:05")

	XString := r.PostForm.Get("X")
	X , err := strconv.ParseFloat(XString, 64)
	if err != nil {
		w.Write([]byte("坐标错误。"))
		success = false
	}
	user.X = X

	YString := r.PostForm.Get("Y")
	Y , err := strconv.ParseFloat(YString, 64)
	if err != nil {
		w.Write([]byte("坐标错误。"))
		success = false
	}
	user.Y = Y

	countryString := r.PostForm.Get("country")
	country, err := strconv.Atoi(countryString)
	if err != nil {
		w.Write([]byte("请输入正确国家。"))
		success = false
	}
	user.Country = country

	cityString := r.PostForm.Get("city")
	city, err := strconv.Atoi(cityString)
	if err != nil {
		w.Write([]byte("请输入正确城市。"))
		success = false
	}
	user.City = city

	tagsString := r.PostForm.Get("tags")
	tags := strings.Split(tagsString, "/")
	tagsInts := ""
	for _, tag := range tags {
		if v, ok := userServiceHandler.TagsMap[tag];ok {
			tagsInts += fmt.Sprint(v) + "/"
		}
	}
	if len(tagsInts) > 0 {
		tagsInts = tagsInts[0:len(tagsInts) - 1]
		user.Tags = tagsInts
	}else{
		w.Write([]byte("请输入标签。"))
		success = false
	}

	user.Pop = 0

	fmt.Println(tx.Query(fmt.Sprintf("INSERT INTO user (gender,tel, name, birthdate, x, y, pop, country, city, tags )VALUES(  '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );",
		user.Gender, user.Tel, user.Name, user.Birthdate, user.X, user.Y, user.Pop,
		user.Country, user.City, user.Tags)))
	tx.Commit()

	if success {
		w.Write([]byte("注册成功。"))
	}
	return
}

func (userServiceHandler UserServiceHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (userServiceHandler UserServiceHandler) Login(w http.ResponseWriter, r *http.Request) {

}
