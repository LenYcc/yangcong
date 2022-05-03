/**
 * @Author: dengmingcong
 * @Description:
 * @File:  dto
 * @Version: 1.0.0
 * @Date: 2022/04/06 1:37 上午
 */

package dto

type SearchRequest struct {
	UserId int64
	Radius int64
	Gender int32
	X      float64
	Y      float64
}

type SearchReply struct {
	Users []*User
	Me   *User
}

type User struct {
	UserId   int64
	Pop      float64
	Gender   int32
	Tags     uint16
	Distance float64
}

type Meta struct {
	Status  StatusCode
	Message string
}

type StatusCode int32

const (
	OK          = 200
	ClientError = 400
	ServerError = 500
)
