/**
 * @Author: dengmingcong
 * @Description:
 * @File:  mysql_driver_test
 * @Version: 1.0.0
 * @Date: 2022/03/13 11:15 下午
 */

package mysql

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	config2 "yangcong/config"
	"yangcong/modules/mysql"
	"yangcong/service/user/dto"
)

func TestName(t *testing.T) {
	go func() {TestNewMySqlConn1(t)}()
	go func() {TestNewMySqlConn2(t)}()
	select {

	}
}

func TestNewMySqlConn1(t *testing.T) {
	config := config2.MysqlConfig{
		Host:     "106.55.149.13",
		Port:     "3306",
		DBName:   "yangcong",
		DBUser:   "root",
		Password: "Dmc362@@",
		Timeout:  0,
		SSLMode:  "",
	}
	db, _ := mysql.NewMySqlConn(config)
	//orm := orm.RegisterDataBase("default", )
	//fmt.Println("12312312")
	//result, _ := db.Query("select * from test1")
	//results := make([]string,3)
	//for result.Next() {
	//	err := result.Scan(&results[0],&results[1],&results[2])
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(results)
	//}
	telMap := make(map[string]bool)
	//rows, _ := db.Query("select * from user")
	user2 := dto.User{}
	//for s, b := range telMap {
	//	fmt.Println(s,b)
	//}

	i := 0
	tx, _ := db.Begin()
	for i < 5000000 {
		//user2.UserId =  i
		user2.Name = "yangcong" + fmt.Sprint(rand.Intn(10000000))
		user2.X = float64(rand.Intn(62)) + 73 + rand.Float64()
		user2.Y = float64(rand.Intn(50)) + 3 + rand.Float64()
		//xx := rand.Intn(18)
		//if xx % 2 == 0 {
		//	user2.X = -1 * user2.X
		//}
		//yy := rand.Intn(18)
		//if yy % 2 == 0 {
		//	user2.Y = -1 * user2.Y
		//}
		zz := rand.Intn(18)
		if zz%2 == 0 {
			user2.Gender = 1
		} else {
			user2.Gender = 0
		}

		telBool := true
		for telBool {
			tel := rand.Intn(9999999999) + 10000000000
			if telMap[fmt.Sprint(tel)] == false {
				user2.Tel = int64(tel)
				telMap[fmt.Sprint(tel)] = true
				telBool = false
			}
		}

		user2.Pop = rand.Float64()

		user2.Name += fmt.Sprint(i)

		user2.City = rand.Intn(30) + 1

		tmm := time.Unix(int64(639215645+rand.Intn(1649058935-639215645)), 0)

		user2.Birthdate = fmt.Sprint(tmm.Format("2006-01-02 03:04:05"))

		numMap := make(map[int]bool)
		xxx := rand.Intn(10)
		numMap[xxx] = true
		tag := fmt.Sprint(xxx) + "/"

		for ii := 0; ii < 10; ii++ {
			xxxx := rand.Intn(10)
			if !numMap[xxxx] {
				tag += fmt.Sprint(xxxx) + "/"
				numMap[xxxx] = true
			}
		}

		//fmt.Println(tag[len(tag) -1 :])
		if tag[len(tag)-1:] == "/" {
			tag = tag[0 : len(tag)-3]
		}

		user2.Tags = tag

		//fmt.Printf("INSERT INTO table_name ( user_id, gender,tel, name, birthdate, x, y, pop, head_img, country, city, tags )VALUES( '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );\n",
		//	user2.UserId, user2.Gender, user2.Tel, user2.Name, user2.Birthdate, user2.X, user2.Y, user2.Pop, user2.HeadImg,
		//	user2.Country, user2.City, user2.Tags)
		//fmt.Println(tx.Query(fmt.Sprintf("INSERT INTO user ( user_id, gender,tel, name, birthdate, x, y, pop, head_img, country, city, tags )VALUES( '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );",
		//	user2.UserId, user2.Gender, user2.Tel, user2.Name, user2.Birthdate, user2.X, user2.Y, user2.Pop, user2.HeadImg,
		//	user2.Country, user2.City, user2.Tags)))

		fmt.Println(tx.Query(fmt.Sprintf("INSERT INTO user (gender,tel, name, birthdate, x, y, pop, head_img, country, city, tags )VALUES(  '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );",
			user2.Gender, user2.Tel, user2.Name, user2.Birthdate, user2.X, user2.Y, user2.Pop, user2.HeadImg,
			user2.Country, user2.City, user2.Tags)))

		i++
		//tx.Commit()
		//time.Sleep(time.Millisecond * 100)
		if i % 10000 == 0 {
			tx.Commit()
			tx, _ = db.Begin()
		}
	}

	db.Close()
}

func TestNewMySqlConn2(t *testing.T) {
	config := config2.MysqlConfig{
		Host:     "81.68.70.6",
		Port:     "3306",
		DBName:   "yangcong",
		DBUser:   "root",
		Password: "Dmc362@@",
		Timeout:  0,
		SSLMode:  "",
	}
	db, _ := mysql.NewMySqlConn(config)
	//orm := orm.RegisterDataBase("default", )
	//fmt.Println("12312312")
	//result, _ := db.Query("select * from test1")
	//results := make([]string,3)
	//for result.Next() {
	//	err := result.Scan(&results[0],&results[1],&results[2])
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(results)
	//}
	telMap := make(map[string]bool)
	rows, _ := db.Query("select * from user")
	user2 := dto.User{}
	for rows.Next() {
		err := rows.Scan(&user2.UserId, &user2.Gender, &user2.Tel, &user2.Name, &user2.Birthdate, &user2.X, &user2.Y, &user2.Pop, &user2.HeadImg,
			&user2.Country, &user2.City, &user2.Tags, &user2.CreateTime, &user2.UpdateTime)
		if err != nil {
			fmt.Println(err)
		}
		telMap[fmt.Sprint(user2.Tel)] = true
		fmt.Println(user2.Tel)
	}
	//for s, b := range telMap {
	//	fmt.Println(s,b)
	//}

	i := 0
	tx, _ := db.Begin()
	for i < 5000000 {
		//user2.UserId =  i
		user2.Name = "yangcong" + fmt.Sprint(rand.Intn(10000000))
		user2.X = float64(rand.Intn(2)) + 115 + rand.Float64()
		user2.Y = float64(rand.Intn(2)) + 39 + rand.Float64()
		//xx := rand.Intn(18)
		//if xx % 2 == 0 {
		//	user2.X = -1 * user2.X
		//}
		//yy := rand.Intn(18)
		//if yy % 2 == 0 {
		//	user2.Y = -1 * user2.Y
		//}
		zz := rand.Intn(18)
		if zz%2 == 0 {
			user2.Gender = 1
		} else {
			user2.Gender = 0
		}

		telBool := true
		for telBool {
			tel := rand.Intn(9999999999) + 10000000000
			if telMap[fmt.Sprint(tel)] == false {
				user2.Tel = int64(tel)
				telMap[fmt.Sprint(tel)] = true
				telBool = false
			}
		}

		user2.Pop = rand.Float64()

		user2.Name += fmt.Sprint(i)

		user2.City = rand.Intn(30) + 1

		tmm := time.Unix(int64(639215645+rand.Intn(1649058935-639215645)), 0)

		user2.Birthdate = fmt.Sprint(tmm.Format("2006-01-02 03:04:05"))


		tag := ""

		for ii := 0; ii < 10; ii++ {
			xxxx := rand.Intn(10)
			if xxxx < 4 {
				tag += fmt.Sprint(ii) + "/"
			}
		}

		tag = tag[0 : len(tag)-1]

		user2.Tags = tag

		//fmt.Printf("INSERT INTO table_name ( user_id, gender,tel, name, birthdate, x, y, pop, head_img, country, city, tags )VALUES( '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );\n",
		//	user2.UserId, user2.Gender, user2.Tel, user2.Name, user2.Birthdate, user2.X, user2.Y, user2.Pop, user2.HeadImg,
		//	user2.Country, user2.City, user2.Tags)
		//fmt.Println(tx.Query(fmt.Sprintf("INSERT INTO user ( user_id, gender,tel, name, birthdate, x, y, pop, head_img, country, city, tags )VALUES( '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v' );",
		//	user2.UserId, user2.Gender, user2.Tel, user2.Name, user2.Birthdate, user2.X, user2.Y, user2.Pop, user2.HeadImg,
		//	user2.Country, user2.City, user2.Tags)))

		fmt.Println(tx.Query(fmt.Sprintf("INSERT INTO user (gender,tel, name, birthdate, x, y, pop, country, city, tags )VALUES(  '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v',  '%v', '%v' );",
			user2.Gender, user2.Tel, user2.Name, user2.Birthdate, user2.X, user2.Y, user2.Pop,
			user2.Country, user2.City, user2.Tags)))

		i++
		//tx.Commit()
		//time.Sleep(time.Millisecond * 100)
		if i % 10000 == 0 {
			tx.Commit()
			tx, _ = db.Begin()
		}
	}

	db.Close()
}
