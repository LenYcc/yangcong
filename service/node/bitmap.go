/**
 * @Author: dengmingcong
 * @Description:
 * @File:  bitmap
 * @Version: 1.0.0
 * @Date: 2022/02/03 3:21 上午
 */

package node

import (
	"fmt"
	"sync"
)

var onceBitMap sync.Once

var BitMap map[string]*bitmap

type bitmap struct {
	keys []byte
	len int
}

const ByteSize = 8

func init() {
	onceBitMap.Do(func() {
		BitMap = make(map[string]*bitmap)
	})
}

func GetBitMap() map[string]*bitmap {
	return BitMap
}

func NewBitMap() *bitmap {
	return &bitmap{keys:make([]byte, 0), len:0}
}

func (b *bitmap)has(v int64) bool {

	if b.len == 0 {
		return false
	}

	index := v / 8

	byteIndex :=byte(v % 8)

	if int(index) >len(b.keys) { //todo not exist
		return false
	}

	if b.keys[index]&(1<<byteIndex) != 0 {
		return true
	}
	return false
}

func (b *bitmap)set(v int64) {

	index := v / 8

	byteIndex :=byte(v % 8)

	for b.len <= int(index) {

		b.keys =append(b.keys, 0)

		b.len++

	}
	b.keys[index] =b.keys[index] | (1 << byteIndex)
}

func (b *bitmap)del(v int64) {

	index := v / 8

	byteIndex :=byte(v % 8)

	for b.len <= int(index) {

		b.keys =append(b.keys, 0)

		b.len++

	}
	b.keys[index] =b.keys[index] | (0 << byteIndex)
}

func (b *bitmap)scan(c int, limit int)(result []int) {
	if b.len == 0 {
		return result
	}
	if c > b.len {
		c = 0
	}
	if limit > b.len {
		for i := 0;i < b.len; i++  {
			for j := byte(0); j <= ByteSize; j ++ {
				//fmt.Println(b.keys[i], 1 << j, b.keys[i] & byte(1 << j))
				if b.keys[i] & byte(1 << j) != 0 {
					result = append(result,  int(j) + i * ByteSize)
					limit --
				}
			}
		}
		return
	}
	for i := c + 1;i != c ; i++ {
		if i == b.len {
			if c == 0 {
				break
			}
			i = 1
		}
		if b.keys[i] == 0 {
			continue
		}
		for j := byte(0); j <= ByteSize; j ++ {
			//fmt.Println(b.keys[i], 1 << j, b.keys[i] & byte(1 << j))
			if b.keys[i] & byte(1 << j) != 0 {
				result = append(result,  int(j) + i * ByteSize)
				limit --
			}
		}
		if limit <= 0 {
			break
		}
	}
	return
}

func (b *bitmap)length()int {
	return b.len
}

func (b *bitmap)print() {
	for _, v :=range b.keys {
		fmt.Printf("%08b\n", v)
	}
}