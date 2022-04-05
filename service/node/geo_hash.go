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

func (prefixTree PrefixTree) Insert(latitude, longitude float64, id int64) string {
	positionRange := &PositionRange{
		MaxLat: MaxLat,
		MinLat: MinLat,
		MaxLon: MatLon,
		MinLon: MinLon,
		GeoHash: make([]uint32, 0),
	}



	for i := 0; i < MaxDeep20; i++ {
		positionRange.MiddleCut(Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	//fmt.Printf( "%b",positionRange.GeoHash)
	mapKey := positionRange.EncodeBase32()
	if prefixTree.GeoHashBitMap[mapKey] != nil {
		prefixTree.GeoHashBitMap[mapKey].Set(id)
	}else{
		prefixTree.InsertMapKey(mapKey)
		prefixTree.GeoHashBitMap[mapKey].Set(id)
	}
	//fmt.Print(GetPrefixTree().StartsWith(mapKey))
	return mapKey
}

func (prefixTree PrefixTree) Delete(latitude, longitude float64,id int64) string {
	positionRange := &PositionRange{
		MaxLat: MaxLat,
		MinLat: MinLat,
		MaxLon: MatLon,
		MinLon: MinLon,
		GeoHash: make([]uint32, 0),
	}

	for i := 0; i < MaxDeep20; i++ {
		positionRange.MiddleCut(Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	//fmt.Printf( "%b",positionRange.GeoHash)
	mapKey := positionRange.EncodeBase32()
	if prefixTree.GeoHashBitMap[mapKey] != nil {
		prefixTree.GeoHashBitMap[mapKey].Del(id)
	}
	//fmt.Print(GetPrefixTree().StartsWith(mapKey))
	return mapKey
}

func (prefixTree PrefixTree) Search(latitude, longitude float64,deep int) []int {
	positionRange := &PositionRange{
		MaxLat: MaxLat,
		MinLat: MinLat,
		MaxLon: MatLon,
		MinLon: MinLon,
		GeoHash: make([]uint32, 0),
	}

	for i := 0; i < MaxDeep20; i++ {
		positionRange.MiddleCut(Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	//fmt.Printf( "%b",positionRange.GeoHash)
	mapKey := positionRange.EncodeBase32()
	//fmt.Println(mapKey[0:deep])
	result := []int{}
	if deep >= len(mapKey) {
		 deep = len(mapKey) - 1
	}
	for _, bitMap := range prefixTree.GetStartsWith(mapKey[0:deep]) {
		result = append(result, bitMap.Scan(rand.Intn(int(bitMap.len)), 2000)...)
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
