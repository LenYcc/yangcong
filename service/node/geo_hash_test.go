/**
 * @Author: dengmingcong
 * @Description:
 * @File:  geo_hash_test
 * @Version: 1.0.0
 * @Date: 2022/01/24 2:31 上午
 */

package node

import (
	"fmt"
	"testing"
)

func TestGeoHash(t *testing.T) {
	//position1 := Position{
	//	Latitude:  39.901136325561474,
	//	Longitude: 116.68228026611021,
	//}
	tree := NewPrefixTree()
	tree.Insert(39.90113632516148,116.68230367445747,  10)
	tree.Insert(39.90113632516148,116.68230367445747,  101)
	tree.Insert(39.90113632516148,116.68230367445747,  102)
	tree.Insert(39.90113632516148,116.68230367445747,  103)
	tree.Insert(39.90113632516148,116.68230367445747,  104)
	tree.Insert(39.90113632516148,116.68230367445747,  105)
	tree.Insert(39.90113632516148,116.68230367445747,  106)
	tree.Insert(-39.90113632516148,116.68230367445747,  109)
	fmt.Println(tree.Search(39.90113632516148,116.68230367445747,0))

}
