/**
 * @Author: dengmingcong
 * @Description:
 * @File:  base32
 * @Version: 1.0.0
 * @Date: 2022/01/24 3:00 上午
 */

package node

const GeoHashIndex = "0123456789bcdefghjkmnpqrstuvwxyz"

func (positionRange PositionRange) EncodeBase32() string {
	result := ""
	length := len(positionRange.GeoHash)
	//fmt.Println(length)
	for i := 0;i < length; i += 5 {
		var index uint32
		//fmt.Printf("%d:\n", i)
		if i > length {
			for _, bit := range positionRange.GeoHash[i : ] {
				index = ( index | bit ) << 1
			}
			index = index >> 1
			//fmt.Println(GeoHashIndex[index])
		}else{
			for _, bit := range positionRange.GeoHash[i : i + 5] {
				index = ( index | bit ) << 1
			}
			index = index >> 1
			//fmt.Printf("%b\n", index)
		}
		//fmt.Println(" ")
		result += string(GeoHashIndex[index])
	}
	//fmt.Println(result)
	return result
}
