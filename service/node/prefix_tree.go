/**
 * @Author: dengmingcong
 * @Description: geoHash 前缀树实现
 * @File:  prefix_tree
 * @Version: 1.0.0
 * @Date: 2022/01/24 3:04 上午
 */

package node

import "fmt"

//const MaxDeep = 8

var globalPrefixTree *PrefixTree

type PrefixTreeNode struct {
	Children map[uint8]*PrefixTreeNode
	Depth int
	IsEnd bool
	Key string
	bitmap *BitMap
}

type PrefixTree struct {
	Root *PrefixTreeNode
	GeoHashBitMap GeoHashBitMap
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		Root: &PrefixTreeNode{
			Children:   make(map[uint8]*PrefixTreeNode),
			Depth:      0,
			IsEnd:      false,
		},
		GeoHashBitMap: make(GeoHashBitMap),
	}
}

func NewPrefixTreeNode() *PrefixTreeNode {
	return & PrefixTreeNode {
			Children:   make(map[uint8]*PrefixTreeNode),
			Depth:      0,
			IsEnd:      false,
	}
}

//func init()  {
//	globalPrefixTree = NewPrefixTree()
//}
//
//func GetPrefixTree() *PrefixTree {
//	return globalPrefixTree
//}

func (prefixTree PrefixTree) InsertMapKey(key string) {
	v := prefixTree.GeoHashBitMap[key]
	if v == nil {
		prefixTree.GeoHashBitMap[key] = NewBitMap()
	}
	node := prefixTree.Root
	for i := 0 ;i< len(key);i++ {
		//fmt.Print(key[i], " " , string(key[i]) , " ")
		_, OK := node.Children[key[i]]
		if !OK {
			node.Children[key[i]] = NewPrefixTreeNode()
		}
		node.Depth = i
		node = node.Children[key[i]]
	}
	node.IsEnd = true
	node.Key = key
	if node.bitmap == nil {
		node.bitmap = prefixTree.GeoHashBitMap[key]
	}
	//fmt.Println()
}

//func (prefixTree PrefixTree) SearchKey(key string) *BitMap {
//	node := prefixTree.Root
//	for i := 0; i < len(key); i++ {
//		_, ok := node.Children[key[i]]
//		if !ok {
//			return nil
//		}
//		node = node.Children[key[i]]
//	}
//	if node.IsEnd && node.bitmap != nil {
//		return node.bitmap
//	}
//	return nil
//}

// 判断字典树中是否有指定前缀的单词
func (prefixTree PrefixTree) StartsWith(prefix string) bool {
	node := prefixTree.Root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.Children[prefix[i]]
		if !ok {
			return false
		}
		node = node.Children[prefix[i]]
	}
	return true
}

// 判断字典树中是否有指定前缀的单词
func (prefixTree PrefixTree) GetStartsWith(prefix string)(result map[string]*BitMap) {
	node := prefixTree.Root
	result = make(map[string]*BitMap)
	for i := 0; i < len(prefix); i++ {
		_, ok := node.Children[prefix[i]]
		if !ok {
			fmt.Println("break")
			return nil
		}
		node = node.Children[prefix[i]]
	}
	//fmt.Println(node.Depth)

	queue := make([]*PrefixTreeNode, 0)
	for _, treeNode := range node.Children {
		//fmt.Println(string(u))
		queue = append(queue, treeNode)
	}
	for len(queue) != 0 {
		tmpQueue := make([]*PrefixTreeNode, 0)
		for _, treeNode := range queue {
			if treeNode.Children != nil {
				for _, prefixTreeNode := range treeNode.Children {
					tmpQueue = append(tmpQueue, prefixTreeNode)
					//fmt.Println(string(u))
				}
			}
			if treeNode.IsEnd{
				//fmt.Println(treeNode.IsEnd, treeNode.Key)
				result[treeNode.Key] = treeNode.bitmap
			}
		}
		queue = tmpQueue
	}
	return result
}
