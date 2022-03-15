/**
 * @Author: dengmingcong
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2022/03/06 12:55 上午
 */

package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"yangcong/service/geohash"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "测试路径,表示程序开始")
}

func GeoHashSetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
	position := geohash.Position{
		Latitude:  39.90113632516148,
		Longitude: 116.68230367445747,
	}
	position.Insert(int64(rand.Intn(200)))
	position = geohash.Position{
		Latitude:  40.90113632516148,
		Longitude: 116.68230367445747,
	}
	position.Insert(int64(rand.Intn(200)))
	position = geohash.Position{
		Latitude:  31.90113632516148,
		Longitude: 120.68230367445747,
	}
	position.Insert(int64(rand.Intn(200)))
	fmt.Fprintln(w, position.Search(1))

}