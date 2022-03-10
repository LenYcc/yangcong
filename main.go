/**
 * @Author: dengmingcong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/01/17 3:32 下午
 */

package main

import "fmt"

const FILE_PATH = "./dict.txt"

func main(){
	//KValue := map[string][][]byte{}
	//f, err0 := os.Open(FILE_PATH)
	//if err0 != nil{
	//	fmt.Errorf("OPen file fial err [%v]", err0)
	//	return
	//}
	//rd := bufio.NewReader(f)
	//for {
	//	line, err1 := rd.ReadString('\n')
	//	if err1 != nil || err1 == io.EOF {
	//		break
	//	}
	//	tmpKV := strings.Fields(line)
	//	//fmt.Println(tmpKV[1])
	//	KValue[tmpKV[1]] = append(KValue[tmpKV[1]], tmpKV[0])
	//}
	//for k, v := range KValue {
	//	fmt.Println(k)
	//	fmt.Println(v)
	//	var tmp [][]byte(){}
	//	for i, str := range v {
	//		byte()
	//	}
	//}
	a := "你好呀"
	b := "阿斯加"
	c := []byte(a)
	d := []byte(b)
	e := []byte{}
	for i, _ := range c {
		x := c[i] | d[i]
		e = append(e,x)
	}
	fmt.Println(e)
	fmt.Println(string(e))
	f := []byte("你斯加")
	for i, _ := range c {
		x := e[i] | f[i]
		f = append(f,x)
	}
	fmt.Println(f)
	fmt.Println(string(f))
}
