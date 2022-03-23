/**
 * @Author: dengmingcong
 * @Description:
 * @File:  geo_hash
 * @Version: 1.0.0
 * @Date: 2022/01/24 2:03 上午
 */

package node

import "math/rand"

//Latitude 维度
//Longitude 经度
const (
	MaxLat = 90
	MinLat = -90
	MatLon = 180
	MinLon = -180
)

const MaxDeep20 = 20 //GeoHash 精度为米    8 * 5 = 40

type PositionRange struct {
	MaxLat float64
	MaxLon float64
	MinLat float64
	MinLon float64
	GeoHash []uint32
}

type Position struct {
	Latitude  float64
	Longitude float64
}

var keyMap map[string]struct{}

func (position Position) Insert(id int64) string {
	positionRange := &PositionRange{
		MaxLat: MaxLat,
		MinLat: MinLat,
		MaxLon: MatLon,
		MinLon: MinLon,
		GeoHash: make([]uint32, 0),
	}

	for i := 0; i < MaxDeep20; i++ {
		positionRange.MiddleCut(position)
	}

	//fmt.Printf( "%b",positionRange.GeoHash)
	mapKey := positionRange.EncodeBase32()
	if GetBitMap()[mapKey] != nil {
		GetBitMap()[mapKey].set(id)
	}else{
		GetPrefixTree().Insert(mapKey)
		GetBitMap()[mapKey].set(id)
	}
	//fmt.Print(GetPrefixTree().StartsWith(mapKey))
	return mapKey
}

func (position Position) Delete(id int64) string {
	positionRange := &PositionRange{
		MaxLat: MaxLat,
		MinLat: MinLat,
		MaxLon: MatLon,
		MinLon: MinLon,
		GeoHash: make([]uint32, 0),
	}

	for i := 0; i < MaxDeep20; i++ {
		positionRange.MiddleCut(position)
	}

	//fmt.Printf( "%b",positionRange.GeoHash)
	mapKey := positionRange.EncodeBase32()
	if GetBitMap()[mapKey] != nil {
		GetBitMap()[mapKey].del(id)
	}
	//fmt.Print(GetPrefixTree().StartsWith(mapKey))
	return mapKey
}

func (position Position) Search(deep int) []int {
	positionRange := &PositionRange{
		MaxLat: MaxLat,
		MinLat: MinLat,
		MaxLon: MatLon,
		MinLon: MinLon,
		GeoHash: make([]uint32, 0),
	}

	for i := 0; i < MaxDeep20; i++ {
		positionRange.MiddleCut(position)
	}

	//fmt.Printf( "%b",positionRange.GeoHash)
	mapKey := positionRange.EncodeBase32()
	//fmt.Println(mapKey[0:deep])
	result := []int{}
	for _, bitMap := range GetPrefixTree().GetStartsWith(mapKey[0:deep]) {
		result = append(result, bitMap.scan(rand.Intn(int(bitMap.len)), 2000)...)
	}
	return result
}

func (positionRange *PositionRange) MiddleCut(position Position) (*PositionRange, error)  {
	latMiddle := (positionRange.MaxLat + positionRange.MinLat) / 2.0
	lonMiddle := (positionRange.MaxLon + positionRange.MinLon) / 2.0
	if position.Longitude > lonMiddle{
		positionRange.MinLon = lonMiddle
		positionRange.GeoHash  = append(positionRange.GeoHash, 1)
	}else{
		positionRange.MaxLon = lonMiddle
		positionRange.GeoHash  = append(positionRange.GeoHash, 0)
	}
	if position.Latitude > latMiddle{
		positionRange.MinLat = latMiddle
		positionRange.GeoHash  = append(positionRange.GeoHash, 1)
	}else{
		positionRange.MaxLat = latMiddle
		positionRange.GeoHash  = append(positionRange.GeoHash, 0)
	}

	return positionRange, nil
}

//todo 持久化 插入 前缀树