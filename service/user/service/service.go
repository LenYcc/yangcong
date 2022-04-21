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
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
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
	Logger log.Logger
}

func NewUserServer() *UserServiceHandler {
	userServiceHandler := &UserServiceHandler{}

	logFile, errOpen := os.OpenFile("./user.log-" + time.Now().Format("2006-01-02"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if errOpen != nil {
		fmt.Errorf("Open file err, %v", errOpen)
	}
	userServiceHandler.Logger.SetOutput(logFile)
	userServiceHandler.Logger.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	userServiceHandler.Logger.Println("服务开始启动。")
	//初始化 Mysql
	mysqlConfig := config.MysqlConfig{
		//Host: "81.68.70.6",
		Host:     "106.55.149.13",
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

func (userServiceHandler UserServiceHandler) Register(c *gin.Context) {
	success := true
	user := dto.User{}
	tx, _ := userServiceHandler.MySqlDB.Begin()

	if c.PostForm("name") != "" {
		user.Name = c.PostForm("name")
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请输入正确的姓名。")
		success = false
	}

	gender := 0
	if c.PostForm("gender") == "male" {
		gender = 1
	}else if c.PostForm("gender") == "female" {
		gender = 0
	}else {
		c.String(http.StatusBadRequest, "message" + ":" + "请输入正确的性别。")
		success = false
	}
	user.Gender = gender
	fmt.Println(gender)

	telString := c.PostForm("tel")
	fmt.Println(telString)
	tel , err := strconv.Atoi(telString)
	if err != nil || len(telString) != 11 {
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"请输入正确的电话号码。"}))
		success = false
	}
	user.Tel = int64(tel)

	birthdateString := c.PostForm("birthdate")
	birthdate , err := strconv.Atoi(birthdateString)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"请输入正确生日。"}))
		success = false
	}
	user.Birthdate = time.Unix(int64(birthdate),0).Format("2006-01-02 03:04:05")

	XString := c.PostForm("X")
	X , err := strconv.ParseFloat(XString, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"坐标错误。"}))
		success = false
	}
	user.X = X

	YString := c.PostForm("Y")
	Y , err := strconv.ParseFloat(YString, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"坐标错误。"}))
		success = false
	}
	user.Y = Y

	countryString := c.PostForm("country")
	country, err := strconv.Atoi(countryString)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"请输入正确国家。"}))
		success = false
	}
	user.Country = country

	cityString := c.PostForm("city")
	city, err := strconv.Atoi(cityString)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"请输入正确城市。"}))
		success = false
	}
	user.City = city

	tagsString := c.PostForm("tags")
	if tagsString != "" {
		user.Tags = tagsString
	}else{
		c.String(http.StatusBadRequest, fmt.Sprint(gin.H{"message":"请输入标签。"}))
		success = false
	}
	user.Pop = 0

	if success {
		fmt.Println(tx.Query(fmt.Sprintf("INSERT INTO user (gender,tel, name, birthdate, x, y, pop, country, city, tags )VALUES(  '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );",
			user.Gender, user.Tel, user.Name, user.Birthdate, user.X, user.Y, user.Pop,
			user.Country, user.City, user.Tags)))
		tx.Commit()
		c.String(http.StatusOK, fmt.Sprint(gin.H{"message":"注册成功。"}))
		userServiceHandler.Logger.Println("注册了一个用户：id 为 ： %v", user.UserId)
	}
	return
}

func (userServiceHandler UserServiceHandler) Update(c *gin.Context) {
	doUpdate := false
	user := dto.User{}

	tx, _ := userServiceHandler.MySqlDB.Begin()
	userId := c.PostForm("userId")
	if userId != "" {
		rows, err := tx.Query(fmt.Sprintf("SELECT * FROM user where user_id = '%v';", userId))
		if err != nil {
			userServiceHandler.Logger.Fatalf("sql err ,%v ", err)
		}else{
			for rows.Next() {
				errScan := rows.Scan(&user.UserId, &user.Gender, &user.Tel, &user.Name, &user.Birthdate, &user.X, &user.Y, &user.Pop, &user.HeadImg,
					&user.Country, &user.City, &user.Tags, &user.CreateTime, &user.UpdateTime)
				if err != nil {
					userServiceHandler.Logger.Fatalf("sql Scan err %v",errScan)
				}
			}
		}	
	}else {
		userServiceHandler.Logger.Fatalf("request err ,no user id")
	}
	

	if value := c.PostForm("name"); value != "" && value != user.Name  {
		user.Name = c.PostForm("name")
		doUpdate = true
	}

	gender := -1
	if c.PostForm("gender") == "male" {
		gender = 1
	}else if c.PostForm("gender") == "female" {
		gender = 0
	}
	if gender != user.Gender && gender != -1{
		user.Gender = gender
		doUpdate = true
	}

	telString := c.PostForm("tel")
	tel , err := strconv.Atoi(telString)
	if err == nil || len(telString) == 11 {
		if int64(tel) != user.Tel {
			user.Tel = int64(tel)
			doUpdate = true
		}
	}

	birthdateString := c.PostForm("birthdate")
	birthdate , err := strconv.Atoi(birthdateString)
	if err == nil {
		if value := time.Unix(int64(birthdate),0).Format("2006-01-02 03:04:05");value != "" && value != user.Birthdate {
			user.Birthdate = value
			doUpdate = true
		}
	}

	countryString := c.PostForm("country")
	if countryString != ""{
		country, err := strconv.Atoi(countryString)
		if err == nil {
			if country != user.Country {
				user.Country = country
				doUpdate = true
			}
		}
	}

	cityString := c.PostForm("city")
	if cityString != "" {
		city, err := strconv.Atoi(cityString)
		if err == nil {
			if city != user.City {
				user.City = city
				doUpdate = true
			}
		}
	}

	tagsString := c.PostForm("tags")
	if tagsString != "" && tagsString != user.Tags{
		user.Tags = tagsString
		doUpdate = true
	}

	if doUpdate {
		fmt.Println(tx.Query(fmt.Sprintf("UPDATE user SET gender = '%v' , tel = '%v', name = '%v', birthdate = '%v', country = '%v', city = '%v', tags = '%v' WHERE user_id = '%v';",
			user.Gender, user.Tel, user.Name, user.Birthdate,
			user.Country, user.City, user.Tags, user.UserId)))
		tx.Commit()
		c.String(http.StatusOK, fmt.Sprint(gin.H{"message":"更新成功。"}))
		userServiceHandler.Logger.Println("更新用户", userId)
	}
	return
}

func (userServiceHandler UserServiceHandler) Login(c *gin.Context) {

}
