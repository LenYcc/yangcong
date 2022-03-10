/**
 * @Author: dengmingcong
 * @Description:
 * @File:  geo_hash_test
 * @Version: 1.0.0
 * @Date: 2022/01/24 2:31 上午
 */

package geohash

import (
	"fmt"
	"testing"
)

func TestGeoHash(t *testing.T) {
	position1 := Position{
		Latitude:  39.901136325561474,
		Longitude: 116.68228026611021,
	}
	key := position1.Insert(10)
	fmt.Println(key)
	position2 := Position{
		Latitude:  39.90113632516148,
		Longitude: 116.68230367445747,
	}
	key = position2.Insert(20)
	fmt.Println(key)
	position3 := Position{
		Latitude:  40.90113632516148,
		Longitude: 116.68230367445747,
	}
	key = position3.Insert(30)
	fmt.Println(key)
	fmt.Println(position2.Search(4))
	fmt.Println("======================================")
	fmt.Println(position2.Search(1))
}
