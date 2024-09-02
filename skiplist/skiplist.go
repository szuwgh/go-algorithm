package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
)

const (
	maxLevel int     = 16   //最大层数
	p        float32 = 0.25 //概率因子
)

type skipList struct {
	header *node
	len    int
	level  int
}

//节点
type node struct {
	key     []byte
	value   []byte
	forward []*node
}

func newSkipList() *skipList {
	return &skipList{
		header: &node{forward: make([]*node, maxLevel)},
	}
}

//两层 1/4
//三层 1/4*1/4
//n层 1/4^(n-1)
func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

func (sl *skipList) Print() {
	for i := sl.level - 1; i >= 0; i-- {
		dummy := sl.header
		for dummy.forward[i] != nil {
			fmt.Print(string(dummy.forward[i].key), "-", string(dummy.forward[i].value), " ")
			dummy = dummy.forward[i]
		}
		fmt.Print("\n")
	}
}

func (sl *skipList) Front() *node {
	return sl.header.forward[0]
}

func (n *node) Next() *node {
	if n != nil {
		return n.forward[0]
	}
	return nil
}

//查找
func (sl *skipList) search(key []byte) *node {
	dummy := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		//寻找每层最后一个比key小的节点
		for dummy.forward[i] != nil && bytes.Compare(dummy.forward[i].key, key) == -1 {
			dummy = dummy.forward[i]
		}
	}
	dummy = dummy.forward[0]
	if dummy != nil && bytes.Compare(dummy.key, key) == 0 {
		return dummy
	}
	return nil
}

/**
  x------>6------------------------------------->nil
  x------>6--------------------------->25------->nil
  x------>6------>9------------------->25------->nil
  x-->3-->6-->7-->9-->12---->19-->21-->25-->26-->nil
						  ^--17
*/
func (sl *skipList) insert(key, value []byte) {
	//需要更新的节点
	update := make([]*node, maxLevel)
	dummy := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		//寻找每层最后一个比key小的节点
		for dummy.forward[i] != nil && bytes.Compare(dummy.forward[i].key, key) == -1 {
			dummy = dummy.forward[i]
		}
		//记录每个level最后一个小于key的节点
		update[i] = dummy
	}
	dummy = dummy.forward[0]
	if dummy != nil && bytes.Compare(dummy.key, key) == 0 {
		dummy.value = value
		return
	}
	level := randomLevel()
	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.header
		sl.level = level
	}
	n := &node{key: key, value: value, forward: make([]*node, maxLevel)}
	for i := 0; i < level; i++ {
		n.forward[i] = update[i].forward[i]
		update[i].forward[i] = n
	}
	sl.len++
	return
}
